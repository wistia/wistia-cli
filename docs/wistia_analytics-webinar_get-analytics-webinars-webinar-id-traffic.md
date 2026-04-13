## wistia analytics-webinar get-analytics-webinars-webinar-id-traffic

Show Webinar Traffic Breakdown

### Synopsis

Retrieve traffic breakdown analytics for a webinar, grouped by a specified dimension
such as UTM campaign, UTM source, UTM medium, or referrer domain.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read detailed stats
```
<!--- /HIDE-MCP -->

```
wistia analytics-webinar get-analytics-webinars-webinar-id-traffic [flags]
```

### Examples

```
  wistia analytics-webinar get-analytics-webinars-webinar-id-traffic --webinar-id <id> --group-by utm_campaign
```

### Options

```
  -g, --group-by string         The dimension to group traffic data by. (options: utm_campaign, utm_source, utm_medium, referrer_domain) [required]
  -h, --help                    help for get-analytics-webinars-webinar-id-traffic
      --sort-by string          The metric to sort results by. (options: registrations, attendees, impressions) (default "registrations")
      --sort-direction string   The sort direction. (options: asc, desc) (default "desc")
  -w, --webinar-id string       The hashed ID of the webinar. [required]
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
