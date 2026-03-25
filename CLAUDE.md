# Durb

A Go CLI client for the Tessitura API, built for the Folger Shakespeare Library.

**Name origin:** "Durb" is short for d'Urbervilles — a play on *Tess of the d'Urbervilles* matching the Tessitura/Tess name. The CLI command is `tess` (not `durb`).

## Architecture

Cobra CLI pattern (like SF CLI / AWS CLI / Twilio CLI).

**Module path:** `github.com/Folger-Shakespeare-Library/durb`

### Directory structure

- `cmd/tess/main.go` — entrypoint; sets version via ldflags
- `internal/cli/` — cobra command definitions
  - `root.go` — root `tess` command; registers subcommands
  - `configure.go` — `tess configure`
  - `constituent.go` — `tess constituent` subcommand group; registers get/search
  - `constituent_get.go` — `tess constituent get`
  - `constituent_search.go` — `tess constituent search`
  - `report.go` — `tess report` subcommand group; registers get/list/request
  - `report_get.go` — `tess report get`
  - `report_list.go` — `tess report list`
  - `report_request.go` — `tess report request` subcommand group; registers get/list/results
  - `report_request_get.go` — `tess report request get`
  - `report_request_list.go` — `tess report request list`
  - `report_request_results.go` — `tess report request results`
- `pkg/tessitura/` — raw Tessitura API client and response structs (mirrors the JSON shape; do not use directly in consumer code)
  - `client.go` — HTTP client, auth, `Get`/`Post`/`Batch` methods
  - `constituents.go` — `ConstituentResult`, `GetConstituentFull`, `GetConstituentsBatch`
  - `search.go` — constituent search types and `SearchConstituents`
  - `affiliations.go` — affiliation structs and `GetAffiliations`
  - `notes.go` — note structs
  - `reports.go` — `APIReport`, `APIReportDetail`, `APIReportParameter`, `ReportResult`; `GetReports`, `GetReport`, `GetReportsBatch`
  - `report_requests.go` — `APIReportRequest`, `APIReportRequestDetail`, `APIReportResult`, `ReportRequestResult`, `ReportResultsParams`; `GetReportRequests`, `GetReportRequest`, `GetReportRequestsBatch`, `GetReportResults`
- `pkg/domain/` — clean domain types mapped from raw API responses (all consumer code uses these)
  - `constituent.go` — `Constituent` and sub-types (`Address`, `Email`, `Phone`, `Salutation`, `Affiliation`, `Note`); `ConstituentFromAPI` and `Attach*` methods
  - `search.go` — `ConstituentSearchResult`; `SearchResultsFromAPI`
  - `report.go` — `Report`, `ReportRef`, `ReportParameter`; `ReportFromAPI` and `AttachDetail`
  - `report_request.go` — `ReportRequest`, `ReportRequestParameter`, `ReportResult`, `ReportResultRef`, `ReportResultReportRef`; `ReportRequestFromAPI`, `AttachRequestDetail`, `ReportResultFromAPI`
- `pkg/config/` — config management (`~/.tess/config.json`)
- `schemas/constituent.schema.json` — JSON Schema for the `Constituent` domain object (**must be updated when domain fields change**)
- `docs/` — Tessitura API docs and swagger file (v16.0.27.97921)

## Key design decisions

- **Domain objects over table endpoints:** The Tessitura API maps to database tables. Durb adds a domain-object layer (`pkg/domain/`) that consolidates related endpoints — e.g., a `Constituent` that folds in addresses, emails, phones from multiple API calls. Raw API types live in `pkg/tessitura/`; all consumer code uses `pkg/domain/` types.
- **Constituent first:** Primary domain object. Expand to others as needed.
- **JSON output only** for now.
- **Auth:** Tessitura uses 4-part basic auth: base64(`username:usergroup:location:password`), stored in `~/.tess/config.json`.
- **`--with` flag** for optional sub-objects (affiliations, notes, etc.). Addresses, emails, phones, and salutations are always included via the `/Detail` endpoint. Use `--with all` to enable everything optional.
- **Batching:** Used in two ways:
  - *Constituents:* When any `--with` extras are requested, `GetConstituentFull` uses the Tessitura `/api/Batch` endpoint to fetch detail + all extras in a single HTTP call. Fixed request IDs: 1=detail, 2=affiliations (individual), 3=affiliations (group), 4=notes. When adding new `--with` options, assign the next available ID and add a case in the switch.
  - *Reports:* `GetReport` always batches request IDs 1=base (`/Reporting/Reports/{id}`) and 2=detail (`/Reporting/Reports/{id}/Details`). This is necessary because `AllowQuery` and `QueryStringAppend` only exist on the base endpoint while `Parameters` only exists on the detail endpoint.
- **Concurrent fetching:** Multiple IDs are fetched with goroutines in `GetConstituentsBatch` and `GetReportsBatch`.
- **Optimistic locking** for updates: Tessitura requires passing `UpdatedDateTime` back on PUT operations.
- **`report list` active-only by default:** The API has no server-side inactive filter, so `report list` filters client-side. Active reports are shown by default; use `--include-inactive` to see all.
- **`report request list` active-only by default:** The API supports `activeOnly` as a server-side query param, so `report request list` passes it directly. Active requests only by default; use `--include-inactive` to see completed/cancelled requests.
- **`report request results` pagination:** The Results endpoint is paginated. Default page size is 100 (matching the API default). If more results exist, the count of remaining results is printed to stderr. Use `--page` to paginate.

## Implemented commands

### `tess configure`
Interactive prompt to set API credentials. Saves to `~/.tess/config.json`.

### `tess constituent get <id> [id...]`
Fetches one or more constituents by ID. Always returns a JSON array.
- `--with affiliations` — attach affiliations (org/household memberships)
- `--with notes` — attach constituent notes
- `--with all` — attach all optional data
- Reads IDs from stdin (one per line) if piped

### `tess constituent search [query]`
Searches constituents. Returns a JSON array of summary records.
- Free-text: positional arg(s)
- Basic structured: `--last-name`, `--first-name`, `--street`, `--postal-code`, `--id`
- Advanced (one at a time): `--email`, `--phone`, `--order-no`, `--web-login`, `--customer-service-no`
- Filters: `--groups` (comma-separated: individuals, organizations, households), `--include-affiliations`

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
# Current platform
make build        # outputs ./tess

# All platforms (outputs to dist/)
make all

# Release archives
make release
```

Or directly:

```bash
go build -o tess ./cmd/tess
```

## Schema maintenance

`schemas/constituent.schema.json` documents the `Constituent` domain object shape. It must be updated manually when fields are added to or removed from `pkg/domain/constituent.go`. It currently does not include the `notes` array.
