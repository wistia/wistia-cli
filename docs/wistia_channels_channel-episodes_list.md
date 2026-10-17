## wistia channels channel-episodes list

List Channel Episodes by Channel

### Synopsis

Lists Channel Episodes belonging to the channel passed in the path.

## Requires api token with one of the following permissions
```
Read all folder and media data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

An [expiring access token](https://docs.wistia.com/reference/post_expiring-token)
created with the `all:delegate_to_contact_permissions` scope and an
authorization naming a channel (any permission) can also be used; it
lists the episodes of the channels the token names.

```
wistia channels channel-episodes list [channel-hashed-id] [flags]
```

### Examples

```
  wistia channels channel-episodes list --channel-hashed-id <id>
```

### Options

```
  -c, --channel-hashed-id string   The hashed ID of the channel to grab channel episodes from. (or pass it as the [channel-hashed-id] argument)
      --cursor string              If 'cursor[enabled]' is set to 1 then cursor pagination is enabled and the
                                   first set of records are fetched up to the 'per_page'. Cursor
                                   pagination will also be turned on if 'cursor[before]' or 'cursor[after]'
                                   are set. Records returned will have a 'cursor' property set which can be used to fetch more records in the same 'sort_by' ordering.
                                   The cursor value of the last record can be used to fetch records after the current result set and
                                   the cursor of the first record can be used to fetch records before the result set.
                                   
                                   NOTE: a cursor value is only valid if the 'sort_by' value hasn't changed from the
                                   last fetch. For example, you cannot fetch using 'sort_by' id and then pass that
                                   cursor value to a 'sort_by' name.
      --hashed-ids stringArray     Filter by hashed id
  -h, --help                       help for list
  -m, --media-id stringArray       Filter by media id. Accepts either the numeric id or the hashed id of a media.
      --page int                   The page number to retrieve. This cannot be combined with 'cursor',
                                   pagination.
      --per-page int               The number of medias per page. Use this for both offset pagination and cursor pagination.
      --published                  Filter by published status.
      --sort-by string             Ordering. Default is ID ASC. When using cursor pagination (see cursor param),
                                   only 'id' and 'created' are supported. All other sort_by options ('position', 'title', 'updated', 'published_at')
                                   require offset pagination.
                                   (options: position, title, created, updated, published_at, id)
      --sort-direction string      Ordering Sort Direction (0 = desc, 1 = asc; default is 1) (options: 0, 1)
  -t, --title string               Filter by channel episode name/title.
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

* [wistia channels channel-episodes](wistia_channels_channel-episodes.md)	 - Operations for channel-episodes

### Machine interface

* `wistia channels channel-episodes list --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia channels channel-episodes list --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
