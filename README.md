# tess

A command-line client for the Tessitura REST API.

## SYNOPSIS

```
tess config init --profile <name>
tess config show | path | list | use <name>

tess crm constituent get [--with <extra>] <id> [id...]
tess crm constituent search [query] [--last-name <name>] [--email <addr>] ...
tess crm constituent create --first <name> --last <name> --email <addr> ...
tess crm constituent update --id <id> --email <addr> --allow-marketing[=false]
tess crm constituent set-status --id <id> --status <desc> [--reason <desc>]

tess crm activity create --constituent-id <id> --activity-type-id <id> --status-id <id> --datetime <dt>
tess crm activity list --constituent-id <id> [--activity-type-id <id>]
tess crm activity delete --activity-id <id>

tess crm attribute set --constituent-id <id> --attribute-type-id <id> --value <val>
tess crm attribute list --constituent-id <id> [--attribute-type-id <id>]
tess crm attribute delete --attribute-id <id>

tess crm interest enable --constituent-id <id> --interest-type-ids <id>[,<id>...]
tess crm interest disable --constituent-id <id> --interest-type-ids <id>[,<id>...]
tess crm interest list --constituent-id <id>

tess report get <id> [id...]
tess report list [--type-ids <ids>] [--category-ids <ids>] [--include-inactive]
tess report request get <id> [id...]
tess report request list [--include-inactive]
tess report request results [--report-id <id>] [--page <n>] ...

tess ref activity-statuses list
tess ref activity-types list
tess ref address-types list
tess ref affiliation-types list
tess ref alias-types list
tess ref appeal-categories list
tess ref association-types list
tess ref business-units list
tess ref campaign-categories list
tess ref constituency-types list
tess ref constituent-groups list
tess ref constituent-inactives list
tess ref constituent-types list
tess ref contact-permission-categories list
tess ref contact-permission-types list
tess ref contribution-designations list
tess ref contribution-import-sets list
tess ref countries list
tess ref delivery-methods list
tess ref designation-codes list
tess ref donation-levels list
tess ref electronic-address-types list
tess ref genders list
tess ref inactive-reasons list
tess ref interest-categories list
tess ref interest-types list
tess ref keyword-categories list
tess ref keywords list
tess ref languages list
tess ref login-types list
tess ref machine-settings list
tess ref membership-benefit-frequencies list
tess ref membership-benefit-types list
tess ref membership-level-categories list
tess ref membership-level-trends list
tess ref membership-periods list
tess ref membership-standings list
tess ref membership-statuses list
tess ref note-types list
tess ref order-categories list
tess ref original-sources list
tess ref payment-types list
tess ref performance-statuses list
tess ref performance-types list
tess ref philanthropy-types list
tess ref plan-priorities list
tess ref plan-sources list
tess ref plan-statuses list
tess ref plan-types list
tess ref planned-giving-codes list
tess ref planned-giving-fundings list
tess ref planned-giving-gift-types list
tess ref planned-giving-on-files list
tess ref planned-giving-purposes list
tess ref planned-giving-sources list
tess ref planned-giving-statuses list
tess ref prefixes list
tess ref price-categories list
tess ref price-type-categories list
tess ref price-type-groups list
tess ref pronouns list
tess ref recognition-types list
tess ref report-categories list
tess ref report-types list
tess ref sales-channels list
tess ref seat-codes list
tess ref seat-statuses list
tess ref seasons list
tess ref sections list
tess ref states list
tess ref suffixes list
tess ref theaters list
tess ref user-groups list
```

## DESCRIPTION

`tess` reads and writes records in a Tessitura ticketing and CRM system. It
talks to the Tessitura REST API and prints JSON to standard output.

The project name is Durb, short for d'Urbervilles (from *Tess of the
d'Urbervilles*). The command name is `tess`.

All output is JSON. Single-record commands return a JSON object. Multi-record
commands return a JSON array. Commands that accept multiple IDs always return
an array, even for a single ID.

## INSTALLATION

```
go install github.com/Folger-Shakespeare-Library/durb/cmd/tess@latest
```

Or build from source:

```
go build -o tess ./cmd/tess
```

## CONFIGURATION

`tess` authenticates using Tessitura's four-part basic auth scheme. The
credential string is `base64(username:usergroup:location:password)`.

Run `tess config init --profile <name>` to store credentials interactively.
You will be prompted for hostname, username, user group, location, and
password. The config file is created at `~/.config/tess/config.json` (or
`$XDG_CONFIG_HOME/tess/config.json`) with mode 0600. `tess` refuses to
start if the file is group- or world-readable.

Multiple profiles can coexist in one config file. The active profile is
resolved in this order:

