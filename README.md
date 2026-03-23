# Durb

A Go CLI client for the Tessitura REST API that maps table-oriented endpoints to domain objects.

The CLI command is `tess`. The project name "Durb" is short for d'Urbervilles -- a play on _Tess of the d'Urbervilles_.

Built by the [Folger Shakespeare Library](https://www.folger.edu).

## Why

The Tessitura API maps directly to database tables, spreading a single concept like a "constituent" across many endpoints. Durb consolidates these into clean domain objects -- a single `tess constituent get` call returns addresses, emails, phones, salutations, and affiliations in one unified JSON response.

## Install

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
tess constituent get 12345 --with all
```

Addresses, emails, phones, and salutations are always included. The `--with` flag adds additional data like affiliations.

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

### Aliases

`constituent` can be shortened to `con`:

```bash
tess con get 12345
tess con search "Smith"
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

- Contact info (addresses, emails, phones, salutations) is filtered to only the constituent's own records -- affiliated records are excluded
- Reference fields like prefix, suffix, and gender are flattened from `{"Id": 5, "Description": "Mr."}` to just `"Mr."`

## Compatibility

Tested against Tessitura API version **16.0.27.97921**. Other versions may work but are not guaranteed.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for release history.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
