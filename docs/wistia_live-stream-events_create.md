## wistia live-stream-events create

Live Stream Event Create

### Synopsis

Create a new live stream event. The event will be created synchronously and return the event details with audience and host links.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

## Rate Limiting
This endpoint is rate limited to 60 requests per minute per IP address.

## Feature Access
This endpoint requires the live streaming feature to be enabled on your account plan.

```
wistia live-stream-events create [flags]
```

### Examples

```
  wistia live-stream-events create --title Wellness Session: Coping with Outie Memories --scheduled-for 2024-03-20T15:30:00-05:00 --event-duration 60
```

### Options

```
      --body string            Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string     The description of the live stream event
  -e, --event-duration int     Duration of the event in minutes (minimum 15) [required]
  -h, --help                   help for create
  -s, --scheduled-for string   The scheduled start time in W3C format with timezone [required]
  -t, --title string           The title of the live stream event [required]
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
