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
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
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
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-keyring             Never read or write the OS keychain; store secrets in the config file instead (env: WISTIA_CLI_NO_KEYRING)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia upload-or-import-media](wistia_upload-or-import-media.md)	 - Operations for upload-or-import-media

### Machine interface

* `wistia upload-or-import-media post-form --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia upload-or-import-media post-form --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
