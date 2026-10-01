## wistia

Data API: Wistia Data API

### Synopsis

Data API: Wistia Data API

```
wistia [flags]
```

### Options

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
  -h, --help                   help for wistia
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

* [wistia account](wistia_account.md)	 - Operations for account
* [wistia account-trials](wistia_account-trials.md)	 - Operations for account-trials
* [wistia allowed-domains](wistia_allowed-domains.md)	 - Operations for allowed-domains
* [wistia analytics-account](wistia_analytics-account.md)	 - Operations for analytics-account
* [wistia analytics-media](wistia_analytics-media.md)	 - Operations for analytics-media
* [wistia analytics-webinar](wistia_analytics-webinar.md)	 - Operations for analytics-webinar
* [wistia auth](wistia_auth.md)	 - Manage authentication credentials
* [wistia background-job-status](wistia_background-job-status.md)	 - Operations for background-job-status
* [wistia brands](wistia_brands.md)	 - Operations for brands
* [wistia bulk](wistia_bulk.md)	 - Operations for bulk
* [wistia bulk-actions](wistia_bulk-actions.md)	 - Operations for bulk-actions
* [wistia captions](wistia_captions.md)	 - Operations for captions
* [wistia channel-collaborators](wistia_channel-collaborators.md)	 - Operations for channel-collaborators
* [wistia channel-episodes](wistia_channel-episodes.md)	 - Operations for channel-episodes
* [wistia channels](wistia_channels.md)	 - Operations for channels
* [wistia configure](wistia_configure.md)	 - Configure authentication credentials and preferences
* [wistia contact](wistia_contact.md)	 - Operations for contact
* [wistia contacts](wistia_contacts.md)	 - Operations for contacts
* [wistia customizations](wistia_customizations.md)	 - Operations for customizations
* [wistia deleted-media](wistia_deleted-media.md)	 - Operations for deleted-media
* [wistia expiring-access-tokens](wistia_expiring-access-tokens.md)	 - Operations for expiring-access-tokens
* [wistia explore](wistia_explore.md)	 - Interactively browse and run commands
* [wistia folder-sharings](wistia_folder-sharings.md)	 - Operations for folder-sharings
* [wistia folders](wistia_folders.md)	 - Operations for folders
* [wistia localizations](wistia_localizations.md)	 - Operations for localizations
* [wistia media](wistia_media.md)	 - Operations for media
* [wistia media-extended-audio-descriptions](wistia_media-extended-audio-descriptions.md)	 - Operations for media-extended-audio-descriptions
* [wistia resource-urls](wistia_resource-urls.md)	 - Operations for resource-urls
* [wistia review-bundles](wistia_review-bundles.md)	 - Operations for review-bundles
* [wistia search](wistia_search.md)	 - Operations for search
* [wistia share-links](wistia_share-links.md)	 - Operations for share-links
* [wistia speakers](wistia_speakers.md)	 - Operations for speakers
* [wistia stats-account](wistia_stats-account.md)	 - Operations for stats-account
* [wistia stats-events](wistia_stats-events.md)	 - Operations for stats-events
* [wistia stats-media](wistia_stats-media.md)	 - Operations for stats-media
* [wistia stats-projects](wistia_stats-projects.md)	 - Operations for stats-projects
* [wistia stats-visitors](wistia_stats-visitors.md)	 - Operations for stats-visitors
* [wistia subfolders](wistia_subfolders.md)	 - Operations for subfolders
* [wistia taggings](wistia_taggings.md)	 - Operations for taggings
* [wistia tags](wistia_tags.md)	 - Operations for tags
* [wistia trims](wistia_trims.md)	 - Operations for trims
* [wistia upload-or-import-media](wistia_upload-or-import-media.md)	 - Operations for upload-or-import-media
* [wistia version](wistia_version.md)	 - Print the CLI version
* [wistia webinar-collaborators](wistia_webinar-collaborators.md)	 - Operations for webinar-collaborators
* [wistia webinar-registrations](wistia_webinar-registrations.md)	 - Operations for webinar-registrations
* [wistia webinars](wistia_webinars.md)	 - Operations for webinars
* [wistia whoami](wistia_whoami.md)	 - Display current authentication configuration

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
