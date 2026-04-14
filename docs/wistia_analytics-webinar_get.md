## wistia analytics-webinar get

Show Webinar Analytics

### Synopsis

Retrieve aggregate analytics for a webinar. This endpoint provides
Bottler-powered analytics including registrations, attendance, engagement,
chat activity, and poll results.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read detailed stats
```
<!--- /HIDE-MCP -->

```
wistia analytics-webinar get [flags]
```

### Examples

```
  wistia analytics-webinar get --webinar-id <id>
```

### Options

```
  -h, --help                           help for get
  -i, --include-post-event             Whether to include on-demand viewing data after the live event ended.
      --post-event-end-date string     End date for the post-event analytics period in ISO 8601 format (YYYY-MM-DD). Exclusive — the range ends before the beginning of this date. Only used when include_post_event is true.
      --post-event-start-date string   Start date for the post-event analytics period in ISO 8601 format (YYYY-MM-DD). Inclusive — the range starts at the beginning of this date. Only used when include_post_event is true.
  -w, --webinar-id string              The hashed ID of the webinar. [required]
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

* [wistia analytics-webinar](wistia_analytics-webinar.md)	 - Operations for analytics-webinar
