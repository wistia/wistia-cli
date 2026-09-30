## wistia analytics-account find-media-by-embed-location

Find Media By Embed Location

### Synopsis

Find the media embedded at a given URL. Returns the hashed IDs of the
account's media that recorded activity at that embed location during the
date range, ranked by plays. The resulting hashed IDs can be passed to
other endpoints, such as Show Account Top Content's `hashed_ids[]` filter,
to fetch analytics for those media.

The domain of `embed_url` is always matched exactly. Its path is matched
exactly by default, or as a prefix with `path_match=prefix` (e.g.
`/pricing` also matching `/pricing/plans`). A path that is empty or `/`
is ignored, returning media across all paths on the domain.

Embed location data is retained for 6 months; a `start_date` older than
that returns a 422 error. When `start_date` and `end_date` are omitted,
the full 6-month queryable window is used.


## Requires api token with one of the following permissions
```
Read detailed stats
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia analytics-account find-media-by-embed-location [flags]
```

### Examples

```
  wistia analytics-account find-media-by-embed-location --embed-url https://milky-jury.com/
```

### Options

```
      --embed-url string    The URL of the page to look up, e.g. 'https://example.com/pricing'. The protocol is optional (https is assumed), so 'example.com/pricing' also works. [required]
      --end-date string     End date for the analytics period in ISO 8601 format (YYYY-MM-DD). Exclusive — the range ends before the beginning of this date. Defaults to tomorrow, so today's activity is included.
  -h, --help                help for find-media-by-embed-location
      --path-match string   How to match the path of 'embed_url' against embed locations. 'exact' requires the path to match exactly; 'prefix' matches any embed path starting with it. (options: exact, prefix) (default "exact")
      --per-page int        Number of media hashed IDs to return (max 1000). (default 100)
  -s, --start-date string   Start date for the analytics period in ISO 8601 format (YYYY-MM-DD). Inclusive — the range starts at the beginning of this date. Must be within the last 6 months. Defaults to 6 months ago, the start of the queryable window.
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

* [wistia analytics-account](wistia_analytics-account.md)	 - Operations for analytics-account

### Machine interface

* `wistia analytics-account find-media-by-embed-location --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia analytics-account find-media-by-embed-location --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
