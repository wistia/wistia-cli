## wistia media list

List Media

### Synopsis

Lists the media belonging to the account. This endpoint can also be used to
do a batch fetch based off of the hashed id.

## Requires api token with one of the following permissions
```
Read all folder and media data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia media list [flags]
```

### Examples

```
  wistia media list
```

### Options

```
  -a, --archived                          Filter by archived status. True will return only archived medias, while false will return only active medias.
  -c, --cursor string                     If 'cursor[enabled]' is set to 1 then cursor pagination is enabled and the
                                          first set of records are fetched up to the 'per_page'. Cursor
                                          pagination will also be turned on if 'cursor[before]' or 'cursor[after]'
                                          are set. Records returned will have a 'cursor' property set which can be used to fetch more records in the same 'sort_by' ordering.
                                          The cursor value of the last record can be used to fetch records after the current result set and
                                          the cursor of the first record can be used to fetch records before the result set.
                                          
                                          NOTE: a cursor value is only valid if the 'sort_by' value hasn't changed from the
                                          last fetch. For example, you cannot fetch using 'sort_by' id and then pass that
                                          cursor value to a 'sort_by' name.
  -f, --folder-id string                  A hashed ID specifying the folder from which you would like to get results.
      --hashed-ids stringArray            Find all of the medias by these hashed_ids.
  -h, --help                              help for list
  -i, --include string                    Set to 'speakers' to add each media's 'speakers' array, its active transcript speaker assignments used for diarization. Webinar hosts and panelists are not included. (options: speakers)
  -n, --name string                       Find a media or medias whose name exactly matches this parameter.
      --page int                          The page number to retrieve. This cannot be combined with 'cursor',
                                          pagination.
      --per-page int                      The number of medias per page. Use this for both offset pagination and cursor pagination.
      --sort-by string                    Ordering. When using cursor pagination (see cursor param),
                                          only 'id' and 'created' are supported. All other sort_by options ('name', 'updated', 'position')
                                          require offset pagination.
                                          (options: name, created, updated, position)
      --sort-direction string             Ordering Sort Direction (0 = desc, 1 = asc; default is 1) (options: 0, 1)
      --speaker-profile-ids stringArray   Filter media assigned to any of these reusable speaker profiles. Blank IDs are ignored; a list with no ID returns 400.
      --tags stringArray                  Find all of the medias that match all of these tag names.
      --type string                       A string specifying which type of media you would like to get. (options: Video, Audio, Image, PdfDocument, MicrosoftOfficeDocument, Swf, UnknownType)
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

* [wistia media](wistia_media.md)	 - Operations for media

### Machine interface

* `wistia media list --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia media list --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
