## wistia channels channel-episodes list

List Channel Episodes by Channel

### Synopsis

Lists Channel Episodes belonging to the channel passed in the path.

## Requires api token with one of the following permissions
```
Read all folder and media data
```

```
wistia channels channel-episodes list [flags]
```

### Examples

```
  wistia channels-channel-episodes list --channel-hashed-id <id>
```

### Options

```
      --channel-hashed-id string   The hashed ID of the channel to grab channel episodes from. [required]
      --cursor cursor[enabled]     If cursor[enabled] is set to 1 then cursor pagination is enabled and the
                                   first set of records are fetched up to the `per_page`. Cursor
                                   pagination will also be turned on if `cursor[before]` or `cursor[after]`
                                   are set. Records returned will have a `cursor` property set which can be used to fetch more records in the same `sort_by` ordering.
                                   The cursor value of the last record can be used to fetch records after the current result set and
                                   the cursor of the first record can be used to fetch records before the result set.
                                   
                                   NOTE: a cursor value is only valid if the `sort_by` value hasn't changed from the
                                   last fetch. For example, you cannot fetch using `sort_by` id and then pass that
                                   cursor value to a `sort_by` name.
                                   
      --hashed-ids stringArray     Filter by hashed id
  -h, --help                       help for list
  -m, --media-id stringArray       Filter by media id. Accepts either the numeric id or the hashed id of a media.
      --page cursor                The page number to retrieve. This cannot be combined with cursor,
                                   pagination.
                                   
      --per-page int               The number of medias per page. Use this for both offset pagination and cursor pagination.
      --published                  Filter by published status.
      --sort-by id                 Ordering. Default is ID ASC. When using cursor pagination (see cursor param),
                                   only id and `created` are supported. All other sort_by options (`position`, `title`, `updated`)
                                   require offset pagination.
                                    (options: position, title, created, updated, id)
      --sort-direction string      Ordering Sort Direction (0 = desc, 1 = asc; default is 1) (options: 0, 1)
  -t, --title string               Filter by channel episode name/title.
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

* [wistia channels channel-episodes](wistia_channels_channel-episodes.md)	 - Operations for channels-channel-episodes
