# Durb

A Go CLI client for the Tessitura API, built for the Folger Shakespeare Library.

**Name origin:** "Durb" is short for d'Urbervilles — a play on *Tess of the d'Urbervilles* matching the Tessitura/Tess name. The CLI command is `tess` (not `durb`).

## Architecture

Cobra CLI pattern (like SF CLI / AWS CLI / Twilio CLI).

**Module path:** `github.com/Folger-Shakespeare-Library/durb`

### Directory structure

- `cmd/tess/main.go` — entrypoint; sets version via ldflags
- `internal/cli/` — cobra command definitions
  - `root.go` — root `tess` command; registers `config`, `crm`, `ref`, `report`
  - `config.go` — `tess config` subcommand group (init, show, path, use, list)
  - `crm.go` — `tess crm` subcommand group; registers activity, attribute, constituent, interest
  - `constituent.go` — `tess crm constituent` subcommand group (alias: `con`); registers get/search/create/update/set-status
  - `constituent_get.go` — `tess crm constituent get`
  - `constituent_search.go` — `tess crm constituent search`
  - `constituent_create.go` — `tess crm constituent create`
  - `constituent_update.go` — `tess crm constituent update`
  - `constituent_set_status.go` — `tess crm constituent set-status`
  - `activity.go` — `tess crm activity` subcommand group; registers create/delete/list
  - `activity_create.go` — `tess crm activity create`
  - `activity_delete.go` — `tess crm activity delete`
  - `activity_list.go` — `tess crm activity list`
  - `attribute.go` — `tess crm attribute` subcommand group; registers set/delete/list
  - `attribute_set.go` — `tess crm attribute set`
  - `attribute_delete.go` — `tess crm attribute delete`
  - `attribute_list.go` — `tess crm attribute list`
  - `interest.go` — `tess crm interest` subcommand group; registers enable/disable/list
  - `interest_enable.go` — `tess crm interest enable`
  - `interest_disable.go` — `tess crm interest disable`
  - `interest_list.go` — `tess crm interest list`
  - `ref.go` — `tess ref` subcommand group; registers all ref subcommands
  - `ref_activity_statuses.go` — `tess ref activity-statuses list`
  - `ref_activity_types.go` — `tess ref activity-types list`
  - `ref_address_types.go` — `tess ref address-types list`
  - `ref_affiliation_types.go` — `tess ref affiliation-types list`
  - `ref_alias_types.go` — `tess ref alias-types list`
  - `ref_appeal_categories.go` — `tess ref appeal-categories list`
  - `ref_association_types.go` — `tess ref association-types list`
  - `ref_business_units.go` — `tess ref business-units list`
  - `ref_campaign_categories.go` — `tess ref campaign-categories list`
  - `ref_constituency_types.go` — `tess ref constituency-types list`
  - `ref_constituent_groups.go` — `tess ref constituent-groups list`
  - `ref_constituent_inactives.go` — `tess ref constituent-inactives list`
  - `ref_constituent_types.go` — `tess ref constituent-types list`
  - `ref_contact_permission_categories.go` — `tess ref contact-permission-categories list`
  - `ref_contact_permission_types.go` — `tess ref contact-permission-types list`
  - `ref_contribution_designations.go` — `tess ref contribution-designations list`
  - `ref_contribution_import_sets.go` — `tess ref contribution-import-sets list`
  - `ref_countries.go` — `tess ref countries list`
  - `ref_delivery_methods.go` — `tess ref delivery-methods list`
  - `ref_designation_codes.go` — `tess ref designation-codes list`
  - `ref_donation_levels.go` — `tess ref donation-levels list`
  - `ref_electronic_address_types.go` — `tess ref electronic-address-types list`
  - `ref_genders.go` — `tess ref genders list`
  - `ref_inactive_reasons.go` — `tess ref inactive-reasons list`
  - `ref_interest_categories.go` — `tess ref interest-categories list`
  - `ref_interest_types.go` — `tess ref interest-types list`
  - `ref_keyword_categories.go` — `tess ref keyword-categories list`
  - `ref_keywords.go` — `tess ref keywords list`
  - `ref_languages.go` — `tess ref languages list`
  - `ref_login_types.go` — `tess ref login-types list`
  - `ref_machine_settings.go` — `tess ref machine-settings list`
  - `ref_membership_benefit_frequencies.go` — `tess ref membership-benefit-frequencies list`
  - `ref_membership_benefit_types.go` — `tess ref membership-benefit-types list`
  - `ref_membership_level_categories.go` — `tess ref membership-level-categories list`
  - `ref_membership_level_trends.go` — `tess ref membership-level-trends list`
  - `ref_membership_periods.go` — `tess ref membership-periods list`
  - `ref_membership_standings.go` — `tess ref membership-standings list`
  - `ref_membership_statuses.go` — `tess ref membership-statuses list`
  - `ref_note_types.go` — `tess ref note-types list`
  - `ref_order_categories.go` — `tess ref order-categories list`
  - `ref_original_sources.go` — `tess ref original-sources list`
  - `ref_payment_types.go` — `tess ref payment-types list`
  - `ref_performance_statuses.go` — `tess ref performance-statuses list`
  - `ref_performance_types.go` — `tess ref performance-types list`
  - `ref_philanthropy_types.go` — `tess ref philanthropy-types list`
  - `ref_plan_priorities.go` — `tess ref plan-priorities list`
  - `ref_plan_sources.go` — `tess ref plan-sources list`
  - `ref_plan_statuses.go` — `tess ref plan-statuses list`
  - `ref_plan_types.go` — `tess ref plan-types list`
  - `ref_planned_giving_codes.go` — `tess ref planned-giving-codes list`
  - `ref_planned_giving_fundings.go` — `tess ref planned-giving-fundings list`
  - `ref_planned_giving_gift_types.go` — `tess ref planned-giving-gift-types list`
  - `ref_planned_giving_on_files.go` — `tess ref planned-giving-on-files list`
  - `ref_planned_giving_purposes.go` — `tess ref planned-giving-purposes list`
  - `ref_planned_giving_sources.go` — `tess ref planned-giving-sources list`
  - `ref_planned_giving_statuses.go` — `tess ref planned-giving-statuses list`
  - `ref_prefixes.go` — `tess ref prefixes list`
  - `ref_price_categories.go` — `tess ref price-categories list`
  - `ref_price_type_categories.go` — `tess ref price-type-categories list`
  - `ref_price_type_groups.go` — `tess ref price-type-groups list`
  - `ref_pronouns.go` — `tess ref pronouns list`
  - `ref_recognition_types.go` — `tess ref recognition-types list`
  - `ref_report_categories.go` — `tess ref report-categories list`
  - `ref_report_types.go` — `tess ref report-types list`
  - `ref_sales_channels.go` — `tess ref sales-channels list`
  - `ref_seat_codes.go` — `tess ref seat-codes list`
  - `ref_seat_statuses.go` — `tess ref seat-statuses list`
  - `ref_seasons.go` — `tess ref seasons list`
  - `ref_sections.go` — `tess ref sections list`
  - `ref_states.go` — `tess ref states list`
  - `ref_suffixes.go` — `tess ref suffixes list`
  - `ref_theaters.go` — `tess ref theaters list`
  - `ref_user_groups.go` — `tess ref user-groups list`
  - `report.go` — `tess report` subcommand group; registers get/list/request
  - `report_get.go` — `tess report get`
  - `report_list.go` — `tess report list`
  - `report_request.go` — `tess report request` subcommand group; registers get/list/results
  - `report_request_get.go` — `tess report request get`
  - `report_request_list.go` — `tess report request list`
  - `report_request_results.go` — `tess report request results`
