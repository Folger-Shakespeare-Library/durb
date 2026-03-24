package domain

import "github.com/Folger-Shakespeare-Library/durb/pkg/tessitura"

// Constituent is the clean domain representation of a Tessitura constituent.
// Fields are flattened from the nested API response into a developer-friendly shape.
type Constituent struct {
	ID               int      `json:"id"`
	FirstName        *string  `json:"firstName"`
	MiddleName       *string  `json:"middleName"`
	LastName         *string  `json:"lastName"`
	DisplayName      *string  `json:"displayName"`
	SortName         *string  `json:"sortName"`
	Prefix           *string  `json:"prefix"`
	Suffix           *string  `json:"suffix"`
	Gender           *string  `json:"gender"`
	Pronoun          *string  `json:"pronoun"`
	ConstituentType  *string  `json:"constituentType"`
	ConstituentGroup *string  `json:"constituentGroup"`
	Inactive         bool     `json:"inactive"`
	InactiveReason   *string  `json:"inactiveReason"`
	ProtectionType   *string  `json:"protectionType"`
	OriginalSource   *string  `json:"originalSource"`
	NameStatus       *string  `json:"nameStatus"`
	MailIndicator    *string  `json:"mailIndicator"`
	EmailIndicator   *string  `json:"emailIndicator"`
	PhoneIndicator   *string  `json:"phoneIndicator"`
	LastActivityDate *string  `json:"lastActivityDate"`
	LastGiftDate     *string  `json:"lastGiftDate"`
	LastTicketDate   *string  `json:"lastTicketDate"`
	CreatedAt        *string  `json:"createdAt"`
	CreatedBy        *string  `json:"createdBy"`
	UpdatedAt        *string  `json:"updatedAt"`
	UpdatedBy        *string  `json:"updatedBy"`
	Addresses        []Address     `json:"addresses,omitempty"`
	Emails           []Email       `json:"emails,omitempty"`
	Phones           []Phone       `json:"phones,omitempty"`
	Salutations      []Salutation  `json:"salutations,omitempty"`
	Affiliations     []Affiliation `json:"affiliations,omitempty"`
	Associations     []Association `json:"associations,omitempty"`
	Notes            []Note        `json:"notes,omitempty"`
}

type Address struct {
	ID          int     `json:"id"`
	Type        *string `json:"type"`
	Primary     bool    `json:"primary"`
	Inactive    bool    `json:"inactive"`
	Street1     *string `json:"street1"`
	Street2     *string `json:"street2"`
	Street3     *string `json:"street3"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	StateCode   *string `json:"stateCode"`
	PostalCode  *string `json:"postalCode"`
	Country     *string `json:"country"`
	CountryCode *string `json:"countryCode"`
	StartDate   *string `json:"startDate"`
	EndDate     *string `json:"endDate"`
	Months      *string `json:"months"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
}

type Email struct {
	ID             int     `json:"id"`
	Type           *string `json:"type"`
	Address        *string `json:"address"`
	Primary        bool    `json:"primary"`
	Inactive       bool    `json:"inactive"`
	IsEmail        bool    `json:"isEmail"`
	AllowHtml      *bool   `json:"allowHtml"`
	AllowMarketing bool    `json:"allowMarketing"`
	StartDate      *string `json:"startDate"`
	EndDate        *string `json:"endDate"`
	Months         *string `json:"months"`
	CreatedAt      *string `json:"createdAt"`
	UpdatedAt      *string `json:"updatedAt"`
}

type Phone struct {
	ID                  int     `json:"id"`
	Type                *string `json:"type"`
	Number              *string `json:"number"`
	Formatted           *string `json:"formatted"`
	InternationalMobile *string `json:"internationalMobile"`
	Primary             bool    `json:"primary"`
	Inactive            bool    `json:"inactive"`
	IsMobile            bool    `json:"isMobile"`
	AllowTelemarketing  *bool   `json:"allowTelemarketing"`
	Country             *string `json:"country"`
	CreatedAt           *string `json:"createdAt"`
	UpdatedAt           *string `json:"updatedAt"`
}

