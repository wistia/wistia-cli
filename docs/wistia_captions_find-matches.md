## wistia captions find-matches

Find Caption Matches

### Synopsis

Finds exact text in caption tracks without modifying them. Matching uses the
same normalization, composite-media boundaries, and time coordinates as the
targeted caption edit endpoint. Fuzzy alternatives are returned separately
as suggestions and are never reported as exact matches. A resolved match
means the wording was located; a later write can still fail authorization,
version, or edit-boundary checks.

When more than 10 exact matches exist, use the one-based `occurrence`
parameter to retrieve a specific later match.

Authentication and request validation failures apply to the whole request.
Missing, inaccessible, or otherwise unreadable media are reported as
per-media statuses without exposing whether an inaccessible ID exists.

## Requires api token with one of the following permissions
```
Read all folder and media data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used.

```
wistia captions find-matches [flags]
```

### Examples

```
  wistia captions find-matches --media-ids <value> --target-text <value>
```

### Options

```
      --body string             Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -e, --end-ms int              Optional end of a time range used to disambiguate the match.
  -h, --help                    help for find-matches
  -l, --language-code string    Exact IETF language tag. Omit when each media has only one caption track.
  -m, --media-ids stringArray   Explicit hashed IDs of the media whose captions should be searched. [required]
      --occurrence int          One-based exact occurrence to return, including occurrences after the first 10.
  -s, --start-ms int            Optional start of a time range used to disambiguate the match.
  -t, --target-text string      Exact caption wording to locate. [required]
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

* [wistia captions](wistia_captions.md)	 - Operations for captions

### Machine interface

* `wistia captions find-matches --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia captions find-matches --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
