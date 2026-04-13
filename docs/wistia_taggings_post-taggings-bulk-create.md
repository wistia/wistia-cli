## wistia taggings post-taggings-bulk-create

Bulk Tag Media

### Synopsis

This method accepts a list of medias to tag. It processes requests asynchronously and will return a background_job_status object rather than the typical Media response object.

The tags will be added to the existing tags on each media file, not replaced.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia taggings post-taggings-bulk-create [flags]
```

### Examples

```
  wistia taggings post-taggings-bulk-create --hashed-ids '["<value 1>","<value 2>","<value 3>"]' --tag-names '["<value 1>"]'
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --hashed-ids stringArray   An array of the media hashed IDs to be tagged. [required]
  -h, --help                     help for post-taggings-bulk-create
  -t, --tag-names stringArray    An array of tag names to add to each media. [required]
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

* [wistia taggings](wistia_taggings.md)	 - Operations for taggings
