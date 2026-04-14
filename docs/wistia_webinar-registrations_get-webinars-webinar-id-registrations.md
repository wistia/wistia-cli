## wistia webinar-registrations get-webinars-webinar-id-registrations

List Webinar Registrations

### Synopsis

Retrieve a paginated list of registrations for a webinar. Returns contact
information, attendance status, engagement metrics, and attribution data
for each registrant.

Pagination uses cursor-based pagination with a `page_info` object in the
response rather than per-record cursors. Use `page_info.end_cursor` as
the `cursor` parameter to fetch the next page.

## Requires api token with one of the following permissions
```
Read all data
```

```
wistia webinar-registrations get-webinars-webinar-id-registrations [flags]
```

### Examples

```
  wistia webinar-registrations get-webinars-webinar-id-registrations --webinar-id <id>
```

### Options

```
  -a, --attendance string             Filter registrations by attendance status. (options: all, attendees, non_attendees) (default "all")
  -c, --cursor page_info.end_cursor   Cursor for pagination. Use the value from the previous response's page_info.end_cursor or `page_info.start_cursor`.
  -e, --emails stringArray            Filter registrations by email addresses.
  -h, --help                          help for get-webinars-webinar-id-registrations
  -p, --per-page int                  Number of results to return per page (max 100). (default 100)
  -r, --restriction string            Filter registrations by restriction status. (options: all, restricted, allowed) (default "all")
  -s, --sort-direction string         Sort direction (0 = desc/previous page, 1 = asc/next page; default is 1) (options: 0, 1) (default "1")
  -w, --webinar-id string             Hashed ID of the webinar. [required]
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

* [wistia webinar-registrations](wistia_webinar-registrations.md)	 - Operations for webinar-registrations
