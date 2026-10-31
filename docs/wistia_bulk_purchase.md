## wistia bulk purchase

Create Bulk Purchase

### Synopsis

Submits either an `actions` array of up to 1000 orders or one `job` that can
resolve to up to 5000 media. Orders are placed asynchronously. Returns a
background job status whose Show endpoint reports aggregate progress and
per-order results.

Orders in the batch can incur charges, so a saved credit card is required.
Supported resource types are `captions` (Wistia-generated English captions),
`localization` (a dubbed, language-specific version of a media),
`extended_audio_description`, and `text_translation` (the media's transcript
translated into another language, audio untouched). Each order's `id` is the
hashed ID of the media to order for.

What an order costs depends on the account, not on this endpoint. Automated
captions are included at no cost on plans that provide them and billed at
the account's configured per-minute rate otherwise; human-reviewed captions
bill per minute at the account's standard or rush rate; localizations bill
per minute once the account's free-dub allowance is used up; text
translations bill as an overage once the account's included translation
minutes are used up. Check the account's plan and billing settings for its
actual rates.

Orders are priced and placed individually: failures -- an ineligible media,
a language that already has a localization, an account not entitled to buy
-- are reported per order and do not stop the rest of the batch. Pricing and
eligibility match the equivalent single-media endpoints exactly.

Use the Create Bulk Actions endpoint for create, update, and delete work; it
does not accept `purchase`, and this endpoint accepts nothing else.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

[Expiring access tokens](https://docs.wistia.com/reference/post_expiring-token)
with authorizations cannot use this endpoint: each action is authorized
later, in a background job, as the persisted contact that submitted it,
and a token's authorizations are not carried into that job. Such requests
are forbidden.

```
wistia bulk purchase [flags]
```

### Examples

```
  wistia bulk purchase
```

### Options

```
  -a, --actions string   The orders to place, one per media. Maximum 1000 per request, and the
                         request body must stay under 2 MB -- whichever limit is reached first. An
                         oversized body is rejected with a '413' and no order in it is placed.
                         
                         Every order is priced and placed independently: one failing (an
                         ineligible media, an account without a saved card, a language that
                         already has a localization) does not stop the rest of the batch.
                         
                         Use 'job' instead to order for a whole folder, channel, or account.
                         (JSON array)
      --body string      Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -h, --help             help for purchase
  -j, --job string       One order placed for many media, named by a parent ('scope') or listed
                         explicitly ('ids'), so ordering captions for a folder of 47 videos takes one
                         job rather than 47 orders.
                         
                         A 'scope' resolves to exactly what List Media returns for that parent,
                         including media in the folder's subfolders and **archived** media, and to at
                         most 5000 media -- beyond that the job is rejected rather than truncated.
                         
                         The job attempts one order for every media it resolves to. Ineligible media
                         fail individually without placing an order; successful orders are metered
                         and may incur charges according to the account's plan. Confirm the scope and
                         potential cost with the customer before submitting. Not available to external
                         contacts.
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-keyring             Never read or write the OS keychain; store secrets in the config file instead (env: WISTIA_CLI_NO_KEYRING)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia bulk](wistia_bulk.md)	 - Operations for bulk

### Machine interface

* `wistia bulk purchase --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia bulk purchase --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
