## wistia channel-episodes list

Channel Episodes List

### Synopsis

Returns all the Channel Episodes associated with the account, and allows for filtering by a particular channel.

## Requires api token with one of the following permissions
```
Read, update & delete anything
Read all data
Read all folder and media data
```

```
wistia channel-episodes list [flags]
```

### Examples

```
  wistia channel-episodes list
```

### Options

```
  -c, --channel-id string       Find episodes for a particular channel by providing the channel hashed ID
      --hashed-id stringArray   Filter by hashed id
  -h, --help                    help for list
  -m, --media-id stringArray    Filter by media id
      --page int                Page number to retrieve
      --per-page int            Number of channels per page
      --published               Filter by published status.
      --sort-by string          Ordering. Default is ID ASC. (options: position, title, created, updated, id)
      --sort-direction string   Ordering Sort Direction (0 = desc, 1 = asc; default is 1) (options: 0, 1)
  -t, --title string            Filter by channel episode name/title.
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
