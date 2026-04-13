## wistia stats-account get-stats-account-by-date

Show Account Stats by Date

### Synopsis

Retrieve account-wide stats organized by day, between a start and end date parameter (inclusive). If start and end date are not provided, defaults to yesterday and today.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read detailed stats
```
<!--- /HIDE-MCP -->

```
wistia stats-account get-stats-account-by-date [flags]
```

### Examples

```
  wistia stats-account get-stats-account-by-date
```

### Options

```
  -e, --end-date string     The end date for the stats, formatted YYYY-MM-DD
  -h, --help                help for get-stats-account-by-date
  -s, --start-date string   The start date for the stats, formatted YYYY-MM-DD
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

* [wistia stats-account](wistia_stats-account.md)	 - Operations for stats-account
