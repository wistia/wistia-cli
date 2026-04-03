## wistia live-stream-events

Operations for live-stream-events

### Synopsis

Operations for live-stream-events

```
wistia live-stream-events [flags]
```

### Options

```
  -h, --help   help for live-stream-events
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

* [wistia](wistia.md)	 - Data API: Wistia Data API
* [wistia live-stream-events create](wistia_live-stream-events_create.md)	 - Live Stream Event Create
* [wistia live-stream-events delete](wistia_live-stream-events_delete.md)	 - Live Stream Event Delete
* [wistia live-stream-events get](wistia_live-stream-events_get.md)	 - Live Stream Event Show
* [wistia live-stream-events list](wistia_live-stream-events_list.md)	 - Live Stream Events List
* [wistia live-stream-events update](wistia_live-stream-events_update.md)	 - Live Stream Event Update
