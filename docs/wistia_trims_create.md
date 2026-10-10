## wistia trims create

Create Media from Trims

### Synopsis

Creates a new media that trims off parts of an existing media.

By default, the `trims` parameter specifies time ranges to **remove** from the media. When `keep_trims` is set to `true`, the `trims` parameter instead specifies time ranges to **keep** in the media.

**NOTE:** currently this endpoint only supports trimming video files.

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
authorization granting the `edit` permission on this media can also be used.

```
wistia trims create [media-hashed-id] [flags]
```

### Examples

```
  wistia trims create --media-hashed-id <id> --trims <value>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -h, --help                     help for create
  -k, --keep-trims               When set to true, the trims parameter is treated as ranges to keep rather than ranges to remove. Defaults to false.
  -m, --media-hashed-id string   The hashed ID of the media. (or pass it as the [media-hashed-id] argument)
  -t, --trims stringArray        An array of strings matching the format of HH:MM:SS.mmm-HH:MM:SS.mmm where HH is hours, MM is minutes, SS is seconds and mmm is milliseconds. When keep_trims is false (default), the ranges specify parts of the media to remove. When keep_trims is true, the ranges specify parts of the media to keep. [required]
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

* [wistia trims](wistia_trims.md)	 - Operations for trims

### Machine interface

* `wistia trims create --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia trims create --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
