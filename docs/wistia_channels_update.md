## wistia channels update

Update Channel

### Synopsis

Updates a channel.

```
wistia channels update [flags]
```

### Examples

```
  wistia channels update --channel-hashed-id <id>
```

### Options

```
  -a, --auto-publish-enabled       Whether the episodes are automatically published when added to the channel. Cannot be enabled if podcasting is on.
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --channel-hashed-id string   The hashed id of the Channel [required]
      --custom-url string          Use if embedding the channel on your own site. The custom URL ensures links always direct to your page and not Wistia's.
      --description string         The channel's description.
  -h, --help                       help for update
  -n, --name string                The display name for the channel
      --podcast-enabled            Whether podcasting is enabled for this channel.
      --podcast-settings string    Podcast specific settings for a channel. These settings only take effect if
                                   podcasting is enabled for the channel.
                                   
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

* [wistia channels](wistia_channels.md)	 - Operations for channels
