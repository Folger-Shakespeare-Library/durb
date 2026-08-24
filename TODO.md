# TODO

- Flatten `report request` to two levels (e.g. `report-request list`) to match CLI best practices (AWS, gh, Stripe all use two-level noun-verb).
- Revisit whether `config init` should suppress password echo during interactive input. Current behavior (visible input) matches AWS/Stripe/Twilio/gh conventions, but worth reconsidering.
- `constituent search --email` results can be misleading. The Tessitura search API matches against all emails on a constituent, but the summary only shows the primary email. Example: searching `--email jane@example.org` may return a constituent whose displayed email is `jane@gmail.com` because `jane@example.org` is a non-primary address that matched but isn't shown. This applies to both `Equals` and `Like` operators. Consider a `--full` flag to fetch complete records for each result, or piping to `constituent get`.