- `pkg/tessitura/` — raw Tessitura API client and response structs (mirrors the JSON shape; do not use directly in consumer code)
  - `client.go` — HTTP client, auth, `Get`/`Post`/`Put`/`Delete`/`Batch` methods
  - `constituents.go` — `ConstituentResult`, `CreateConstituentParams`, `GetConstituentDetail`, `GetConstituentFull`, `GetConstituentsBatch`, `CreateConstituent`, `SetConstituentStatus`
  - `search.go` — constituent search types and `SearchConstituents`
  - `affiliations.go` — affiliation structs and `GetAffiliations`
  - `associations.go` — association structs
  - `aliases.go` — alias structs
  - `logins.go` — web login structs
  - `notes.go` — note structs
  - `activities.go` — `APISpecialActivity`, `CreateActivityParams`, `GetActivities`, `CreateActivity`, `DeleteActivity`
  - `attributes.go` — `APIAttribute`, `GetAttributes`, `CreateAttribute`, `UpdateAttribute`, `DeleteAttribute`
  - `interests.go` — `APIInterest`, `GetInterests`, `CreateInterest`, `UpdateInterest`
  - `electronic_addresses.go` — `GetElectronicAddresses`, `UpdateElectronicAddress`
  - `reference.go` — all reference data API structs and getter functions; types include `APIRefItem`, `APISeatStatus`, `APIMachineSetting`, `APIConstituentType`, `APIOriginalSource`, `APISpecialActivityType`, `APISpecialActivityStatus`, `APIKeyword`, `APIInterestType`, `APIInactiveReason`, `APIControlGroupRef`, `APIRelationshipCategoryRef`, `APIReportCategory`, `APIReportType`, `APISeason`, `APIPerformanceStatus`, `APIPerformanceType`, `APIPriceCategory`, `APIPriceTypeCategory`, `APIPriceTypeGroup`, `APIPaymentType`, `APIDeliveryMethod`, `APIOrderCategory`, `APISalesChannel`, `APINoteType`, `APIElectronicAddressType`, `APIContactPermissionCategory`, `APIContactPermissionType`, `APIAffiliationType`, `APIAssociationType`, `APIAliasType`, `APILoginType`, `APIGender`, `APIPronoun`, `APIPrefix`, `APISuffix`, `APICountry`, `APIState`, `APIAddressType`, `APILanguage`, `APIConstituencyType`, `APIConstituentGroup`, `APIKeywordCategory`, `APIInterestCategory`, `APIUserGroup`, `APIBusinessUnit`, `APITheater`, `APISection`, `APISeatCode`, `APIAppealCategory`, `APICampaignCategory`, `APIContributionDesignation`, `APIContributionImportSet`, `APIDesignationCode`, `APIDonationLevel`, `APIPhilanthropyType`, `APIPlanPriority`, `APIPlanSource`, `APIPlanStatus`, `APIPlanType`, `APIPlannedGivingCode`, `APIPlannedGivingFunding`, `APIPlannedGivingGiftType`, `APIPlannedGivingOnFile`, `APIPlannedGivingPurpose`, `APIPlannedGivingSource`, `APIPlannedGivingStatus`, `APIRecognitionType`, `APIMembershipBenefitFrequency`, `APIMembershipBenefitType`, `APIMembershipLevelCategory`, `APIMembershipLevelTrend`, `APIMembershipPeriod`, `APIMembershipStanding`, `APIMembershipStatus`
  - `reports.go` — `APIReport`, `APIReportDetail`, `APIReportParameter`, `ReportResult`; `GetReports`, `GetReport`, `GetReportsBatch`
  - `report_requests.go` — `APIReportRequest`, `APIReportRequestDetail`, `APIReportResult`, `ReportRequestResult`, `ReportResultsParams`; `GetReportRequests`, `GetReportRequest`, `GetReportRequestsBatch`, `GetReportResults`
