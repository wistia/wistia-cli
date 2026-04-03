## wistia media copy

Media Copy

### Synopsis

Copy a media.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia media copy [flags]
```

### Examples

```
  wistia media copy --media-hashed-id <id>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                     help for copy
  -m, --media-hashed-id string   The hashed ID of the media. [required]
      --owner string             An email address specifying the owner of the new media. Defaults to the source media’s current owner if omitted or invalid.
  -p, --project-id int           The ID of the project where you want the new copy placed. Defaults to the source media’s current project if omitted or invalid.
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

* [wistia media](wistia_media.md)	 - Operations for media
