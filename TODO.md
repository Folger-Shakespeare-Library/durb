# TODO

## Reference data commands

Candidate `tess ref <resource> list` commands from `/ReferenceData/`, ranked by value.

### High — supports existing commands

These are referenced by ID in existing commands. A `ref` lookup avoids guessing IDs.

- [x] ConstituentTypes — `constituent create --constituent-type-id`
- [x] OriginalSources — `constituent create --original-source-id`
- [x] SpecialActivityTypes — `activity create --activity-type-id`
- [x] SpecialActivityStatuses — `activity create --status-id`
- [x] Keywords — `attribute set --attribute-type-id` (keywords = attribute types in Tessitura)
- [x] InterestTypes — `interest enable/disable --interest-type-ids`
- [x] InactiveReasons — used internally by `set-status`; exposing lets users discover valid reason descriptions

### Medium — useful for future commands or common lookups

- [x] ReportCategories — `report list --category-ids`
- [x] ReportTypes — `report list --type-ids`
- [x] Seasons — season references for reporting and ticketing
- [x] PerformanceStatuses / PerformanceTypes — event classification
- [x] PriceCategories / PriceTypeCategories / PriceTypeGroups — pricing
- [x] PaymentTypes — payment method lookups
- [x] DeliveryMethods — ticket delivery options
- [x] OrderCategories — order classification
- [x] SalesChannels — box office vs. online vs. phone
- [x] NoteTypes — classify constituent notes
- [x] ElectronicAddressTypes — email/phone/web type IDs
- [x] ContactPermissionTypes / ContactPermissionCategories — marketing consent
- [x] AffiliationTypes / AssociationTypes / AliasTypes / LoginTypes — classify related records

### Low — rarely needed or installation-specific

- [ ] Genders / Pronouns / Prefixes / Suffixes — demographic data
- [ ] Countries / States / AddressTypes — address lookups
- [ ] Languages — language codes
- [ ] ConstituencyTypes / ConstituentGroups — constituent classification
- [ ] KeywordCategories / InterestCategories — metadata for attribute/interest types
- [ ] UserGroups / BusinessUnits — org structure
- [ ] Theaters / Sections / SeatCodes — venue configuration

Everything else (200+ resources) is deeply administrative or tied to features not yet in the CLI.
