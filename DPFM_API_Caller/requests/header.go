package requests

type header struct {
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBusinessPartner    *int    `json:"FinInstBusinessPartner"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	AccountBusinessPartner    *int    `json:"AccountBusinessPartner"`
	ValidityEndDate           *string `json:"ValidityEndDate"`
	ValidityStartDate         *string `json:"ValidityStartDate"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
