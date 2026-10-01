## wistia analytics-webinar get-registration

Show Webinar Registration Timeseries

### Synopsis

Retrieve registration timeseries data for a webinar with configurable
granularity. Returns an array of timestamped registration metric buckets
including impressions, registrations, and completion rates.


## Requires api token with one of the following permissions
```
Read detailed stats
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia analytics-webinar get-registration [webinar-id] [flags]
```

### Examples

```
  wistia analytics-webinar get-registration --webinar-id <id> --granularity monthly
```

### Options

```
  -g, --granularity string             The time granularity for the timeseries data. (options: daily, weekly, monthly) [required]
  -h, --help                           help for get-registration
  -i, --include-post-event             Whether to include on-demand viewing data after the live event ended.
      --post-event-end-date string     End date for the post-event analytics period in ISO 8601 format (YYYY-MM-DD). Exclusive — the range ends before the beginning of this date. Only used when include_post_event is true.
      --post-event-start-date string   Start date for the post-event analytics period in ISO 8601 format (YYYY-MM-DD). Inclusive — the range starts at the beginning of this date. Only used when include_post_event is true.
  -w, --webinar-id string              The hashed ID of the webinar. (or pass it as the [webinar-id] argument)
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

* [wistia analytics-webinar](wistia_analytics-webinar.md)	 - Operations for analytics-webinar

### Machine interface

* `wistia analytics-webinar get-registration --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia analytics-webinar get-registration --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
