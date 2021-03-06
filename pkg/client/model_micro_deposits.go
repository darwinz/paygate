/*
 * Paygate API
 *
 * PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.  An organization is a value used to isolate models from each other. This can be set to a \"user ID\" from your authentication service or any value your system has to identify.  There are also [admin endpoints](https://moov-io.github.io/paygate/admin/) for back-office operations.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// MicroDeposits struct for MicroDeposits
type MicroDeposits struct {
	// A microDepositID to identify this set of credits to an external account
	MicroDepositID string `json:"microDepositID"`
	// An array of transferID values created from this micro-deposit
	TransferIDs []string       `json:"transferIDs"`
	Destination Destination    `json:"destination"`
	Amounts     []Amount       `json:"amounts"`
	Status      TransferStatus `json:"status"`
	ProcessedAt *time.Time     `json:"processedAt,omitempty"`
	Created     time.Time      `json:"created"`
}