- `pkg/domain/` — clean domain types mapped from raw API responses (all consumer code uses these)
  - `constituent.go` — `Constituent` and sub-types (`Address`, `Email`, `DigitalAddress`, `Phone`, `Salutation`, `Affiliation`, `Note`, `Association`, `Login`, `Alias`); `ConstituentFromAPI` and `Attach*` methods
  - `search.go` — `ConstituentSearchResult`; `SearchResultsFromAPI`
  - `report.go` — `Report`, `ReportRef`, `ReportParameter`; `ReportFromAPI` and `AttachDetail`
  - `report_request.go` — `ReportRequest`, `ReportRequestParameter`, `ReportResult`, `ReportResultRef`, `ReportResultReportRef`; `ReportRequestFromAPI`, `AttachRequestDetail`, `ReportResultFromAPI`
- `pkg/config/` — config management; config path is `$XDG_CONFIG_HOME/tess/config.json` (defaults to `~/.config/tess/config.json`)
- `schemas/constituent.schema.json` — JSON Schema for the `Constituent` domain object (**must be updated when domain fields change**)
- `swagger.json` — Tessitura API swagger file (v16.0.27.97921); OpenAPI 3.0 format — model schemas are under `components.schemas` (not `definitions`), and response `$ref` paths use `#/components/schemas/ModelName`

## Swagger lookup recipe

To get API struct fields for a reference data endpoint:

```python
import json
with open('swagger.json') as f:
    swagger = json.load(f)
model = swagger['components']['schemas']['ModelName']  # e.g. 'ConstituentProtectionType'
for name, prop in model.get('properties', {}).items():
    print(name, prop.get('type', prop.get('$ref', '').split('/')[-1]), 'nullable' if prop.get('nullable') else '')
```

