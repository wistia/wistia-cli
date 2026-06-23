## wistia analytics-media get-embed-locations-timeseries

Show Media Embed Locations Timeseries

### Synopsis

Retrieve timeseries analytics for a video broken down by embed location. Returns
an array of timestamped buckets, each containing metrics for the top embed
locations (ranked by the chosen metric) plus an "All other" entry aggregating
the remaining locations.

The date range between `start_date` and `end_date` must not exceed 2 years.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read detailed stats
```
<!--- /HIDE-MCP -->

```
wistia analytics-media get-embed-locations-timeseries [flags]
```

### Examples

```
  wistia analytics-media get-embed-locations-timeseries --media-id <id> --start-date 2025-02-07 --end-date 2025-10-04 --granularity weekly
```

### Options

```
      --embed-url string     Filter results to a single embed URL. When provided, only analytics for
                             the page matching this URL are returned. Must be a valid HTTP or HTTPS URL.
                             
      --end-date string      End date for the analytics period in ISO 8601 format (YYYY-MM-DD). Exclusive — the range ends before the beginning of this date. [required]
  -g, --granularity string   The time granularity for the timeseries data. (options: daily, weekly, monthly) [required]
  -h, --help                 help for get-embed-locations-timeseries
  -m, --media-id string      The hashed ID of the video. [required]
  -p, --per-page int         Number of top embed locations per time bucket (max 100). Remaining locations are aggregated into an "All other" entry. (default 5)
      --sort-by string       The metric used to rank and select the top embed locations. (options: plays, loads, engagement_rate, play_rate) (default "plays")
      --start-date string    Start date for the analytics period in ISO 8601 format (YYYY-MM-DD). Inclusive — the range starts at the beginning of this date. [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia analytics-media](wistia_analytics-media.md)	 - Operations for analytics-media
