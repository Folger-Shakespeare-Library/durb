# Durb

A CLI for the Tessitura REST API. Consolidates table-oriented API endpoints into domain objects. All output is JSON.

Named for d'Urbervilles — *Tess of the d'Urbervilles*, matching the Tessitura/Tess name.

## Install

```bash
go install github.com/Folger-Shakespeare-Library/durb/cmd/tess@latest
```

## Configuration

```bash
tess configure
```

Prompts for hostname, username, user group, location, and password. Credentials are stored in `~/.tess/config.json` with mode `0600`. On Linux and macOS, `tess` refuses to run if the config file is group- or world-readable.

## Commands

### tess constituent get \<id\> [id...]

Fetch constituents by ID. Accepts multiple IDs or reads them from stdin.

Addresses, emails, phones, digital addresses, and salutations are always included. Optional attachments via `--with`:

    affiliations    organization/household memberships
    aliases         alternative names
    associations    relationships to other constituents
    logins          web login credentials
    notes           constituent notes
    all             all of the above

```bash
tess constituent get 12345 --with affiliations,notes
tess constituent search "Smith" | jq -r '.[].id' | tess constituent get --with all
```

### tess constituent search [query]

Search constituents. Returns summary records.

    --last-name, --first-name, --street, --postal-code, --id
    --email, --phone, --order-no, --web-login, --customer-service-no
    --groups individuals,organizations,households
    --include-affiliations

The advanced flags (`--email`, `--phone`, `--order-no`, `--web-login`, `--customer-service-no`) are mutually exclusive with each other and with free-text search.

### tess report get \<id\> [id...]

Fetch reports by ID. Merges the base report and detail endpoint (parameters, indicators) in a single batched call. Accepts multiple IDs or stdin.

### tess report list

List reports. Active only by default.

    --type-ids 6          filter by report type (comma-delimited)
    --category-ids 9,12   filter by category (comma-delimited)
    --include-inactive

### tess report request get \<id\> [id...]

Fetch report requests by ID. Includes parameter values via batched call. Accepts multiple IDs or stdin.

### tess report request list

List report requests. Active only by default.

    --include-inactive    include completed and cancelled

### tess report request results

Paginated view combining request, schedule, and report data.

    --report-id perfseatingbook
    --schedule-name "Daily Seating"
    --start-date 2025-06-01    --end-date 2025-06-30
    --my-reports-only          --recent-only
    --include-public           --include-errors    --include-deleted
    --page 2                   --page-size 50

When more results exist, the remaining count is printed to stderr.

## Aliases

    constituent → con
    configure   → config

## Examples

```bash
tess con get 12345 | jq '.[0].displayName'
tess con get 12345 | jq '.[0].emails[].address'
tess con search --last-name Smith | jq -r '.[].id' | tess con get --with all
tess report list --category-ids 9 | jq -r '.[].id' | tess report get
```

## Compatibility

Tested against Tessitura API v16.0.27.97921.

## License

[GNU General Public License v3.0](LICENSE)
