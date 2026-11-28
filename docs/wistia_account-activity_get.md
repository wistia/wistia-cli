## wistia account-activity get

List Account Activity

### Synopsis

Lists the account's activity, the same events shown on the Team Activity
page, newest first. Account Owners and Managers see activity for the whole
account. Other team members see only their own activity.

Results are always cursor paginated: pass the last record's `cursor` as
`cursor[after]` to fetch the next page.


## Requires api token with one of the following permissions
```
Read all folder and media data
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia account-activity get [flags]
```

### Examples

```
  wistia account-activity get
```

### Options

```
  -a, --actions stringArray     Only return activity with these actions. (options: create, update, destroy, publish, unpublish, moved, ordered, download, replace)
      --contact-id string       Only return activity by the team member with this hashed ID.
      --created-after string    Only return activity at or after this time: ISO 8601 at any precision ("2026", "2026-08", "2026-08-01", or "2026-08-01T09:30:00Z") or epoch seconds. A truncated value starts at the beginning of its year, month, or day, in UTC.
      --created-before string   Only return activity at or before this time, in the same forms as 'created_after'. A truncated value runs to the end of its year, month, or day, so 'created_after=2026-08' with 'created_before=2026-08' covers all of August 2026.
      --cursor string           Cursor pagination. When cursor pagination is enabled, each record returned
                                has a 'cursor': pass the last one as 'cursor[after]' for the next page, or
                                the first one as 'cursor[before]' for the previous page. A cursor only works with the 'sort_by' it was fetched with.
                                For example, you cannot fetch using 'sort_by' id and then pass that cursor
                                value to a 'sort_by' name.
                                
                                Setting 'cursor[enabled]' to 1 fetches the first page with cursor pagination
                                on, up to 'per_page' records. Setting 'cursor[before]' or 'cursor[after]' turns
                                it on too.
  -h, --help                    help for get
  -p, --per-page int            The number of activity records per page (default 10, max 100).
  -t, --types stringArray       Only return activity of these types. 'account' is changes to account
                                settings (name, URL, owner, billing email), not all account activity.
                                (options: account, archived_item, captions, channel, channel_episode, contact, folder, localization, media, media_assignment, media_workflow_status, subfolder, subscription_change, team_comment, trial, webinar)
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

* [wistia account-activity](wistia_account-activity.md)	 - Operations for account-activity

### Machine interface

* `wistia account-activity get --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia account-activity get --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
