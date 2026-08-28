# Durb

`tess` - read and write Tessitura records

The command is `tess`, short for Tessitura. The project name
Durb is short for d'Urbervilles, from *Tess of the d'Urbervilles*.

    go install github.com/Folger-Shakespeare-Library/durb/cmd/tess@latest

## Configuration

Run `tess config init --profile default` to store credentials.
You will be prompted for the Tessitura hostname, username, user
group, location, and password.

Credentials are saved to `~/.config/tess/config.json` (or
`$XDG_CONFIG_HOME/tess/config.json`). The file is created mode
0600 and `tess` will refuse to start if it is group- or
world-readable.

Multiple profiles can coexist in the same config file. The
active profile is determined by, in order:

1. The `--profile` flag
2. `TESSITURA_PROFILE` environment variable
3. The `default_profile` field in config.json
4. The literal name `"default"`

Individual fields can be overridden with `TESSITURA_HOSTNAME`,
`TESSITURA_USERNAME`, `TESSITURA_USER_GROUP`, `TESSITURA_LOCATION`,
and `TESSITURA_PASSWORD`. This is the easiest way to use `tess`
in CI or scripts without a config file on disk.

Other config commands:

    tess config show       print the active profile (password masked)
    tess config path       print the config file path
    tess config list       list all profiles (* marks active)
    tess config use NAME   set the default profile

## Constituents

`tess crm constituent get` returns a constituent record. The
response always includes addresses, emails, phones, digital
addresses, and salutations. Records from affiliated constituents
are filtered out.

Additional data requires `--with`:

    --with affiliations     organization and household memberships
    --with aliases          alternative names
    --with associations     relationships to other constituents
    --with logins           web login credentials
    --with notes            constituent notes
    --with all              all of the above

All `get` commands (constituent, report, and report request)
accept multiple IDs as arguments and read IDs from standard
input, one per line:

    tess crm constituent search --last-name Smith \
      | jq -r '.[].id' \
      | tess crm constituent get --with all

### Search

`tess crm constituent search` operates in one of three modes,
and the modes are mutually exclusive. Mixing them is an error.

**Free text** uses a positional argument:

    tess crm constituent search "Jane Smith"

**Structured search** uses named flags that can be combined:
`--last-name`, `--first-name`, `--street`, `--postal-code`, `--id`.

    tess crm constituent search --last-name Smith --first-name Jane

**Advanced search** uses a single lookup field: `--email`,
`--phone`, `--order-no`, `--web-login`, or `--customer-service-no`.
Only one may be specified. `--op` sets the comparison (Equals,
Like, LessThan, GreaterThan; default Equals).

    tess crm constituent search --email jane@example.com

`--groups` filters by type (individuals, organizations,
households). Results are deduplicated by constituent ID.

### Create

    tess crm constituent create --first Jane --last Smith \
        --email jane@example.com --constituent-type-id 1 \
        --original-source-id 10

`--first` is silently truncated to 20 characters, `--last` to
55, and `--postal-code` to 10.

### Set-status

    tess crm constituent set-status --id 12345 --status Inactive --reason Deceased

`--status` and `--reason` take human-readable descriptions, not
numeric codes. `tess` resolves them against Tessitura's reference
data by case-insensitive match. A `--reason` is required when
setting a constituent inactive.

## Activities, Attributes, and Interests

These are records attached to a constituent by `--constituent-id`.

`tess crm attribute set` is an upsert: if the constituent already
has an attribute of the given type, it updates the value; otherwise
it creates one. `tess crm interest enable` and `disable` work the
same way and will not create duplicate records.

    tess crm attribute set --constituent-id 12345 \
        --attribute-type-id 7 --value "Yes"
    tess crm interest enable --constituent-id 12345 \
        --interest-type-ids 1,5,12

`tess crm activity create` creates an activity record. The
`--unique` flag makes it idempotent: if an activity with the same
type and datetime already exists, it returns the existing record.

    tess crm activity create --constituent-id 12345 \
        --activity-type-id 3 --status-id 1 \
        --datetime 2025-06-15 --unique

## Reports

    tess report get perfseatingbook
    tess report list --category-ids 9

`tess report get` returns the report definition with its
parameters. `tess report list` shows active reports only by
default; the API has no server-side inactive filter, so `tess`
filters client-side. Pass `--include-inactive` to see all.

    tess report request get 12345
    tess report request list
    tess report request results --report-id perfseatingbook \
        --start-date 2025-06-01 --end-date 2025-06-30

`tess report request results` is paginated (`--page`,
`--page-size`, default 100).

## Reference Data

    tess ref constituent-inactives list
    tess ref seat-statuses list

## All Commands

    tess config init | show | path | list | use
    tess crm constituent get | search | create | update | set-status
    tess crm activity create | list | delete
    tess crm attribute set | list | delete
    tess crm interest enable | disable | list
    tess report get | list
    tess report request get | list | results
    tess ref constituent-inactives list
    tess ref seat-statuses list

Run `tess <command> --help` for flags and usage.
