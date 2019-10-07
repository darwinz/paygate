/*
 * Paygate API
 *
 * Paygate is a RESTful API enabling Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transactions to be submitted and received without a deep understanding of a full NACHA file specification.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// Depository struct for Depository
type Depository struct {
	// Depository ID
	ID string `json:"ID,omitempty"`
	// Legal name of the financial institution.
	BankName string `json:"bankName,omitempty"`
	// Legal holder name on the account
	Holder string `json:"holder"`
	// Defines the type of entity of the account holder as an *individual* or *company*
	HolderType string `json:"holderType"`
	// Defines the account as *checking* or *savings*
	Type string `json:"type"`
	// The ABA routing transit number for the depository account.
	RoutingNumber string `json:"routingNumber"`
	// The account number for the depository account
	AccountNumber string `json:"accountNumber"`
	// Defines the status of the Depository account
	Status string `json:"status,omitempty"`
	// Additional meta data to be used for display only
	Metadata    string       `json:"metadata,omitempty"`
	ReturnCodes []ReturnCode `json:"returnCodes,omitempty"`
	Created     time.Time    `json:"created,omitempty"`
	Updated     time.Time    `json:"updated,omitempty"`
}
