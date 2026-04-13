## wistia media put-medias-copy

Bulk Copy Media

### Synopsis

This method accepts a list of medias to copy to a destination folder. It processes requests asynchronously and will return a background_job_status object rather than the typical Media response object.

Each media will be duplicated and the copy will be placed in the specified destination folder. The original media files will not be affected.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia media put-medias-copy [flags]
```

### Examples

```
  wistia media put-medias-copy --hashed-ids '["<value 1>"]' --folder-id <id>
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --folder-id string         The hashed ID of the destination folder where the copies will be placed. [required]
      --hashed-ids stringArray   An array of the media hashed IDs to be copied. [required]
  -h, --help                     help for put-medias-copy
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
