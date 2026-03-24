package tessitura

type APINoteTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APINote struct {
	Id              *int                `json:"Id"`
	Constituent     *APIEntity          `json:"Constituent"`
	NoteType        *APINoteTypeSummary `json:"NoteType"`
	Note            *string             `json:"Note"`
	EditIndicator   bool                `json:"EditIndicator"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreatedBy       *string             `json:"CreatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
	UpdatedBy       *string             `json:"UpdatedBy"`
}
