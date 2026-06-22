## wistia webinars create

Create Webinar

### Synopsis

Creates a new webinar.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia webinars create [flags]
```

### Examples

```
  wistia webinars create --title Wellness Session: Coping with Outie Memories --scheduled-for 2024-03-20T15:30:00-05:00 --event-duration 60 --time-zone America/New_York
```

### Options

```
      --body string          Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --description string   The description of the webinar
  -e, --event-duration int   Duration of the event in minutes (minimum 15) [required]
  -f, --folder-id string     Hashed ID of the folder to place this webinar in. Defaults to the account's default webinar folder if not provided.
  -h, --help                 help for create
  -s, --scheduled-for Z      The scheduled start time as a UTC formatted ISO 8601 string (offset Z or `+00:00`). [required]
      --time-zone string     The IANA time zone identifier the webinar is scheduled in. [required]
      --title string         The title of the webinar [required]
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

* [wistia webinars](wistia_webinars.md)	 - Operations for webinars
