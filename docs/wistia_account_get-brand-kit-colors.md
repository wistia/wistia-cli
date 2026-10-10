## wistia account get-brand-kit-colors

Get Brand Kit Colors

### Synopsis

Retrieves the current account's brand colors for the Wistia desktop
app's background picker.

`colors` lists solid colors: every brand kit's color tokens (the colors
the web editor offers as "Brand colors"), then each brand's primary and
page background color when it is solid, default brand first. Values are
six-digit hex strings, and a repeated color is listed once. Tokens whose
value isn't a hex color are left out. An account without a brand kit
gets its player color in place of the kit, which is what its default
brand kit would hold.

`brand_gradients` lists each brand's primary and page background color
that is set to a gradient, as color stops sorted by position, default
brand first. Stops whose color isn't a hex color are left out, and a
gradient with fewer than two hex stops left isn't listed.

## Requires api token with one of the following permissions
```
(any scope allowed)
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia account get-brand-kit-colors [flags]
```

### Examples

```
  wistia account get-brand-kit-colors
```

### Options

```
  -h, --help   help for get-brand-kit-colors
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

* `wistia account get-brand-kit-colors --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia account get-brand-kit-colors --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
