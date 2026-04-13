## wistia search search

Search

### Synopsis

Searches across folders, subfolders, medias, channels, channel episodes, and webinars.
Also searches through video transcripts, so media results may include transcript matches with
timestamps when the query matches spoken content.

## Requires api token with one of the following permissions
```
Read all data
```

```
wistia search search [flags]
```

### Examples

```
  wistia search search --q screencast
```

### Options

```
      --created-after string        Filter results created on or after this datetime. Must be a valid ISO8601 timestamp in UTC (ending with 'Z').
      --created-before string       Filter results created on or before this datetime. Must be a valid ISO8601 timestamp in UTC (ending with 'Z').
  -h, --help                        help for search
      --q string                    The search query string [required]
  -r, --resource-type stringArray   Filter results by one or more resource types.
  -t, --tags stringArray            Filter results by one or more tag names. When multiple tags are provided, results matching any of the specified tags are returned (OR logic).
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

* [wistia search](wistia_search.md)	 - Operations for search
