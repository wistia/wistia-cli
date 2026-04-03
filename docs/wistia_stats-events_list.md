## wistia stats-events list

Stats:Events List

### Synopsis

Retrieve a list of events. Please note that due to our data retention policy,
only events from the last 2 years are available.

## Requires api token with one of the following permissions
```
Read, update & delete anything
Read all data
Read all folder and media data
```

```
wistia stats-events list [flags]
```

### Examples

```
  wistia stats-events list
```

### Options

```
  -e, --end-date string      End date in the format 'YYYY-MM-DD'.
  -h, --help                 help for list
  -m, --media-id string      An optional identifier for a specific video.
      --page int             The page of events to get data from.
      --per-page int         Maximum number of events to retrieve (capped at 100).
  -s, --start-date string    Start date in the format 'YYYY-MM-DD'.
  -v, --visitor-key string   An optional identifier for a specific visitor.
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

* [wistia stats-events](wistia_stats-events.md)	 - Operations for stats-events
