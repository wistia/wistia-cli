## wistia channel-episodes create

Create Channel Episode

### Synopsis

Creates a new channel episode in a channel.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia channel-episodes create [flags]
```

### Examples

```
  wistia channel-episodes create --channel-hashed-id <id>
```

### Options

```
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -c, --channel-hashed-id string   The hashed ID of the channel to add the episode to. [required]
      --description string         The episode's description or episode notes.
      --episode-number int         The episode number for this episode in your podcast.  This parameter only takes effect if podcasting is enabled for the channel.
      --episode-type string        The episode type for your podcast.  This parameter only takes effect if podcasting is enabled for the channel. (options: full, trailer, bonus)
      --explicit-content           Whether this episode contains explicit content.  This parameter only takes effect if podcasting is enabled for the channel.
  -h, --help                       help for create
      --hide-from-feed             Whether or not to hide this episode from your podcast feed.  Set to true to hide the episode, false to show the episode.  This parameter only takes effect if podcasting is enabled for the channel.
  -m, --media-id string            The alphanumeric hashed ID of the media to be added as a channel episode.
      --publish-at string          The date and time when the episode should be published in UTC timezone. Required when publish_status is 'scheduled'. Must be a valid ISO8601 timestamp in UTC (ending with 'Z').  Can only be provided when publish_status is 'scheduled.'
      --publish-status string      The status of whether or not the episode has been published to your channel. (options: draft, published, scheduled)
  -s, --summary string             A short summary of the episode that is displayed when space is limited.
  -t, --title string               The episode's title.  If not provided, the channel episode uses the title of the media used to create it.
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
