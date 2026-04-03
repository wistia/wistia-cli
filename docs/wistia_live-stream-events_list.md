## wistia live-stream-events list

Live Stream Events List

### Synopsis

Use this endpoint to request a list of Live Stream Events in your Wistia account. This request supports paging and sorting.

## Requires api token with one of the following permissions
```
Read, update & delete anything
Read all data
Read all folder and media data
```

```
wistia live-stream-events list [flags]
```

### Examples

```
  wistia live-stream-events list
```

### Options

```
      --hashed-ids stringArray   Filter by specific event IDs
  -h, --help                     help for list
      --page int                 Page number to retrieve
      --per-page int             Number of events per page (maximum 100)
      --sort-by string           Field to sort by (options: scheduled_for, id)
      --sort-direction string    Sort direction (1 for ascending, -1 for descending) (options: 1, -1)
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

* [wistia live-stream-events](wistia_live-stream-events.md)	 - Operations for live-stream-events
