## wistia media import-url

Import Media from URL

### Synopsis

This endpoint imports a media file from a given URL. The import is processed
asynchronously and will return a background_job_status object rather than the
typical Media response object. You can poll the background job status endpoint
to check on the progress of the import.

If no folder_id is provided, a new folder called "Untitled Folder" will be
created and the imported media will be placed there.

The URL must be publicly accessible — Wistia's servers need to be able to fetch the file directly.

Note: imports from certain domains (e.g. vimeo.com, wistia.com) are not permitted.

<!--- HIDE-MCP -->
## Requires api token with one of the following permissions
```
Read, update & delete anything
```
<!--- /HIDE-MCP -->

```
wistia media import-url [flags]
```

### Examples

```
  wistia media import-url --url https://example.com/video.mp4
```

### Options

```
      --body string        Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -f, --folder-id string   The hashed ID of the folder (project) to import the media into. If not provided, a new folder will be created.
  -h, --help               help for import-url
  -u, --url string         The publicly accessible URL of the media file to import. [required]
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
