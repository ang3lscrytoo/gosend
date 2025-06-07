package gosend

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type InvoiceHandler func(invoice *Invoice)

type TrackedInvoice struct {
	InvoiceId int64
	AddedAt   time.Time
}

type PollingManager struct {
	Period          time.Duration
	invoiceHandler  InvoiceHandler
	trackedInvoices map[int64]*TrackedInvoice
	mu              sync.RWMutex
	ctx             context.Context
	cancel          context.CancelFunc
	pollingActive   bool
	pollingWg       sync.WaitGroup
}

func (pm *PollingManager) TrackInvoice(invoice *Invoice) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.trackedInvoices[invoice.InvoiceId] = &TrackedInvoice{
		InvoiceId: invoice.InvoiceId,
		AddedAt:   time.Now(),
	}
}

func (client *Client) StartPolling() error {
	if client.pollingManager.pollingActive {
		return fmt.Errorf("polling already active")
	}

	client.pollingManager.pollingActive = true
	client.pollingManager.pollingWg.Add(1)

	go func() {
		defer client.pollingManager.pollingWg.Done()

		ticker := time.NewTicker(client.pollingManager.Period * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-client.pollingManager.ctx.Done():
				return
			case <-ticker.C:
				if err := client.checkInvoices(); err != nil {
					time.Sleep(5 * time.Second)
				}
			}
		}
	}()

	return nil
}

func (client *Client) StopPolling() {
	if !client.pollingManager.pollingActive {
		return
	}
	client.pollingManager.pollingActive = false
	client.pollingManager.cancel()
	client.pollingManager.pollingWg.Wait()
}

func (client *Client) checkInvoices() error {
	client.pollingManager.mu.RLock()
	invoices := make([]*TrackedInvoice, 0, len(client.pollingManager.trackedInvoices))
	for _, inv := range client.pollingManager.trackedInvoices {
		invoices = append(invoices, inv)
	}
	client.pollingManager.mu.RUnlock()

	if len(invoices) == 0 {
		return nil
	}

	const batchSize = 100
	for i := 0; i < len(invoices); i += batchSize {
		end := i + batchSize
		if end > len(invoices) {
			end = len(invoices)
		}

		batch := invoices[i:end]
		_ = client.checkInvoicesBatch(batch)
	}

	return nil
}

func (client *Client) checkInvoicesBatch(batch []*TrackedInvoice) error {
	if len(batch) == 0 {
		return nil
	}

	var invoiceIDs []string
	for _, inv := range batch {
		invoiceIDs = append(invoiceIDs, strconv.Itoa(int(inv.InvoiceId)))
	}

	invoices, err := client.GetInvoices(GetInvoicesOptions{
		InvoiceIds: invoiceIDs,
	})
	if err != nil {
		return err
	}

	for _, invoice := range invoices {
		client.pollingManager.mu.RLock()
		_, exists := client.pollingManager.trackedInvoices[invoice.InvoiceId]
		client.pollingManager.mu.RUnlock()

		if !exists {
			continue
		}

		if invoice.Status == InvoiceStatusPaid {

			if client.pollingManager.invoiceHandler != nil {
				client.pollingManager.invoiceHandler(invoice)
			}

			client.pollingManager.mu.Lock()
			delete(client.pollingManager.trackedInvoices, invoice.InvoiceId)
			client.pollingManager.mu.Unlock()
		}
	}

	return nil
}
