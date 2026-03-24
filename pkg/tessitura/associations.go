package tessitura

type APIAssociationTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIGenderSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
}

type APIAssociation struct {
	Id                    *int                       `json:"Id"`
	Constituent           *APIEntity                 `json:"Constituent"`
	AssociatedConstituent *APIEntity                 `json:"AssociatedConstituent"`
	AssociatedName        *string                    `json:"AssociatedName"`
	AssociationType       *APIAssociationTypeSummary `json:"AssociationType"`
	Gender                *APIGenderSummary          `json:"Gender"`
	BirthDate             *string                    `json:"BirthDate"`
	Inactive              *bool                      `json:"Inactive"`
	StartDate             *string                    `json:"StartDate"`
	EndDate               *string                    `json:"EndDate"`
	Note                  *string                    `json:"Note"`
	ReciprocalAssociation *APIEntity                 `json:"ReciprocalAssociation"`
	EditIndicator         bool                       `json:"EditIndicator"`
	CreatedDateTime       *string                    `json:"CreatedDateTime"`
	CreatedBy             *string                    `json:"CreatedBy"`
	UpdatedDateTime       *string                    `json:"UpdatedDateTime"`
	UpdatedBy             *string                    `json:"UpdatedBy"`
}
