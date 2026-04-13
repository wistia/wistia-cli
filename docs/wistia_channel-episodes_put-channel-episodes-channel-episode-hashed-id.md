## wistia channel-episodes put-channel-episodes-channel-episode-hashed-id

Update Channel Episode

### Synopsis

Updates an existing channel episode in a channel.
## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia channel-episodes put-channel-episodes-channel-episode-hashed-id [flags]
```

### Examples

```
  wistia channel-episodes put-channel-episodes-channel-episode-hashed-id --channel-episode-hashed-id <id>
```

### Options

```
      --body string                          Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --channel-episode-hashed-id string     The hashed id of the Channel Episode [required]
      --description string                   The episode's description or episode notes.
      --episode-notes string                 Additional notes for the episode.
      --episode-number int                   The number of the episode. This parameter only takes effect if podcasting is enabled for the channel.
      --episode-type string                  The type of episode. This parameter only takes effect if podcasting is enabled for the channel. (options: full, trailer, bonus)
      --explicit-content                     Whether the episode contains explicit content. This parameter only takes effect if podcasting is enabled for the channel.
  -h, --help                                 help for put-channel-episodes-channel-episode-hashed-id
      --hide-from-feed                       Whether to hide the episode from the podcast feed. This parameter only takes effect if podcasting is enabled for the channel.
  -l, --live-stream-event-hashed-id string   The unique alphanumeric identifier for the live stream event associated with this channel episode.
  -m, --media-hashed-id string               The unique alphanumeric identifier for the media associated with this channel episode.
      --publish-at string                    The date and time when the episode is scheduled to be published in UTC timezone.
      --publish-status string                The status of whether or not the episode has been published to your channel. (options: draft, published, scheduled)
      --season-number int                    The season number of the episode. This parameter only takes effect if podcasting is enabled for the channel.
      --summary string                       A short summary of the episode that is displayed when space is limited.
  -t, --title string                         The episode's title. If not provided, the channel episode uses the title of the media used to create it.
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

* [wistia channel-episodes](wistia_channel-episodes.md)	 - Operations for channel-episodes
