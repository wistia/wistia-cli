## wistia channels list

Channels List

### Synopsis

Returns all the Channels associated with the account.

## Requires api token with one of the following permissions
```
Read, update & delete anything
Read all data
Read all folder and media data
```

```
wistia channels list [flags]
```

### Examples

```
  wistia channels list
```

### Options

```
      --hashed-ids stringArray   Find all of the channels limited to these hashed_ids.
  -h, --help                     help for list
      --page int                 Page number to retrieve
      --per-page int             Number of channels per page
      --sort-by string           Ordering. Default is ID ASC. (options: created, id, updated, name)
      --sort-direction string    Ordering Sort Direction (0 = desc, 1 = asc; default is 1) (options: 0, 1)
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
