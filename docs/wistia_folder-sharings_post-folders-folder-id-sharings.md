## wistia folder-sharings post-folders-folder-id-sharings

Create Folder Sharing

### Synopsis

Creates a new sharing object for a folder by specifying the email of the person to share with and other optional parameters.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia folder-sharings post-folders-folder-id-sharings [flags]
```

### Examples

```
  wistia folder-sharings post-folders-folder-id-sharings --folder-id <id> --sharing '{"with":"Katrine83@hotmail.com","requirePassword":false,"canShare":false,"canDownload":false,"canUpload":false}'
```

### Options

```
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --folder-id string   Hashed ID of the folder to be shared [required]
  -h, --help               help for post-folders-folder-id-sharings
  -s, --sharing string     [required]
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

* [wistia folder-sharings](wistia_folder-sharings.md)	 - Operations for folder-sharings
