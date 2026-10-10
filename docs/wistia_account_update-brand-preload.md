## wistia account update-brand-preload

Update Brand Preload

### Synopsis

Persists the account's default page logo (by Bakery hashed_id) and
default player color. Both fields are optional independently — omit a
field to leave that account setting untouched. Passing an empty string
for `selected_logo_hashed_id` clears the logo.

Requires the OAuth contact to be an owner or manager of the account
(or a Wistia admin) — mirrors the auth check on the underlying
`updateWtwBrandKitAccountSettings` GraphQL mutation.

Deliberately narrower than the mutation: this endpoint does not
create/update BrandKits or set body font family. Glass's onboarding
customize step writes only these two fields; broader brand-kit
editing continues to happen through the WTW web UI + GraphQL.

## Requires api token with one of the following permissions
```
(any scope allowed)
```

```
wistia account update-brand-preload [flags]
```

### Examples

```
  wistia account update-brand-preload
```

### Options

```
      --body string                      Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -h, --help                             help for update-brand-preload
      --schema                           Print the exact JSON Schema of the request body and exit
      --selected-logo-hashed-id string   Bakery hashed_id of an uploaded logo image, which will become the
                                         account's default page logo. Omit to leave the current logo
                                         untouched. Pass an empty string to clear the logo.
      --selected-player-color string     Hex color string (e.g. "#3366FF") for the account's default player
                                         color — 6 hex digits, with or without the leading '#'. Omit or send
                                         an empty string to leave the current color untouched (there is no
                                         clear operation — color always has a value). Malformed values are
                                         rejected at the API boundary; without this check, the model's
                                         sanitize step would return nil and silently reset the account color
                                         to the global default.
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

* [wistia account](wistia_account.md)	 - Operations for account

### Machine interface

* `wistia account update-brand-preload --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia account update-brand-preload --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `wistia account update-brand-preload --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
