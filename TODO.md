# TODO

- Flatten `report request` to two levels (e.g. `report-request list`) to match CLI best practices (AWS, gh, Stripe all use two-level noun-verb).
- Revisit whether `config init` should suppress password echo during interactive input. Current behavior (visible input) matches AWS/Stripe/Twilio/gh conventions, but worth reconsidering.
- Replace `tess crm constituent set-status` with `enable`/`disable` commands. The Tessitura "constituent inactive" lookup table has four values (Active ID=1, Inactive ID=2, Purge ID=-1, Merged ID=5), but only Active and Inactive are settable through the normal PUT endpoint. Purge requires separate Schedule/Unschedule Purge endpoints, and Merged is destructive (constituent disappears from search). So `set-status` with its `--status` flag is overengineered — in practice you only toggle between Active and Inactive. Proposed: `tess crm constituent enable --id 12345` (sets to Active, no reason allowed) and `tess crm constituent disable --id 12345 --reason "Management request"` (sets to Inactive, reason required). Hardcoding Active=1 and Inactive=2 is safe — these are Tessitura-shipped defaults confirmed in their API docs. The `--reason` flag should look up from `/ReferenceData/InactiveReasons` by description (already implemented in `set-status`). This also means removing Purge and Merged from the valid status list, since they can't be set this way anyway.
- `constituent search --email` results can be misleading. The Tessitura search API matches against all emails on a constituent, but the summary only shows the primary email. Example: searching `--email jane@example.org` may return a constituent whose displayed email is `jane@gmail.com` because `jane@example.org` is a non-primary address that matched but isn't shown. This applies to both `Equals` and `Like` operators. Consider a `--full` flag to fetch complete records for each result, or piping to `constituent get`.

## Reference data commands

Candidate `tess ref <resource> list` commands from `/ReferenceData/`, ranked by value.

### High — supports existing commands

These are referenced by ID in existing commands. A `ref` lookup avoids guessing IDs.

- **ConstituentTypes** — `constituent create --constituent-type-id`
- **OriginalSources** — `constituent create --original-source-id`
- **SpecialActivityTypes** — `activity create --activity-type-id`
- **SpecialActivityStatuses** — `activity create --status-id`
- **Keywords** — `attribute set --attribute-type-id` (keywords = attribute types in Tessitura)
- **InterestTypes** — `interest enable/disable --interest-type-ids`
- **InactiveReasons** — used internally by `set-status`; exposing lets users discover valid reason descriptions

### Medium — useful for future commands or common lookups

- **ReportCategories** — `report list --category-ids`
- **ReportTypes** — `report list --type-ids`
- **Seasons** — season references for reporting and ticketing
- **PerformanceStatuses** / **PerformanceTypes** — event classification
- **PriceCategories** / **PriceTypeCategories** / **PriceTypeGroups** — pricing
- **PaymentTypes** — payment method lookups
- **DeliveryMethods** — ticket delivery options
- **OrderCategories** — order classification
- **SalesChannels** — box office vs. online vs. phone
- **NoteTypes** — classify constituent notes
- **ElectronicAddressTypes** — email/phone/web type IDs
- **ContactPermissionTypes** / **ContactPermissionCategories** — marketing consent
- **AffiliationTypes** / **AssociationTypes** / **AliasTypes** / **LoginTypes** — classify related records

### Low — rarely needed or installation-specific

- **Genders** / **Pronouns** / **Prefixes** / **Suffixes** — demographic data
- **Countries** / **States** / **AddressTypes** — address lookups
- **Languages** — language codes
- **ConstituencyTypes** / **ConstituentGroups** — constituent classification
- **KeywordCategories** / **InterestCategories** — metadata for attribute/interest types
- **UserGroups** / **BusinessUnits** — org structure
- **Theaters** / **Sections** / **SeatCodes** — venue configuration

Everything else (200+ resources) is deeply administrative or tied to features not yet in the CLI.
