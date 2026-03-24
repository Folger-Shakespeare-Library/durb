package tessitura

type APIAliasTypeSummary struct {
	Id          *int    `json:"Id"`
	Description *string `json:"Description"`
	Inactive    *bool   `json:"Inactive"`
}

type APIAlias struct {
	Id             *int                 `json:"Id"`
	Constituent    *APIEntity           `json:"Constituent"`
	AliasFirstName *string              `json:"AliasFirstName"`
	AliasLastName  *string              `json:"AliasLastName"`
	AliasType      *APIAliasTypeSummary `json:"AliasType"`
	EditIndicator  bool                 `json:"EditIndicator"`
	CreateLocation *string              `json:"CreateLocation"`
	CreatedDateTime *string             `json:"CreatedDateTime"`
	CreatedBy      *string              `json:"CreatedBy"`
	UpdatedDateTime *string             `json:"UpdatedDateTime"`
	UpdatedBy      *string              `json:"UpdatedBy"`
}
