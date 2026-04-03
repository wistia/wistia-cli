## wistia live-stream-events update

Live Stream Event Update

### Synopsis

Update an existing live stream event.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia live-stream-events update [flags]
```

### Examples

```
  wistia live-stream-events update --id <id>
```

### Options

```
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                       help for update
  -i, --id string                  The hashed ID of the live stream event [required]
  -l, --live-stream-event string   JSON object
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