1. `--profile` flag (available on every command)
2. `TESSITURA_PROFILE` environment variable
3. `default_profile` field in config.json
4. The literal name `"default"`

Individual fields can be overridden with environment variables. This is the
easiest way to run `tess` in CI or scripts without a config file:

    TESSITURA_HOSTNAME
    TESSITURA_USERNAME
    TESSITURA_USER_GROUP
    TESSITURA_LOCATION
    TESSITURA_PASSWORD

### Config subcommands

**tess config init** `--profile <name>`

Set up credentials for a named profile. Prompts for each field.
Overwrites an existing profile after confirmation.

**tess config show** `[--profile <name>]`

Print the resolved profile. The password is partially masked.

**tess config path**

Print the config file path.

**tess config list**

List all profiles. The active profile is marked with `*`.

**tess config use** `<name>`

Set the default profile for subsequent commands.

## CONSTITUENTS

### tess crm constituent get

```
tess crm constituent get [--with <extra>] <id> [id...]
```

Fetch one or more constituents by ID. Always returns a JSON array.

The base record always includes addresses, emails, phones, digital
addresses, and salutations. These come from the `/Detail` endpoint.
Records belonging to affiliated constituents are filtered out.

IDs can be passed as positional arguments, read from standard input (one
per line), or both. Multiple IDs are fetched concurrently. The alias
`con` can be used in place of `constituent`.

**--with** *extra*

Attach additional data to each constituent. May be specified more than
once, or as a comma-separated list. Available extras:

    affiliations    organization and household memberships
    aliases         alternative names
    associations    relationships to other constituents
    logins          web login credentials
    notes           constituent notes
    all             all of the above

When any extras are requested, `tess` uses the Tessitura batch endpoint
(`/api/Batch`) to fetch the detail record and all extras in a single
HTTP call.

Examples:

```
tess crm constituent get 12345
tess crm constituent get 12345 67890 --with affiliations --with notes
tess crm constituent get --with all 12345
echo "12345" | tess crm constituent get
tess crm constituent search "Smith" | jq -r '.[].id' | tess crm constituent get --with all
```

### tess crm constituent search

```
tess crm constituent search [query] [flags]
```

Search for constituents. Returns a JSON array of summary records.

The search operates in one of three modes. The modes are mutually
exclusive; combining them is an error.

**Free-text search.** Pass the query as a positional argument.

```
tess crm constituent search "Jane Smith"
```

**Structured search.** Use named flags. These can be combined.

    --last-name <name>      last name
    --first-name <name>     first name
    --street <addr>         street address
    --postal-code <code>    postal / ZIP code
    --id <id>               constituent ID

```
tess crm constituent search --last-name Smith --first-name Jane
```

**Advanced search.** Use exactly one of these flags. Only one may be
specified per invocation.

    --email <addr>              email address
    --phone <number>            phone number
    --order-no <number>         order number
    --web-login <username>      web login username
    --customer-service-no <no>  customer service number

    --op <operator>     comparison operator: Equals (default), Like,
                        LessThan, GreaterThan

```
tess crm constituent search --email jane@example.com
tess crm constituent search --phone 5551234567 --op Like
```

**Filters.** These work with any search mode.

    --groups <list>             comma-separated: individuals, organizations,
                                households
    --include-affiliations      include affiliated constituents in results

### tess crm constituent create

```
tess crm constituent create --first <name> --last <name> --email <addr>
    --constituent-type-id <id> --original-source-id <id> [flags]
```

Create a new constituent. Returns the created record as JSON.

All five flags above are required. There are no defaults for write
operations. The sort name is auto-generated as `"LastName, FirstName"`.

    --first <name>              first name
    --last <name>               last name
    --email <addr>              email address
    --constituent-type-id <id>  constituent type (e.g. 1 = Individual)
    --original-source-id <id>   original source ID
    --street <addr>             street address line 1
    --postal-code <code>        postal / ZIP code
    --allow-marketing           allow email marketing (default false)

`--first` is limited to 20 characters, `--last` to 55, and
`--postal-code` to 10. Exceeding these limits is a hard error.

### tess crm constituent update

```
tess crm constituent update --id <id> --email <addr> --allow-marketing[=false]
```

Update properties on an existing constituent. Currently supports
updating the email marketing flag.

Requires `--id` and `--email` to identify the electronic address
record. Uses optimistic locking: `tess` fetches the current record,
reads its `UpdatedDateTime`, and passes it back on the PUT.

### tess crm constituent set-status

```
tess crm constituent set-status --id <id> --status <description> [--reason <description>]
```

