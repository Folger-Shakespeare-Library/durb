# Durb

A Go CLI client for the Tessitura REST API that maps table-oriented endpoints to domain objects.

The CLI command is `tess`. The project name "Durb" is short for d'Urbervilles -- a play on _Tess of the d'Urbervilles_.

Built by the [Folger Shakespeare Library](https://www.folger.edu).

## Why

The Tessitura API maps directly to database tables, spreading a single concept like a "constituent" or "report" across many endpoints. Durb consolidates these into clean domain objects -- a single `tess constituent get` call returns addresses, emails, phones, salutations, and affiliations in one unified JSON response; a `tess report get` merges the base report and its detail (including parameters) into one record.

## Install

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

## Setup

```bash
tess configure
```

This prompts for your Tessitura API credentials and saves them to `~/.tess/config.json`. You'll need:

- **Hostname** -- your Tessitura REST API base URL
- **Username**, **User Group**, **Location**, **Password** -- your Tessitura service credentials

On macOS and Linux, the config file is created with `0600` permissions (owner-only read/write), and `tess` will refuse to run if the file is readable by other users. On Windows, the file is protected by the user's own profile directory permissions.

## Usage

### Get a constituent

```bash
tess constituent get 12345
```

Always returns a JSON array (even for a single ID).

### Get multiple constituents

```bash
tess constituent get 12345 67890

# Or pipe IDs from another command
tess constituent search "Smith" | jq -r '.[].id' | tess constituent get
```

Multiple IDs are fetched concurrently.

### Include related data

```bash
tess constituent get 12345 --with affiliations
tess constituent get 12345 --with notes
tess constituent get 12345 --with affiliations,notes
tess constituent get 12345 --with all
```

Addresses, emails, phones, digital addresses, and salutations are always included. The `--with` flag adds optional data:

- `affiliations` -- organization/household memberships
- `aliases` -- alternative names
- `associations` -- relationships to other constituents
- `logins` -- web login credentials
- `notes` -- constituent notes
- `all` -- include everything above

### Search for constituents

```bash
# Free-text search
tess constituent search "Smith"

# Structured search
tess constituent search --last-name Smith
tess constituent search --first-name Jane --last-name Smith
tess constituent search --street "Main Street"
tess constituent search --postal-code 10001

# Advanced search
tess constituent search --email user@example.com
tess constituent search --phone 5551234567
tess constituent search --order-no 99999
tess constituent search --web-login jsmith
tess constituent search --customer-service-no 12345

# Filter by constituent group
tess constituent search "Smith" --groups individuals

# Include affiliated constituents
tess constituent search --last-name Smith --include-affiliations
```

Search returns a JSON array of summary records.

### Piping search into get

```bash
# Get full details for all search results
tess constituent search --last-name Smith | jq -r '.[].id' | tess constituent get

# With affiliations
tess constituent search --last-name Smith | jq -r '.[].id' | tess constituent get --with affiliations
```

### Get a report

```bash
tess report get perfseatingbook
```

Always returns a JSON array. The response includes the full report record: base fields, indicators (`publicIndicator`, `warningIndicator`, `utilityIndicator`), and the `parameters` array -- fetched via a single batched API call.

### Get multiple reports

```bash
tess report get perfseatingbook annualgifts

# Or pipe IDs from a list
tess report list | jq -r '.[].id' | tess report get
```

### List reports

```bash
# All active reports
tess report list

# Filter by type or category
tess report list --type-ids 6
tess report list --category-ids 9
tess report list --type-ids 6 --category-ids 9,12

# Include inactive reports
tess report list --include-inactive
```

Returns a JSON array of summary records. Active reports only by default.

### Get a report request

```bash
tess report request get 12345
```

Always returns a JSON array. Includes parameter values via a batched API call.

### Get multiple report requests

```bash
tess report request get 12345 67890

# Or pipe IDs
tess report request results --report-id perfseatingbook | jq -r '.[].id' | tess report request get
```

### List report requests

```bash
# Active requests only (default)
tess report request list

# Include completed and cancelled
tess report request list --include-inactive
```

### Report request results

A richer, paginated view combining request, schedule, and report data:

```bash
# All results (page 1, 100 per page)
tess report request results

# Filter by report or schedule
tess report request results --report-id perfseatingbook
tess report request results --schedule-name "Daily Seating"

# Date range
tess report request results --start-date 2025-06-01 --end-date 2025-06-30

# Only my results
tess report request results --my-reports-only

# Paginate
tess report request results --page 2 --page-size 50
```

If there are more results beyond the current page, the remaining count is printed to stderr.

### Aliases

`constituent` can be shortened to `con`, and `configure` to `config`:

```bash
tess con get 12345
tess con search "Smith"
tess config
```

## Output

All commands output JSON to stdout. Use `jq` for filtering:

```bash
# Get just the display name
tess con get 12345 | jq '.[0].displayName'

# Get all email addresses
tess con get 12345 | jq '.[0].emails[].address'

# Get primary address
tess con get 12345 | jq '.[0].addresses[] | select(.primary)'
```

## Domain model

Durb maps Tessitura's table-oriented API responses to clean domain objects:

**Constituents**
- Contact info (addresses, emails, phones, digital addresses, salutations) is filtered to only the constituent's own records -- affiliated records are excluded
- Reference fields like prefix, suffix, and gender are flattened from `{"Id": 5, "Description": "Mr."}` to just `"Mr."`
- Aliases, associations, and web logins are available as optional `--with` attachments

**Reports**
- `report get` merges the base report and its detail endpoint into one object via a batched API call
- `category` and `reportType` are sub-objects (`{"id": 9, "description": "..."}`) rather than flat fields
- `report list` filters inactive reports client-side (the Tessitura API has no server-side inactive filter)

**Report requests**
- `report request get` batches the base request and its detail (parameter values) in a single API call
- `report request list` uses the API's server-side `activeOnly` filter; active requests only by default
- `report request results` is a paginated combined entity (request + schedule + report); remaining result count is printed to stderr when more pages exist

## Compatibility

Tested against Tessitura API version **16.0.27.97921**. Other versions may work but are not guaranteed.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
