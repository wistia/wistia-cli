## wistia

Data API: Wistia Data API

### Synopsis

Data API: Wistia Data API

```
wistia [flags]
```

### Options

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDE_CODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview the request that would be sent without executing it (output to stderr)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                   help for wistia
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

* [wistia account](wistia_account.md)	 - Operations for account
* [wistia allowed-domains](wistia_allowed-domains.md)	 - Operations for allowed-domains
* [wistia auth](wistia_auth.md)	 - Manage authentication credentials
* [wistia background-job-status](wistia_background-job-status.md)	 - Operations for background-job-status
* [wistia captions](wistia_captions.md)	 - Operations for captions
* [wistia channel-episodes](wistia_channel-episodes.md)	 - Operations for channel-episodes
* [wistia channels](wistia_channels.md)	 - Operations for channels
* [wistia configure](wistia_configure.md)	 - Configure authentication credentials and preferences
* [wistia customizations](wistia_customizations.md)	 - Operations for customizations
* [wistia expiring-access-tokens](wistia_expiring-access-tokens.md)	 - Operations for expiring-access-tokens
* [wistia explore](wistia_explore.md)	 - Interactively browse and run commands
* [wistia live-stream-events](wistia_live-stream-events.md)	 - Operations for live-stream-events
* [wistia localizations](wistia_localizations.md)	 - Operations for localizations
* [wistia media](wistia_media.md)	 - Operations for media
* [wistia project-sharings](wistia_project-sharings.md)	 - Operations for project-sharings
* [wistia projects](wistia_projects.md)	 - Operations for projects
* [wistia search](wistia_search.md)	 - Operations for search
* [wistia stats-account](wistia_stats-account.md)	 - Operations for stats-account
* [wistia stats-events](wistia_stats-events.md)	 - Operations for stats-events
* [wistia stats-media](wistia_stats-media.md)	 - Operations for stats-media
* [wistia stats-projects](wistia_stats-projects.md)	 - Operations for stats-projects
* [wistia stats-visitors](wistia_stats-visitors.md)	 - Operations for stats-visitors
* [wistia subfolders](wistia_subfolders.md)	 - Operations for subfolders
* [wistia tags](wistia_tags.md)	 - Operations for tags
* [wistia trims](wistia_trims.md)	 - Operations for trims
* [wistia version](wistia_version.md)	 - Print the CLI version
* [wistia whoami](wistia_whoami.md)	 - Display current authentication configuration