Change a constituent's inactive status. `--status` and `--reason` take
human-readable descriptions (e.g. "Active", "Inactive", "Deceased"),
not numeric codes. `tess` resolves them against Tessitura's reference
data by case-insensitive match.

`--reason` is required when setting a constituent inactive. It is an
error to pass `--reason` when setting status to Active.

```
tess crm constituent set-status --id 12345 --status Inactive --reason Duplicate
tess crm constituent set-status --id 12345 --status Active
```

## ACTIVITIES

```
tess crm activity create --constituent-id <id> --activity-type-id <id>
    --status-id <id> --datetime <datetime> [--notes <text>] [--unique]
tess crm activity list --constituent-id <id> [--activity-type-id <id>]
tess crm activity delete --activity-id <id>
```

Activities are special activity records attached to a constituent.

**tess crm activity create** creates a record. `--datetime` accepts a
full ISO 8601 datetime or a bare date (`YYYY-MM-DD`), which defaults to
midnight in the system's local timezone.

The `--unique` flag makes creation idempotent: if an activity with the
same type and datetime already exists on the constituent, `tess` returns
the existing record instead of creating a duplicate.

**tess crm activity list** lists activities for a constituent.
`--activity-type-id` filters to a specific type.

**tess crm activity delete** deletes a single activity record by its ID.

## ATTRIBUTES

```
tess crm attribute set --constituent-id <id> --attribute-type-id <id> --value <value>
tess crm attribute list --constituent-id <id> [--attribute-type-id <id>]
tess crm attribute delete --attribute-id <id>
```

Attributes are keyword/value pairs attached to a constituent.

**tess crm attribute set** is an upsert. If the constituent already has
an attribute of the given type, `tess` updates its value. Otherwise it
creates a new record. Uses optimistic locking for updates.

**tess crm attribute list** lists attributes on a constituent.
`--attribute-type-id` filters to a specific type.

**tess crm attribute delete** deletes a single attribute record by its ID.

## INTERESTS

```
tess crm interest enable --constituent-id <id> --interest-type-ids <id>[,<id>...]
tess crm interest disable --constituent-id <id> --interest-type-ids <id>[,<id>...]
tess crm interest list --constituent-id <id>
```

Interests are boolean flags attached to a constituent.

**tess crm interest enable** and **disable** accept a comma-separated
list of interest type IDs. They will not create duplicate records: if
the interest assignment already exists, `tess` updates it. Uses
optimistic locking for updates.

**tess crm interest list** lists all interest assignments for a
constituent.

## REPORTS

### tess report get

```
tess report get <id> [id...]
```

Fetch one or more report definitions by ID. Always returns a JSON
array. Always includes full detail (parameters, indicators) via a
batched API call that combines the base endpoint and the detail
endpoint. Multiple IDs are fetched concurrently.

Reads IDs from standard input (one per line) if piped.

### tess report list

```
tess report list [--type-ids <ids>] [--category-ids <ids>] [--include-inactive]
```

List all reports. Returns a JSON array of summary records (no parameters
or detail-only fields; use `report get` for the full record).

Active reports only by default. The API has no server-side inactive
filter, so `tess` filters client-side.

    --type-ids <ids>        filter by report type ID (comma-delimited)
    --category-ids <ids>    filter by category ID (comma-delimited)
    --include-inactive      include inactive reports

### tess report request get

```
tess report request get <id> [id...]
```

Fetch one or more report requests by ID. Always returns a JSON array.
Always includes parameter values from the detail endpoint via a batched
API call. Multiple IDs are fetched concurrently.

Reads IDs from standard input (one per line) if piped.

### tess report request list

```
tess report request list [--include-inactive]
```

List all report requests. Active requests only by default. Unlike
`report list`, this uses a server-side `activeOnly` filter.

    --include-inactive      include completed and cancelled requests

### tess report request results

```
tess report request results [flags]
```

List scheduled report results. Returns a combined view that merges
report request, schedule, and report definition data. Paginated.

    --report-id <id>            filter by report ID
    --schedule-name <name>      filter by schedule name
    --start-date <YYYY-MM-DD>   start of date range
    --end-date <YYYY-MM-DD>     end of date range
    --my-reports-only           only results for the current user
    --recent-only               only recent results
    --include-public            include public report results
    --include-errors            include errored results
    --include-deleted           include results whose output was deleted
    --page <n>                  page number (default 1)
    --page-size <n>             results per page (default 100)

If more results exist beyond the current page, the count of remaining
results is printed to standard error.

## REFERENCE DATA

