## wistia upload-or-import-media post-form

Upload or Import Media

### Synopsis

Endpoint to upload media files from a local system or import from a web URL.

- Use `multipart/form-data` with a `file` parameter to upload from local system
- Use `application/x-www-form-urlencoded` with a `url` parameter to import from web URL

```
wistia upload-or-import-media post-form [flags]
```

### Examples

```
  wistia upload-or-import-media post-form --url http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4
```

### Options

```
  -a, --access-token string   A 64 character hex string. This parameter can be found on your API access page OR can be the token you received from authenticating via Oauth2. Note this approach is legacy and discouraged. We recommend using Bearer Token authentication.
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --contact-id int        A Wistia contact id.
      --description string    A description to use for the media in Wistia.
  -h, --help                  help for post-form
  -l, --low-priority          Inform the encoding service that this upload can be considered lower priority than others. This is especially useful for platform customers doing bulk uploads or migrations. Setting this to "false" has no effect.
  -n, --name string           A display name to use for the media in Wistia.
  -p, --project-id string     The hashed id of the project to upload media into.
  -u, --url string            The publicly accessible web location of the media file to import. [required]
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

* [wistia upload-or-import-media](wistia_upload-or-import-media.md)	 - Operations for upload-or-import-media