type Salutation struct {
	ID                  int     `json:"id"`
	Type                *string `json:"type"`
	BusinessTitle       *string `json:"businessTitle"`
	EnvelopeSalutation1 *string `json:"envelopeSalutation1"`
	EnvelopeSalutation2 *string `json:"envelopeSalutation2"`
	LetterSalutation    *string `json:"letterSalutation"`
	Default             bool    `json:"default"`
	CreatedAt           *string `json:"createdAt"`
	UpdatedAt           *string `json:"updatedAt"`
}

type Affiliation struct {
	ID               int     `json:"id"`
	Type             *string `json:"type"`
	GroupName        *string `json:"groupName,omitempty"`
	GroupID          *int    `json:"groupId,omitempty"`
	IndividualName   *string `json:"individualName,omitempty"`
	IndividualID     *int    `json:"individualId,omitempty"`
	Title            *string `json:"title"`
	Primary          bool    `json:"primary"`
	Inactive         bool    `json:"inactive"`
	StartDate        *string `json:"startDate"`
	EndDate          *string `json:"endDate"`
	Note             *string `json:"note"`
	CreatedAt        *string `json:"createdAt"`
	UpdatedAt        *string `json:"updatedAt"`
}

// refDesc extracts the Description string from an API ref object.
func refDesc(r *tessitura.APIRef) *string {
	if r == nil || r.Description == nil {
		return nil
	}
	return r.Description
}

// ConstituentFromAPI maps a raw Tessitura API response to a clean domain Constituent.
func ConstituentFromAPI(d *tessitura.APIConstituentDetail) *Constituent {
	c := Constituent{
		ID:               derefInt(d.Id),
		FirstName:        d.FirstName,
		MiddleName:       d.MiddleName,
		LastName:         d.LastName,
		DisplayName:      d.DisplayName,
		SortName:         d.SortName,
		Prefix:           refDesc(d.Prefix),
		Suffix:           refDesc(d.Suffix),
		Gender:           refDesc(d.Gender),
		Pronoun:          refDesc(d.Pronoun),
		Inactive:         isInactive(d.Inactive),
		InactiveReason:   refDesc(d.InactiveReason),
		ProtectionType:   refDesc(d.ProtectionType),
		OriginalSource:   refDesc(d.OriginalSource),
		NameStatus:       refDesc(d.NameStatus),
		MailIndicator:    refDesc(d.MailIndicator),
		EmailIndicator:   refDesc(d.EmarketIndicator),
		PhoneIndicator:   refDesc(d.PhoneIndicator),
		LastActivityDate: d.LastActivityDate,
		LastGiftDate:     d.LastGiftDate,
		LastTicketDate:   d.LastTicketDate,
		CreatedAt:        d.CreatedDateTime,
		CreatedBy:        d.CreatedBy,
		UpdatedAt:        d.UpdatedDateTime,
		UpdatedBy:        d.UpdatedBy,
	}

	if d.ConstituentType != nil {
		c.ConstituentType = d.ConstituentType.Description
		if d.ConstituentType.ConstituentGroup != nil {
			c.ConstituentGroup = d.ConstituentType.ConstituentGroup.Description
		}
	}

	for _, s := range d.Salutations {
		if s.Constituent == nil || s.Constituent.Id == c.ID {
			c.Salutations = append(c.Salutations, salutationFromAPI(s))
		}
	}

	for _, a := range d.Addresses {
		if a.Constituent == nil || a.Constituent.Id == c.ID {
			c.Addresses = append(c.Addresses, addressFromAPI(a))
		}
	}
	for _, e := range d.ElectronicAddresses {
		if e.Constituent == nil || e.Constituent.Id == c.ID {
			c.Emails = append(c.Emails, emailFromAPI(e))
		}
	}
	for _, p := range d.PhoneNumbers {
		if p.Constituent == nil || p.Constituent.Id == c.ID {
			c.Phones = append(c.Phones, phoneFromAPI(p))
		}
	}

	return &c
}

