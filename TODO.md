# TODO

- Flatten `report request` to two levels (e.g. `report-request list`) to match CLI best practices (AWS, gh, Stripe all use two-level noun-verb).
- Revisit whether `config init` should suppress password echo during interactive input. Current behavior (visible input) matches AWS/Stripe/Twilio/gh conventions, but worth reconsidering.
