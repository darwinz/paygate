// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

type TELDetail struct {
	PhoneNumber string `json:"phoneNumber"`
}

// createTELBatch creates and returns a TEL ACH batch for use after receiving oral authorization to debit a customer's account.
//
// TEL batches require a telephone number that's answered during typical business hours along with a date and statement of oral
// authorization for a one-time funds transfer. Recurring transfers must contain the total amount of transfers or conditions for
// scheduling transfers. Originators must retain written notice of the authorization for two years.
func createTELBatch(id, userId string, transfer *Transfer, receiver *Receiver, receiverDep *Depository, orig *Originator, origDep *Depository) (ach.Batcher, error) {
	batchHeader := ach.NewBatchHeader()
	batchHeader.ID = id
	batchHeader.ServiceClassCode = ach.DebitsOnly
	batchHeader.CompanyName = orig.Metadata
	batchHeader.StandardEntryClassCode = ach.TEL
	batchHeader.CompanyIdentification = orig.Identification
	batchHeader.CompanyEntryDescription = transfer.Description
	batchHeader.EffectiveEntryDate = base.Now().AddBankingDay(1).Format("060102") // Date to be posted, YYMMDD
	batchHeader.ODFIIdentification = aba8(origDep.RoutingNumber)

	// Add EntryDetail to PPD batch
	entryDetail := ach.NewEntryDetail()
	entryDetail.ID = id
	entryDetail.TransactionCode = determineTransactionCode(transfer, origDep)
	entryDetail.RDFIIdentification = aba8(receiverDep.RoutingNumber)
	entryDetail.CheckDigit = abaCheckDigit(receiverDep.RoutingNumber)
	entryDetail.DFIAccountNumber = receiverDep.AccountNumber
	entryDetail.Amount = transfer.Amount.Int()
	entryDetail.IdentificationNumber = createIdentificationNumber() // TODO(adam): should this be the [required] phone number?
	entryDetail.IndividualName = receiver.Metadata
	entryDetail.DiscretionaryData = transfer.Description // TODO(adam): Or should this be the phone number
	entryDetail.TraceNumber = createTraceNumber(origDep.RoutingNumber)

	// For now just create PPD
	batch, err := ach.NewBatch(batchHeader)
	if err != nil {
		return nil, fmt.Errorf("ACH file %s (userId=%s): failed to create batch: %v", id, userId, err)
	}
	batch.AddEntry(entryDetail)
	batch.SetControl(ach.NewBatchControl())

	if err := batch.Create(); err != nil {
		return batch, err
	}
	return batch, nil
}