func addressFromAPI(a tessitura.APIAddress) Address {
	addr := Address{
		ID:         derefInt(a.Id),
		Type:       refDesc(a.AddressType),
		Primary:    derefBool(a.PrimaryIndicator),
		Inactive:   derefBool(a.Inactive),
		Street1:    a.Street1,
		Street2:    a.Street2,
		Street3:    a.Street3,
		City:       a.City,
		PostalCode: a.PostalCode,
		StartDate:  a.StartDate,
		EndDate:    a.EndDate,
		Months:     a.Months,
		CreatedAt:  a.CreatedDateTime,
		UpdatedAt:  a.UpdatedDateTime,
	}
	if a.State != nil {
		addr.State = a.State.Description
		addr.StateCode = a.State.StateCode
	}
	if a.Country != nil {
		addr.Country = a.Country.Description
		addr.CountryCode = a.Country.IsoAlpha2
	}
	return addr
}

func emailFromAPI(e tessitura.APIElectronicAddress) Email {
	return Email{
		ID:             derefInt(e.Id),
		Type:           refDesc(e.ElectronicAddressType),
		Address:        e.Address,
		Primary:        derefBool(e.PrimaryIndicator),
		Inactive:       derefBool(e.Inactive),
		IsEmail:        derefBool(e.IsEmail),
		AllowHtml:      e.AllowHtmlFormat,
		AllowMarketing: derefBool(e.AllowMarketing),
		StartDate:      e.StartDate,
		EndDate:        e.EndDate,
		Months:         e.Months,
		CreatedAt:      e.CreatedDateTime,
		UpdatedAt:      e.UpdatedDateTime,
	}
}

func phoneFromAPI(p tessitura.APIPhone) Phone {
	ph := Phone{
		ID:                  derefInt(p.Id),
		Type:                refDesc(p.PhoneType),
		Number:              p.PhoneNumber,
		Formatted:           p.PhoneFormatted,
		InternationalMobile: p.InternationalMobile,
		Primary:             derefBool(p.PrimaryIndicator),
		Inactive:            derefBool(p.Inactive),
		IsMobile:            derefBool(p.IsMobile),
		AllowTelemarketing:  p.AllowTelemarketing,
		CreatedAt:           p.CreatedDateTime,
		UpdatedAt:           p.UpdatedDateTime,
	}
	if p.Country != nil {
		ph.Country = p.Country.Description
	}
	return ph
}

func derefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func derefBool(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}

func salutationFromAPI(s tessitura.APISalutation) Salutation {
	return Salutation{
		ID:                  derefInt(s.Id),
		Type:                refDesc(s.SalutationType),
		BusinessTitle:       s.BusinessTitle,
		EnvelopeSalutation1: s.EnvelopeSalutation1,
		EnvelopeSalutation2: s.EnvelopeSalutation2,
		LetterSalutation:    s.LetterSalutation,
		Default:             s.DefaultIndicator,
		CreatedAt:           s.CreatedDateTime,
		UpdatedAt:           s.UpdatedDateTime,
	}
}

func affiliationFromAPI(a tessitura.APIAffiliation, parentID int) Affiliation {
	aff := Affiliation{
		ID:        derefInt(a.Id),
		Title:     a.Title,
		Primary:   a.PrimaryIndicator,
		Inactive:  derefBool(a.Inactive),
		StartDate: a.StartDate,
		EndDate:   a.EndDate,
		Note:      a.Note,
		CreatedAt: a.CreatedDateTime,
		UpdatedAt: a.UpdatedDateTime,
	}
	if a.AffiliationType != nil {
		aff.Type = a.AffiliationType.Description
	}
	// Show the "other" side of the relationship.
	if a.GroupConstituent != nil && a.GroupConstituent.Id != parentID {
		id := a.GroupConstituent.Id
		aff.GroupID = &id
		aff.GroupName = a.GroupConstituentName
	}
	if a.IndividualConstituent != nil && a.IndividualConstituent.Id != parentID {
		id := a.IndividualConstituent.Id
		aff.IndividualID = &id
		aff.IndividualName = a.IndividualConstituentName
	}
	return aff
}

