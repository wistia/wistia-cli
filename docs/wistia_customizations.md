## wistia customizations

Operations for customizations

### Synopsis

Operations for customizations

```
wistia customizations [flags]
```

### Options

```
  -h, --help   help for customizations
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

* [wistia](wistia.md)	 - Data API: Wistia Data API
* [wistia customizations create](wistia_customizations_create.md)	 - Create Customizations
* [wistia customizations delete](wistia_customizations_delete.md)	 - Delete Customizations
* [wistia customizations get](wistia_customizations_get.md)	 - Show Customizations
* [wistia customizations get-access](wistia_customizations_get-access.md)	 - Show Access Customizations
* [wistia customizations get-accessibility](wistia_customizations_get-accessibility.md)	 - Show Accessibility Customizations
* [wistia customizations get-appearance](wistia_customizations_get-appearance.md)	 - Show Appearance Customizations
* [wistia customizations get-chapters](wistia_customizations_get-chapters.md)	 - Show Chapters Customizations
* [wistia customizations get-engagement](wistia_customizations_get-engagement.md)	 - Show Engagement Customizations
* [wistia customizations get-lead-capture](wistia_customizations_get-lead-capture.md)	 - Show Lead Capture Customizations
* [wistia customizations get-playback](wistia_customizations_get-playback.md)	 - Show Playback Customizations
* [wistia customizations get-related-media](wistia_customizations_get-related-media.md)	 - Show Related Media Customizations
* [wistia customizations get-sharing](wistia_customizations_get-sharing.md)	 - Show Sharing Customizations
* [wistia customizations get-thumbnail](wistia_customizations_get-thumbnail.md)	 - Show Thumbnail Customizations
* [wistia customizations update](wistia_customizations_update.md)	 - Update Customizations
* [wistia customizations update-access](wistia_customizations_update-access.md)	 - Update Access Customizations
* [wistia customizations update-accessibility](wistia_customizations_update-accessibility.md)	 - Update Accessibility Customizations
* [wistia customizations update-appearance](wistia_customizations_update-appearance.md)	 - Update Appearance Customizations
* [wistia customizations update-chapters](wistia_customizations_update-chapters.md)	 - Update Chapters Customizations
* [wistia customizations update-engagement](wistia_customizations_update-engagement.md)	 - Update Engagement Customizations
* [wistia customizations update-lead-capture](wistia_customizations_update-lead-capture.md)	 - Update Lead Capture Customizations
* [wistia customizations update-playback](wistia_customizations_update-playback.md)	 - Update Playback Customizations
* [wistia customizations update-related-media](wistia_customizations_update-related-media.md)	 - Update Related Media Customizations
* [wistia customizations update-sharing](wistia_customizations_update-sharing.md)	 - Update Sharing Customizations
* [wistia customizations update-thumbnail](wistia_customizations_update-thumbnail.md)	 - Update Thumbnail Customizations

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