```
tess ref activity-statuses list
tess ref activity-types list
tess ref address-types list
tess ref affiliation-types list
tess ref alias-types list
tess ref appeal-categories list
tess ref association-types list
tess ref business-units list
tess ref campaign-categories list
tess ref constituency-types list
tess ref constituent-groups list
tess ref constituent-inactives list
tess ref constituent-types list
tess ref contact-permission-categories list
tess ref contact-permission-types list
tess ref contribution-designations list
tess ref contribution-import-sets list
tess ref countries list
tess ref delivery-methods list
tess ref designation-codes list
tess ref donation-levels list
tess ref electronic-address-types list
tess ref genders list
tess ref inactive-reasons list
tess ref interest-categories list
tess ref interest-types list
tess ref keyword-categories list
tess ref keywords list
tess ref languages list
tess ref login-types list
tess ref machine-settings list
tess ref membership-benefit-frequencies list
tess ref membership-benefit-types list
tess ref membership-level-categories list
tess ref membership-level-trends list
tess ref membership-periods list
tess ref membership-standings list
tess ref membership-statuses list
tess ref note-types list
tess ref order-categories list
tess ref original-sources list
tess ref payment-types list
tess ref performance-statuses list
tess ref performance-types list
tess ref philanthropy-types list
tess ref plan-priorities list
tess ref plan-sources list
tess ref plan-statuses list
tess ref plan-types list
tess ref planned-giving-codes list
tess ref planned-giving-fundings list
tess ref planned-giving-gift-types list
tess ref planned-giving-on-files list
tess ref planned-giving-purposes list
tess ref planned-giving-sources list
tess ref planned-giving-statuses list
tess ref prefixes list
tess ref price-categories list
tess ref price-type-categories list
tess ref price-type-groups list
tess ref pronouns list
tess ref recognition-types list
tess ref report-categories list
tess ref report-types list
tess ref sales-channels list
tess ref seat-codes list
tess ref seat-statuses list
tess ref seasons list
tess ref sections list
tess ref states list
tess ref suffixes list
tess ref theaters list
tess ref user-groups list
```

List reference (lookup table) data from Tessitura.

**tess ref activity-statuses list** returns available special activity
statuses. Used with `activity create --status-id`.

**tess ref activity-types list** returns available special activity
types. Used with `activity create --activity-type-id`.

**tess ref address-types list** returns available address types.

**tess ref affiliation-types list** returns available affiliation types
for classifying constituent affiliations (org/household memberships).

**tess ref alias-types list** returns available alias types for
classifying constituent aliases.

**tess ref appeal-categories list** returns available categories for
fundraising appeals.

**tess ref association-types list** returns available association types
for classifying constituent associations.

**tess ref business-units list** returns available business units.

**tess ref campaign-categories list** returns available categories for
fundraising campaigns.

**tess ref constituency-types list** returns available constituency
types for constituent classification.

**tess ref constituent-groups list** returns available constituent
groups (individuals, organizations, households).

**tess ref constituent-inactives list** returns available inactive status
types for constituents (`id`, `description`, `inactive`).

**tess ref constituent-types list** returns available constituent types.
Used with `constituent create --constituent-type-id`.

**tess ref contact-permission-categories list** returns available contact
permission categories for marketing consent groupings.

**tess ref contact-permission-types list** returns available contact
permission types for marketing consent configuration.

**tess ref contribution-designations list** returns available
contribution designation codes for directing contributions.

**tess ref contribution-import-sets list** returns import set
definitions for contribution imports.

**tess ref countries list** returns available countries.

**tess ref delivery-methods list** returns available ticket delivery
methods.

**tess ref designation-codes list** returns available codes for directing
contributions.

**tess ref donation-levels list** returns available giving levels/tiers
with recognition type references.

**tess ref electronic-address-types list** returns available electronic
address types (email, phone, web type IDs).

**tess ref genders list** returns available genders with default prefix
and pronoun references.

**tess ref inactive-reasons list** returns available inactive reasons for
constituents. Used with `constituent set-status --reason`.

**tess ref interest-categories list** returns available interest
categories (metadata for interest types).

**tess ref interest-types list** returns available interest types. Used
with `interest enable/disable --interest-type-ids`.

**tess ref keyword-categories list** returns available keyword categories
(metadata for attribute types).

**tess ref keywords list** returns available keywords (attribute types in
Tessitura). Used with `attribute set --attribute-type-id`.

**tess ref languages list** returns available language codes.

**tess ref login-types list** returns available login types for
classifying constituent web logins.

**tess ref machine-settings list** returns machine settings (workstation
name, card reader configuration, merchant IDs, audit fields).

**tess ref membership-benefit-frequencies list** returns available
frequency options for membership benefits.

**tess ref membership-benefit-types list** returns available types of
membership benefits.