// AttachAffiliations maps raw API affiliations and attaches them to the constituent.
// Only the "other" side of each affiliation is shown (group side for individuals,
// individual side for households/orgs).
func (c *Constituent) AttachAffiliations(apiAffiliations []tessitura.APIAffiliation) {
	for _, a := range apiAffiliations {
		c.Affiliations = append(c.Affiliations, affiliationFromAPI(a, c.ID))
	}
}

type Note struct {
	ID        int     `json:"id"`
	Type      *string `json:"type"`
	Text      *string `json:"text"`
	CreatedAt *string `json:"createdAt"`
	CreatedBy *string `json:"createdBy"`
	UpdatedAt *string `json:"updatedAt"`
	UpdatedBy *string `json:"updatedBy"`
}

func noteFromAPI(n tessitura.APINote) Note {
	note := Note{
		ID:        derefInt(n.Id),
		Text:      n.Note,
		CreatedAt: n.CreatedDateTime,
		CreatedBy: n.CreatedBy,
		UpdatedAt: n.UpdatedDateTime,
		UpdatedBy: n.UpdatedBy,
	}
	if n.NoteType != nil {
		note.Type = n.NoteType.Description
	}
	return note
}

type Association struct {
	ID             int     `json:"id"`
	Type           *string `json:"type"`
	AssociatedID   *int    `json:"associatedId,omitempty"`
	AssociatedName *string `json:"associatedName,omitempty"`
	Gender         *string `json:"gender,omitempty"`
	BirthDate      *string `json:"birthDate,omitempty"`
	Inactive       bool    `json:"inactive"`
	StartDate      *string `json:"startDate,omitempty"`
	EndDate        *string `json:"endDate,omitempty"`
	Note           *string `json:"note,omitempty"`
	CreatedAt      *string `json:"createdAt"`
	CreatedBy      *string `json:"createdBy"`
	UpdatedAt      *string `json:"updatedAt"`
	UpdatedBy      *string `json:"updatedBy"`
}

func associationFromAPI(a tessitura.APIAssociation) Association {
	assoc := Association{
		ID:             derefInt(a.Id),
		AssociatedName: a.AssociatedName,
		BirthDate:      a.BirthDate,
		Inactive:       derefBool(a.Inactive),
		StartDate:      a.StartDate,
		EndDate:        a.EndDate,
		Note:           a.Note,
		CreatedAt:      a.CreatedDateTime,
		CreatedBy:      a.CreatedBy,
		UpdatedAt:      a.UpdatedDateTime,
		UpdatedBy:      a.UpdatedBy,
	}
	if a.AssociationType != nil {
		assoc.Type = a.AssociationType.Description
	}
	if a.AssociatedConstituent != nil {
		id := a.AssociatedConstituent.Id
		assoc.AssociatedID = &id
	}
	if a.Gender != nil {
		assoc.Gender = a.Gender.Description
	}
	return assoc
}

// AttachAssociations maps raw API associations and attaches them to the constituent.
func (c *Constituent) AttachAssociations(apiAssociations []tessitura.APIAssociation) {
	for _, a := range apiAssociations {
		c.Associations = append(c.Associations, associationFromAPI(a))
	}
}

// AttachNotes maps raw API notes and attaches them to the constituent.
func (c *Constituent) AttachNotes(apiNotes []tessitura.APINote) {
	for _, n := range apiNotes {
		c.Notes = append(c.Notes, noteFromAPI(n))
	}
}

// isInactive reads the Inactive boolean from the ConstituentInactiveSummary ref.
func isInactive(r *tessitura.APIInactiveSummary) bool {
	if r == nil || r.Inactive == nil {
		return false
	}
	return *r.Inactive
}
