## wistia brands apply

Apply Brand

### Synopsis

Applies a brand to a media, folder, or channel, so that resource is styled
by the brand's colors, fonts, logos, and layout.

A brand has no effect until it is applied to something. Media inherit from
their folder, and folders from the account's default brand, so applying a
brand to a folder styles everything inside it that has no brand of its own.

Applying the account-level default brand (`is_default: true`) is how a
resource is un-branded: it detaches the resource so it inherits again.

By default this also clears any brand-mapped appearance settings the
resource had set directly, so the brand is what shows. Pass
`clear_overrides: false` to leave those in place.

Responds with the brand now in effect on the resource, which is not always
the one you applied — detaching a media returns the brand it falls back to.

Webinars can't be branded through this endpoint yet.

## Requires api token with one of the following permissions
```
All data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia brands apply [brand-id] [flags]
```

### Examples

```
  wistia brands apply --brand-id <id> --resource-type media --resource-id abcde12345
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -b, --brand-id string        The id of the brand to apply (or pass it as the [brand-id] argument)
  -c, --clear-overrides        When true (the default), appearance settings the resource had set directly are cleared for the fields the brand controls, so the brand is what shows. Set to false to leave them in place, in which case they continue to win over the brand. (default true)
  -h, --help                   help for apply
      --resource-id string     The id of the resource being branded. [required]
      --resource-type string   The kind of resource being branded. Webinars can't be branded through this endpoint yet. (options: media, folder, channel) [required]
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
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia brands](wistia_brands.md)	 - Operations for brands

### Machine interface

* `wistia brands apply --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia brands apply --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