Common sub-object refs resolve to other `components.schemas` entries (e.g. `ControlGroupSummary`, `RefItem`). Map these to existing Go types like `APIControlGroupRef`, `APIRefItem`.

## Adding a `tess ref` command

Each `tess ref <resource> list` command needs three things:

1. **API struct + getter** in `pkg/tessitura/reference.go` — struct mirrors the JSON shape from swagger; getter calls `c.Get(ctx, "/ReferenceData/<Resource>")`, unmarshals into a slice, returns it. Use the full endpoint (not `/Summary`) per the "full records" rule.
2. **CLI file** `internal/cli/ref_<resource>.go` — defines the cobra command vars and `runRef<Resource>List` function. Follow the pattern in any existing `ref_*.go` file.
3. **Registration** in `internal/cli/ref.go` — add `refCmd.AddCommand(ref<Resource>Cmd)` in `init()`, alphabetically.

After adding, update `CLAUDE.md` (directory listing, type list in `reference.go` entry, implemented commands section) and `README.md`.

## Standing orders

- When a new command is added, update both CLAUDE.md and README.md to reflect it.

## Key design decisions

- **Domain objects over table endpoints:** The Tessitura API maps to database tables. Durb adds a domain-object layer (`pkg/domain/`) that consolidates related endpoints — e.g., a `Constituent` that folds in addresses, emails, phones from multiple API calls. Raw API types live in `pkg/tessitura/`; all consumer code uses `pkg/domain/` types.
- **Constituent first:** Primary domain object. Expand to others as needed.
- **JSON output only** for now.
- **Full records, no summaries:** Use the full endpoints, not `/Summary` variants. Always return the complete API response — do not filter, abbreviate, or reshape fields. Consumers pipe through `jq`.
- **Auth:** Tessitura uses 4-part basic auth: base64(`username:usergroup:location:password`). Config supports multiple named profiles.
- **Profile resolution:** `--profile` flag > `TESSITURA_PROFILE` env var > `default_profile` in config > `"default"`. Individual fields can be overridden with `TESSITURA_HOSTNAME`, `TESSITURA_USERNAME`, `TESSITURA_USER_GROUP`, `TESSITURA_LOCATION`, `TESSITURA_PASSWORD`.
- **Config file permissions:** On non-Windows, enforces `0600` permissions on the config file.
- **`--with` flag** on `constituent get` for optional sub-objects (affiliations, associations, aliases, logins, notes). Addresses, emails, phones, and salutations are always included via the `/Detail` endpoint. Use `--with all` to enable everything optional.
- **Batching:** Used in two ways:
  - *Constituents:* When any `--with` extras are requested, `GetConstituentFull` uses the Tessitura `/api/Batch` endpoint to fetch detail + all extras in a single HTTP call. Fixed request IDs: 1=detail, 2=affiliations (individual), 3=affiliations (group), 4=notes, 5=associations, 6=logins, 7=aliases. When adding new `--with` options, assign the next available ID and add a case in the switch.
  - *Reports:* `GetReport` always batches request IDs 1=base (`/Reporting/Reports/{id}`) and 2=detail (`/Reporting/Reports/{id}/Details`). This is necessary because `AllowQuery` and `QueryStringAppend` only exist on the base endpoint while `Parameters` only exists on the detail endpoint.
- **Concurrent fetching:** Multiple IDs are fetched with goroutines in `GetConstituentsBatch` and `GetReportsBatch`.
- **Optimistic locking** for updates: Tessitura requires passing `UpdatedDateTime` back on PUT operations. Used in `constituent update`, `attribute set`, and `interest enable`/`disable`.
- **Field length limits:** The API silently truncates certain fields (not documented in the swagger). Known limits: first name 20 chars, last name 55 chars, postal code 10 chars. `constituent search` warns on stderr when values exceed these limits. `constituent create` returns a hard error.
- **`report list` active-only by default:** The API has no server-side inactive filter, so `report list` filters client-side. Active reports are shown by default; use `--include-inactive` to see all.
- **`report request list` active-only by default:** The API supports `activeOnly` as a server-side query param, so `report request list` passes it directly. Active requests only by default; use `--include-inactive` to see completed/cancelled requests.
- **`report request results` pagination:** The Results endpoint is paginated. Default page size is 100 (matching the API default). If more results exist, the count of remaining results is printed to stderr. Use `--page` to paginate.

