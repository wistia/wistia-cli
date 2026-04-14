## wistia remix post-remixes

Create Remix

### Synopsis

Start a new video remix job. The remix is processed asynchronously — poll
the show endpoint to check status and get preview URLs when ready.

Remix uses AI to analyze video transcripts and create edited versions
(highlight reels, trailers, cut-downs, etc.) based on your instructions.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read, update & delete anything
```
<!--- /HIDE-MCP -->

```
wistia remix post-remixes [flags]
```

### Examples

```
  wistia remix post-remixes --media-hashed-ids '["abc123","def456"]' --instructions Create a 60-second highlight reel focusing on the product demo section
```

### Options

```
      --body string                    Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                           help for post-remixes
  -i, --instructions string            Natural language instructions describing the desired remix (e.g., "create a 60-second highlight reel"). [required]
  -m, --media-hashed-ids stringArray   Array of source media hashed IDs to remix from. [required]
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

* [wistia remix](wistia_remix.md)	 - Operations for remix