**tess ref membership-level-categories list** returns available categories
for membership levels.

**tess ref membership-level-trends list** returns available trend
classifications for membership levels.

**tess ref membership-periods list** returns available membership period
definitions.

**tess ref membership-standings list** returns available standing types for
memberships.

**tess ref membership-statuses list** returns available status values for
memberships.

**tess ref note-types list** returns available note types for classifying
constituent notes.

**tess ref order-categories list** returns available order categories for
order classification.

**tess ref original-sources list** returns available original sources.
Used with `constituent create --original-source-id`.

**tess ref payment-types list** returns available payment types for
payment method lookups.

**tess ref performance-statuses list** returns available performance
statuses for event classification.

**tess ref performance-types list** returns available performance types
for event classification.

**tess ref philanthropy-types list** returns available types of
philanthropic activity.

**tess ref plan-priorities list** returns available priority levels for
development plans.

**tess ref plan-sources list** returns available sources for development
plans.

**tess ref plan-statuses list** returns available status values for
development plans.

**tess ref plan-types list** returns available types of
development/solicitation plans.

**tess ref planned-giving-codes list** returns available codes for
planned giving.

**tess ref planned-giving-fundings list** returns available funding types
for planned gifts.

**tess ref planned-giving-gift-types list** returns available types of
planned gifts.

**tess ref planned-giving-on-files list** returns available on-file
statuses for planned gifts.

**tess ref planned-giving-purposes list** returns available purposes for
planned gifts.

**tess ref planned-giving-sources list** returns available sources for
planned gifts.

**tess ref planned-giving-statuses list** returns available statuses for
planned gifts.

**tess ref prefixes list** returns available name prefixes (Mr., Mrs.,
Dr., etc.).

**tess ref price-categories list** returns available price categories.

**tess ref price-type-categories list** returns available price type
categories.

**tess ref price-type-groups list** returns available price type groups.

**tess ref pronouns list** returns available pronoun sets.

**tess ref recognition-types list** returns available donor recognition
type definitions.

**tess ref report-categories list** returns available report categories.
Used with `report list --category-ids`.

**tess ref report-types list** returns available report types. Used with
`report list --type-ids`.

**tess ref sales-channels list** returns available sales channels (box
office, online, phone).

**tess ref seat-codes list** returns available seat codes for venue
configuration.

**tess ref seat-statuses list** returns seat statuses (`id`,
`description`, `statusCode`, `statusLegend`, `statusPriority`,
`inactive`).

**tess ref seasons list** returns available seasons for reporting and
ticketing.

**tess ref sections list** returns available venue sections.

**tess ref states list** returns available states and provinces with
country references.

**tess ref suffixes list** returns available name suffixes (Jr., Sr.,
III, etc.).

**tess ref theaters list** returns available theaters (venues) with
address and capacity data.

**tess ref user-groups list** returns available user groups with
division and permission data.

## ENVIRONMENT

    TESSITURA_PROFILE       profile name (overrides default_profile in config)
    TESSITURA_HOSTNAME      API hostname (overrides profile value)
    TESSITURA_USERNAME      API username (overrides profile value)
    TESSITURA_USER_GROUP    API user group (overrides profile value)
    TESSITURA_LOCATION      API location (overrides profile value)
    TESSITURA_PASSWORD      API password (overrides profile value)

## FILES

    ~/.config/tess/config.json      credentials and profile configuration

## EXAMPLES

Pipe a search into a full fetch:

```
tess crm constituent search --last-name Smith \
  | jq -r '.[].id' \
  | tess crm constituent get --with all
```

Create a constituent and capture the new ID:

```
NEW_ID=$(tess crm constituent create --first Jane --last Smith \
    --email jane@example.com --constituent-type-id 1 \
    --original-source-id 10 | jq '.id')
```

Idempotent activity creation (safe to run repeatedly):

```
tess crm activity create --constituent-id 12345 \
    --activity-type-id 3 --status-id 1 \
    --datetime 2025-06-15 --unique
```

Set an attribute and enable interests in one session:

```
tess crm attribute set --constituent-id 12345 \
    --attribute-type-id 7 --value "Yes"
tess crm interest enable --constituent-id 12345 \
    --interest-type-ids 1,5,12
```

List active reports in a category:

```
tess report list --category-ids 9
```

Use environment variables for CI:

```
export TESSITURA_HOSTNAME=https://example.tnhs.cloud/tessitura
export TESSITURA_USERNAME=api_user
export TESSITURA_USER_GROUP=API
export TESSITURA_LOCATION=1
export TESSITURA_PASSWORD=secret
tess crm constituent get 12345
```