## Implemented commands

### `tess config init`
Interactive prompt to set API credentials for a named profile. Requires `--profile`.
Saves to `$XDG_CONFIG_HOME/tess/config.json` (default `~/.config/tess/config.json`).

### `tess config show`
Displays the active profile's configuration (password masked).

### `tess config path`
Prints the config file path.

### `tess config use <profile>`
Sets the default profile.

### `tess config list`
Lists all configured profiles. Active profile is marked with `*`.

### `tess crm constituent get <id> [id...]`
Fetches one or more constituents by ID. Always returns a JSON array.
- `--with affiliations` — attach affiliations (org/household memberships)
- `--with associations` — attach associations
- `--with aliases` — attach aliases
- `--with logins` — attach web logins
- `--with notes` — attach constituent notes
- `--with all` — attach all optional data
- Reads IDs from stdin (one per line) if piped

### `tess crm constituent search [query]`
Searches constituents. Returns a JSON array of summary records.
- Free-text: positional arg(s)
- Basic structured: `--last-name`, `--first-name`, `--street`, `--postal-code`, `--id`
- Advanced (one at a time): `--email`, `--phone`, `--order-no`, `--web-login`, `--customer-service-no`
- Operator: `--op` (Equals, Like, LessThan, GreaterThan) — applies to advanced search only
- Filters: `--groups` (comma-separated: individuals, organizations, households), `--include-affiliations`
- Warns on stderr when `--first-name` > 20, `--last-name` > 55, or `--postal-code` > 10 chars (API truncates silently)

### `tess crm constituent create`
Creates a new constituent. Returns the created record as JSON.
- `--first` — first name (required, max 20 chars)
- `--last` — last name (required, max 55 chars)
- `--email` — email address (required)
- `--constituent-type-id` — constituent type ID (required)
- `--original-source-id` — original source ID (required)
- `--street` — street address
- `--postal-code` — postal/ZIP code (max 10 chars)
- `--allow-marketing` — allow email marketing
- Auto-generates SortName as "LastName, FirstName"
- Returns a hard error if field length limits are exceeded

### `tess crm constituent update`
Updates properties on an existing constituent. Currently supports the email marketing flag only.
- `--id` — constituent ID (required)
- `--email` — email address to update (required)
- `--allow-marketing` / `--allow-marketing=false`
- Uses optimistic locking via `UpdatedDateTime`

### `tess crm constituent set-status`
Changes a constituent's inactive status. Looks up status and reason by description from reference data.
- `--id` — constituent ID (required)
- `--status` — status description, e.g. "Active", "Inactive" (required)
- `--reason` — inactive reason description (required when inactivating)

### `tess crm activity create`
Creates a special activity record on a constituent.
- `--constituent-id` — constituent ID (required)
- `--activity-type-id` — activity type ID (required)
- `--status-id` — activity status ID (required)
- `--datetime` — ISO 8601 date or datetime (required); bare dates default to midnight local time
- `--notes` — activity notes
- `--unique` — skip creation if a matching activity (same type + datetime) already exists; returns the existing record

### `tess crm activity delete`
Deletes an activity record by ID.
- `--activity-id` — activity record ID (required)

### `tess crm activity list`
Lists activities for a constituent.
- `--constituent-id` — constituent ID (required)
- `--activity-type-id` — filter by activity type ID

### `tess crm attribute set`
Sets an attribute on a constituent. Creates the attribute if none exists for that keyword; updates the value if one already exists.
- `--constituent-id` — constituent ID (required)
- `--attribute-type-id` — attribute type / keyword ID (required)
- `--value` — attribute value (required)
- Uses optimistic locking on update

### `tess crm attribute delete`
Deletes an attribute record by ID.
- `--attribute-id` — attribute record ID (required)

### `tess crm attribute list`
Lists attributes on a constituent.
- `--constituent-id` — constituent ID (required)
- `--attribute-type-id` — filter by attribute type ID

### `tess crm interest enable`
Enables one or more interests on a constituent. Creates the interest assignment if it doesn't exist; updates it if it does.
- `--constituent-id` — constituent ID (required)
- `--interest-type-ids` — comma-separated interest type IDs (required)
- Uses optimistic locking on update

