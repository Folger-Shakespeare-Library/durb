# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [0.1.0] - 2026-03-21

### Added

- `tess configure` — interactive credential setup with secure file permissions (0600)
- `tess constituent get` — fetch one or more constituents by ID with full domain object mapping
  - Addresses, emails, phones, and salutations included by default
  - Affiliated contact info filtered out (only the constituent's own records)
  - `--with affiliations` to attach affiliation data
  - Multiple IDs fetched concurrently; related data batched per constituent
  - Accepts IDs as arguments or via stdin for piping
  - Always returns a JSON array
- `tess constituent search` — search by free text, structured fields, or advanced criteria
  - Basic: `--last-name`, `--first-name`, `--street`, `--postal-code`, `--id`
  - Advanced: `--email`, `--phone`, `--order-no`, `--web-login`, `--customer-service-no`
  - Filters: `--groups`, `--include-affiliations`
  - Duplicate results automatically removed
- Cross-platform builds for macOS, Linux, and Windows
- Version stamped from git tags at build time (`tess --version`)
