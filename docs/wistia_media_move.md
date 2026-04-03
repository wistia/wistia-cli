## wistia media move

Media Move

### Synopsis

Move one or many media to a different project and optionally to a specific subfolder.
Max 100 media per request, and max 10 requests in 5 minutes.
Note: this is a different rate limit than applies to the rest of the api!

If a subfolder_id is provided, media will be moved to that subfolder. The subfolder
must belong to the specified project.

Returns a Background Job as the move is async.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia media move [flags]
```

### Examples

```
  wistia media move --hashed-ids '["<value 1>","<value 2>"]' --project-id <id>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --hashed-ids stringArray   An array of the media hashed IDs to be moved. [required]
  -h, --help                     help for move
  -p, --project-id string        The hashed ID of the project where you want the media moved. [required]
  -s, --subfolder-id string      Optional. The hashed ID of the subfolder where you want the media moved. If not provided, media will be moved to the project's default subfolder. The subfolder must belong to the specified project.
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
