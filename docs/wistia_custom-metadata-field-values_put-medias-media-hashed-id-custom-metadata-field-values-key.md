## wistia custom-metadata-field-values put-medias-media-hashed-id-custom-metadata-field-values-key

Set Custom Metadata Field Value

### Synopsis

Sets (or replaces) the value of a custom metadata field on a media, addressed by the field definition's immutable key. The lookup is case-insensitive.

The request body carries a single polymorphic `value` field whose JSON type must match the definition's `field_type`:

- text-like types (`text`, `short_text`, `url`, `email`, `money`, `time`, `datetime`) — a string; format-validated per type (e.g. money is `"USD 12.34"`, time is 24-hour `"14:30"`, datetime is UTC ISO 8601 `"2026-07-10T14:30:00Z"`)
- `number` — a JSON number (a numeric string is coerced)
- `date` — an ISO 8601 date string like `"2026-07-10"`
- `boolean` — a JSON boolean; `false` persists as false (it does not clear the field)
- `single_select` — the chosen option's key (a string); unknown option keys return a 422
- `multi_select` — an array of the chosen options' keys (strings); unknown option keys or a non-array value return a 422
- `contact_ref` — a contact reference object `{"type": "contact" | "contact_group", "id": "<hashed_id>"}`; unknown ids, ids from another account, and group references on fields that do not allow groups return a 422
- `contact_multi_ref` — an array of contact reference objects; the same 422 rules apply per reference, and a non-array value returns a 422

A null or absent `value` clears the field (equivalent to the DELETE endpoint), as does an empty array for `multi_select` and `contact_multi_ref`. Type mismatches and format violations return a 422 with a field-level message.

Only values for active field definitions can be written. Requires the custom metadata feature to be available on your account.


## Requires api token with one of the following permissions
```
Read, update & delete anything
Upload, read & update all media
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia custom-metadata-field-values put-medias-media-hashed-id-custom-metadata-field-values-key [flags]
```

### Examples

```
  wistia custom-metadata-field-values put-medias-media-hashed-id-custom-metadata-field-values-key --media-hashed-id <id> --key client
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -h, --help                     help for put-medias-media-hashed-id-custom-metadata-field-values-key
  -k, --key string               The field definition's immutable key. [required]
  -m, --media-hashed-id string   The hashed ID of the media whose custom metadata field value is to be set. [required]
  -v, --value string             JSON value (one of: string | number | boolean | array of any | { "type": string, "id": string })
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

* [wistia custom-metadata-field-values](wistia_custom-metadata-field-values.md)	 - Operations for custom-metadata-field-values

### Machine interface

* `wistia custom-metadata-field-values put-medias-media-hashed-id-custom-metadata-field-values-key --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia custom-metadata-field-values put-medias-media-hashed-id-custom-metadata-field-values-key --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
