## wistia channel-episodes

Operations for channel-episodes

### Synopsis

Operations for channel-episodes

```
wistia channel-episodes [flags]
```

### Options

```
  -h, --help   help for channel-episodes
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
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia](wistia.md)	 - Data API: Wistia Data API
* [wistia channel-episodes create](wistia_channel-episodes_create.md)	 - Create Channel Episode
* [wistia channel-episodes delete](wistia_channel-episodes_delete.md)	 - Delete Channel Episode
* [wistia channel-episodes get](wistia_channel-episodes_get.md)	 - Show Channel Episode
* [wistia channel-episodes list](wistia_channel-episodes_list.md)	 - List Channel Episodes
* [wistia channel-episodes publish](wistia_channel-episodes_publish.md)	 - Publish Channel Episode
* [wistia channel-episodes unpublish](wistia_channel-episodes_unpublish.md)	 - Un-publish Channel Episode
* [wistia channel-episodes update](wistia_channel-episodes_update.md)	 - Update Channel Episode

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