### `tess crm interest disable`
Disables one or more interests on a constituent.
- `--constituent-id` — constituent ID (required)
- `--interest-type-ids` — comma-separated interest type IDs (required)

### `tess crm interest list`
Lists all interest assignments for a constituent.
- `--constituent-id` — constituent ID (required)

### `tess ref activity-statuses list`
Lists available special activity statuses. Used with `activity create --status-id`.

### `tess ref activity-types list`
Lists available special activity types. Used with `activity create --activity-type-id`.

### `tess ref address-types list`
Lists available address types.

### `tess ref affiliation-types list`
Lists available affiliation types for classifying constituent affiliations.

### `tess ref alias-types list`
Lists available alias types for classifying constituent aliases.

### `tess ref appeal-categories list`
Lists available appeal categories for fundraising appeals.

### `tess ref association-types list`
Lists available association types for classifying constituent associations.

### `tess ref business-units list`
Lists available business units.

### `tess ref campaign-categories list`
Lists available campaign categories for fundraising campaigns.

### `tess ref constituency-types list`
Lists available constituency types for constituent classification.

### `tess ref constituent-groups list`
Lists available constituent groups (individuals, organizations, households).

### `tess ref constituent-inactives list`
Lists available inactive status types for constituents.

### `tess ref constituent-types list`
Lists available constituent types. Used with `constituent create --constituent-type-id`.

### `tess ref contact-permission-categories list`
Lists available contact permission categories for marketing consent groupings.

### `tess ref contact-permission-types list`
Lists available contact permission types for marketing consent configuration.

### `tess ref contribution-designations list`
Lists available contribution designation codes for directing contributions.

### `tess ref contribution-import-sets list`
Lists import set definitions for contribution imports.

### `tess ref countries list`
Lists available countries.

### `tess ref delivery-methods list`
Lists available ticket delivery methods.

### `tess ref designation-codes list`
Lists available codes for directing contributions.

### `tess ref donation-levels list`
Lists available giving levels/tiers with recognition type references.

### `tess ref electronic-address-types list`
Lists available electronic address types (email, phone, web type IDs).

### `tess ref genders list`
Lists available genders with default prefix and pronoun references.

### `tess ref inactive-reasons list`
Lists available inactive reasons for constituents. Used with `constituent set-status --reason`.

### `tess ref interest-categories list`
Lists available interest categories (metadata for interest types).

### `tess ref interest-types list`
Lists available interest types. Used with `interest enable/disable --interest-type-ids`.

### `tess ref keyword-categories list`
Lists available keyword categories (metadata for attribute types).

### `tess ref keywords list`
Lists available keywords (attribute types in Tessitura). Used with `attribute set --attribute-type-id`.

### `tess ref languages list`
Lists available language codes.

### `tess ref login-types list`
Lists available login types for classifying constituent web logins.

### `tess ref machine-settings list`
Lists machine settings (workstation name, card reader configuration, merchant IDs, audit fields). Uses the full `/ReferenceData/MachineSettings` endpoint (not Summary, which returns empty descriptions).

### `tess ref membership-benefit-frequencies list`
Lists available frequency options for membership benefits.

### `tess ref membership-benefit-types list`
Lists available types of membership benefits.

### `tess ref membership-level-categories list`
Lists available categories for membership levels.

### `tess ref membership-level-trends list`
Lists available trend classifications for membership levels.

### `tess ref membership-periods list`
Lists available membership period definitions.

### `tess ref membership-standings list`
Lists available standing types for memberships.

### `tess ref membership-statuses list`
Lists available status values for memberships.

### `tess ref note-types list`
Lists available note types for classifying constituent notes.

### `tess ref order-categories list`
Lists available order categories for order classification.

### `tess ref original-sources list`
Lists available original sources. Used with `constituent create --original-source-id`.

### `tess ref payment-types list`
Lists available payment types for payment method lookups.

### `tess ref performance-statuses list`
Lists available performance statuses for event classification.

### `tess ref performance-types list`
Lists available performance types for event classification.

### `tess ref philanthropy-types list`
Lists available types of philanthropic activity.

### `tess ref plan-priorities list`
Lists available priority levels for development plans.

### `tess ref plan-sources list`
Lists available sources for development plans.

### `tess ref plan-statuses list`
Lists available status values for development plans.

### `tess ref plan-types list`
Lists available types of development/solicitation plans.

### `tess ref planned-giving-codes list`
Lists available codes for planned giving.

