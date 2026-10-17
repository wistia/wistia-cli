## wistia custom-metadata-field-definitions put-custom-metadata-field-definitions-key

Update Custom Metadata Field Definition

### Synopsis

Updates the editable attributes of a custom metadata field definition addressed by its immutable key: `label`, `default_value`, `position`, and single_select or multi_select options via `config.options` (add, edit, or remove). The `key` and `field_type` are immutable — attempting to change them returns a 422.

Requires the custom metadata feature to be available on your account.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

An [expiring access token](https://docs.wistia.com/reference/post_expiring-token)
created with the `all:delegate_to_contact_permissions` scope and an
`account` authorization granting the `manage-custom-metadata` permission
can also be used.

```
wistia custom-metadata-field-definitions put-custom-metadata-field-definitions-key [key] [flags]
```

### Examples

```
  wistia custom-metadata-field-definitions put-custom-metadata-field-definitions-key --key client
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -c, --config-param string    Type-specific configuration changes. Only valid for field types that have any; currently option add, edit, or remove operations for a single_select or multi_select field.
      --default-value string   A default value for the field, matching the field_type's format.
  -h, --help                   help for put-custom-metadata-field-definitions-key
  -k, --key string             The field's immutable key (or pass it as the [key] argument)
  -l, --label string           The field's display name. Must be unique per account among active fields (case-insensitive).
  -p, --position string        The field's display order within the account, ascending from 0.
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

* [wistia custom-metadata-field-definitions](wistia_custom-metadata-field-definitions.md)	 - Operations for custom-metadata-field-definitions

### Machine interface

* `wistia custom-metadata-field-definitions put-custom-metadata-field-definitions-key --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia custom-metadata-field-definitions put-custom-metadata-field-definitions-key --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
