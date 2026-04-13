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

* [wistia](wistia.md)	 - Data API: Wistia Data API
* [wistia channel-episodes create](wistia_channel-episodes_create.md)	 - Create Channel Episode
* [wistia channel-episodes delete-channel-episodes-channel-episode-hashed-id](wistia_channel-episodes_delete-channel-episodes-channel-episode-hashed-id.md)	 - Delete Channel Episode
* [wistia channel-episodes get](wistia_channel-episodes_get.md)	 - Show Channel Episode
* [wistia channel-episodes list](wistia_channel-episodes_list.md)	 - List Channel Episodes
* [wistia channel-episodes put-channel-episodes-channel-episode-hashed-id](wistia_channel-episodes_put-channel-episodes-channel-episode-hashed-id.md)	 - Update Channel Episode
* [wistia channel-episodes put-channel-episodes-channel-episode-hashed-id-publish](wistia_channel-episodes_put-channel-episodes-channel-episode-hashed-id-publish.md)	 - Publish Channel Episode
* [wistia channel-episodes put-channel-episodes-channel-episode-hashed-id-unpublish](wistia_channel-episodes_put-channel-episodes-channel-episode-hashed-id-unpublish.md)	 - Un-publish Channel Episode