### `tess ref planned-giving-fundings list`
Lists available funding types for planned gifts.

### `tess ref planned-giving-gift-types list`
Lists available types of planned gifts.

### `tess ref planned-giving-on-files list`
Lists available on-file statuses for planned gifts.

### `tess ref planned-giving-purposes list`
Lists available purposes for planned gifts.

### `tess ref planned-giving-sources list`
Lists available sources for planned gifts.

### `tess ref planned-giving-statuses list`
Lists available statuses for planned gifts.

### `tess ref prefixes list`
Lists available name prefixes (Mr., Mrs., Dr., etc.).

### `tess ref price-categories list`
Lists available price categories.

### `tess ref price-type-categories list`
Lists available price type categories.

### `tess ref price-type-groups list`
Lists available price type groups.

### `tess ref pronouns list`
Lists available pronoun sets.

### `tess ref recognition-types list`
Lists available donor recognition type definitions.

### `tess ref report-categories list`
Lists available report categories. Used with `report list --category-ids`.

### `tess ref report-types list`
Lists available report types. Used with `report list --type-ids`.

### `tess ref sales-channels list`
Lists available sales channels (box office, online, phone).

### `tess ref seat-codes list`
Lists available seat codes for venue configuration.

### `tess ref seat-statuses list`
Lists available seat statuses.

### `tess ref seasons list`
Lists available seasons for reporting and ticketing.

### `tess ref sections list`
Lists available venue sections.

### `tess ref states list`
Lists available states and provinces with country references.

### `tess ref suffixes list`
Lists available name suffixes (Jr., Sr., III, etc.).

### `tess ref theaters list`
Lists available theaters (venues) with address and capacity data.

### `tess ref user-groups list`
Lists available user groups with division and permission data.

### `tess report get <id> [id...]`
Fetches one or more reports by ID. Always returns a JSON array. Always includes full detail (parameters, indicators) via a batched API call. Multiple IDs are fetched concurrently.
- Reads IDs from stdin (one per line) if piped

The domain `Report` object includes: base fields (`id`, `name`, `description`, `reportPath`, `category`, `reportType`, `allowSchedule`, `allowQuery`, `queryStringAppend`, `parameterWindow`, `parameterWindowIndicator`, `inactive`, `lastRequestId`, audit fields) plus detail-only fields (`publicIndicator`, `warningIndicator`, `utilityIndicator`, `window`, `applicationId`, `parameters`).

`category` and `reportType` are sub-objects: `{"id": 9, "description": "Ticketing Box Office"}`.

### `tess report list`
Lists all reports. Returns a JSON array. Active reports only by default (client-side filter).
- `--type-ids 6` — filter by report type ID (comma-delimited)
- `--category-ids 9` — filter by category ID (comma-delimited)
- `--include-inactive` — include inactive reports

Note: `report list` returns summary records from the base list endpoint only (no parameters or detail-only fields). Use `report get` for the full record.

### `tess report request get <id> [id...]`
Fetches one or more report requests by ID. Always returns a JSON array. Always includes parameter values (from the `/Detail` endpoint) via a batched API call. Multiple IDs are fetched concurrently.
- Reads IDs from stdin (one per line) if piped

### `tess report request list`
Lists all report requests. Returns a JSON array. Active requests only by default (server-side filter via `activeOnly` param).
- `--include-inactive` — include completed and cancelled requests

### `tess report request results`
Lists scheduled report results — a combined entity merging `ReportRequest`, `ReportSchedule`, and `Report` data. Returns a paginated JSON array.
- `--report-id perfseatingbook` — filter by report ID
- `--schedule-name "Daily Seating"` — filter by schedule name
- `--start-date 2025-01-01` / `--end-date 2025-12-31` — date range filter
- `--my-reports-only` — only show results for the current user
- `--recent-only` — only show recent results
- `--include-public` — include public report results
- `--include-errors` — include errored results
- `--include-deleted` — include results whose output has been deleted
- `--page 2` / `--page-size 100` — pagination (default page size: 100)

If more results exist beyond the current page, the remaining count is printed to stderr.

## Building

```bash
go build -o tess ./cmd/tess
```

Releases are handled by GoReleaser via GitHub Actions on tag push.

## Schema maintenance

`schemas/constituent.schema.json` documents the `Constituent` domain object shape. It must be updated manually when fields are added to or removed from `pkg/domain/constituent.go`. It currently does not include the `notes` array.
