## wistia search search

Search

### Synopsis

Searches across folders, subfolders, medias, channels, channel episodes, and webinars.
Also searches through video transcripts, so media results may include transcript matches with
timestamps when the query matches spoken content.

## Requires api token with one of the following permissions
```
Read all data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia search search [flags]
```

### Examples

```
  wistia search search --q screencast
```

### Options

```
      --created-after string        Filter results created on or after this datetime. Must be a valid ISO8601 timestamp in UTC (ending with 'Z').
      --created-before string       Filter results created on or before this datetime. Must be a valid ISO8601 timestamp in UTC (ending with 'Z').
      --custom-metadata string      Filter media by custom metadata field value, keyed by field key:
                                    'custom_metadata[<field_key>]=<value>'. Only available on accounts with access to
                                    custom metadata (other accounts receive a 403 when this parameter is passed).
                                    Custom metadata only exists on media, so results contain media only and
                                    'resource_type' must include 'media'. Use an empty 'q' to match all media.
                                    
                                    The value shape depends on the field's type:
                                    - Select and text fields take a value
                                      ('custom_metadata[region]=emea') or an array of values matched as OR
                                      ('custom_metadata[region][]=emea&custom_metadata[region][]=amer'). Select fields
                                      match on option keys.
                                    - Boolean fields take 'true' or 'false'.
                                    - Number and time fields take an exact number ('custom_metadata[year]=2026')
                                      or a range object ('custom_metadata[budget][min]=100&custom_metadata[budget][max]=500';
                                      either bound may be omitted).
                                    - Date and datetime fields take a 'YYYY-MM-DD' date matching that UTC day, or a
                                      range object with ISO8601 bounds ('custom_metadata[shoot_date][after]=2026-01-01',
                                      'custom_metadata[shoot_date][before]=2026-02-01T00:00:00Z'). A bare-date bound
                                      covers its whole UTC day: 'after' starts at the day's beginning and 'before'
                                      runs through the day's end.
                                    - Any field type accepts a presence filter: 'custom_metadata[region][exists]=false'
                                      returns media missing the field entirely (useful for metadata coverage audits),
                                      and 'exists=true' returns media that have any value for it.
                                    
                                    Unknown or archived field keys return a 400, as do select option keys that don't
                                    exist on the field. The primary match set holds at most 100 media with no
                                    pagination; a non-blank 'q' can add up to 100 more transcript-only matches, and
                                    an empty-'q' audit returns at most 100. Narrow large audits (e.g. with
                                    'created_after'/'created_before') to complete full coverage.
  -h, --help                        help for search
  -i, --include string              Pass 'custom_metadata' to include each media result's custom metadata field values (same shape as the Get Custom Metadata Field Values endpoint). Only available on accounts with access to custom metadata (other accounts receive a 403 when this parameter is passed). (options: custom_metadata)
      --q string                    The search query string [required]
  -r, --resource-type stringArray   Filter results by one or more resource types. (options: media, folder, subfolder, channel, channel_episode, webinar)
  -t, --tags stringArray            Filter results by one or more tag names. When multiple tags are provided, results matching any of the specified tags are returned (OR logic).
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
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia search](wistia_search.md)	 - Operations for search

### Machine interface

* `wistia search search --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia search search --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
