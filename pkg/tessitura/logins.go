package tessitura

type APIWebLogin struct {
	Id                 *int       `json:"Id"`
	Constituent        *APIEntity `json:"Constituent"`
	Login              *string    `json:"Login"`
	LoginType          *APIRef    `json:"LoginType"`
	PrimaryIndicator   *bool      `json:"PrimaryIndicator"`
	TemporaryIndicator *bool      `json:"TemporaryIndicator"`
	Inactive           *bool      `json:"Inactive"`
	LastLoginDate      *string    `json:"LastLoginDate"`
	LockedDate         *string    `json:"LockedDate"`
	FailedAttempts     *int       `json:"FailedAttempts"`
	Email              *APIEntity `json:"Email"`
	CreatedDateTime    *string    `json:"CreatedDateTime"`
	CreatedBy          *string    `json:"CreatedBy"`
	UpdatedDateTime    *string    `json:"UpdatedDateTime"`
	UpdatedBy          *string    `json:"UpdatedBy"`
}
