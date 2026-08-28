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
  - `ref.go` — `tess ref` subcommand group; registers constituent-inactives, seat-statuses
  - `ref_constituent_inactives.go` — `tess ref constituent-inactives list`
  - `ref_machine_settings.go` — `tess ref machine-settings list`
  - `ref_seat_statuses.go` — `tess ref seat-statuses list`
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
  - `reference.go` — `APIRefItem`, `APISeatStatus`, `APIMachineSetting`, `GetConstituentInactiveStatuses`, `GetConstituentInactiveReasons`, `GetMachineSettings`, `GetSeatStatuses`
  - `reports.go` — `APIReport`, `APIReportDetail`, `APIReportParameter`, `ReportResult`; `GetReports`, `GetReport`, `GetReportsBatch`
  - `report_requests.go` — `APIReportRequest`, `APIReportRequestDetail`, `APIReportResult`, `ReportRequestResult`, `ReportResultsParams`; `GetReportRequests`, `GetReportRequest`, `GetReportRequestsBatch`, `GetReportResults`
- `pkg/domain/` — clean domain types mapped from raw API responses (all consumer code uses these)
  - `constituent.go` — `Constituent` and sub-types (`Address`, `Email`, `DigitalAddress`, `Phone`, `Salutation`, `Affiliation`, `Note`, `Association`, `Login`, `Alias`); `ConstituentFromAPI` and `Attach*` methods
  - `search.go` — `ConstituentSearchResult`; `SearchResultsFromAPI`
  - `report.go` — `Report`, `ReportRef`, `ReportParameter`; `ReportFromAPI` and `AttachDetail`
  - `report_request.go` — `ReportRequest`, `ReportRequestParameter`, `ReportResult`, `ReportResultRef`, `ReportResultReportRef`; `ReportRequestFromAPI`, `AttachRequestDetail`, `ReportResultFromAPI`
- `pkg/config/` — config management; config path is `$XDG_CONFIG_HOME/tess/config.json` (defaults to `~/.config/tess/config.json`)
- `schemas/constituent.schema.json` — JSON Schema for the `Constituent` domain object (**must be updated when domain fields change**)
- `swagger.json` — Tessitura API swagger file (v16.0.27.97921)

## Standing orders

- When a new command is added, update both CLAUDE.md and README.md to reflect it.

## Key design decisions

- **Domain objects over table endpoints:** The Tessitura API maps to database tables. Durb adds a domain-object layer (`pkg/domain/`) that consolidates related endpoints — e.g., a `Constituent` that folds in addresses, emails, phones from multiple API calls. Raw API types live in `pkg/tessitura/`; all consumer code uses `pkg/domain/` types.
- **Constituent first:** Primary domain object. Expand to others as needed.
- **JSON output only** for now.
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

### `tess ref constituent-inactives list`
Lists available inactive status types for constituents.

### `tess ref machine-settings list`
Lists machine settings (workstation name, card reader configuration, merchant IDs, audit fields). Uses the full `/ReferenceData/MachineSettings` endpoint (not Summary, which returns empty descriptions).

### `tess ref seat-statuses list`
Lists available seat statuses.

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
