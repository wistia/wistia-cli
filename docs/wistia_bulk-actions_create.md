## wistia bulk-actions create

Create Bulk Actions

### Synopsis

Submits a batch of up to 1000 create, update, delete, and move actions to be
processed asynchronously. Returns a background job status whose Show
endpoint reports aggregate progress and per-action results, including the
hashed IDs of created records.

Supported resource types are `media`, `folder`, `subfolder`, `channel`,
`channel_episode`, `captions`, and the ten `customization_*` concerns. A
`folder` is a top-level folder (previously called a project); a `subfolder`
is nested inside one and requires `folder_id` and `name` when created. A
`captions` action operates on one caption track -- one media in one
language.

Because caption actions carry SRT contents inline, they are the resource
type most likely to reach the request body limit before the action cap.
Purchasing captions is not available here -- it has its own endpoint.

A `move` action targets one media and accepts a destination `folder_id` and
optional `subfolder_id`. Bulk moves can use different destinations and are
not subject to the Move Media endpoint's 100-item limit or separate throttle.

Player customizations are addressed one concern at a time
(`customization_appearance`, `customization_playback`, and so on), matching
the Update Customizations endpoints; each accepts `update` only, takes the
media's hashed ID as its `id`, and takes the same payload as its
corresponding endpoint. There is no batch equivalent of the broad customize
endpoint, so a batch always states which slice of the player it is changing.

A `media` update payload can also carry a `custom_metadata` object mapping
field keys to the values to set (`null` clears a field; omitted fields are
left untouched). Values are validated against each field's type exactly as
the Set Custom Metadata Field Value endpoint validates them, and each write
is recorded with its actor and source. Requires the custom metadata feature
on the account; without it actions carrying `custom_metadata` fail
individually.

Deleting a folder or subfolder also soft-deletes its media. An account owner
or manager can restore that media from the trash until it purges. To keep the
media when deleting a subfolder, use the Delete Subfolder endpoint; it moves
the media to the folder's root level instead.

Each action in the batch is authorized and processed independently:
failures (including authorization failures) are reported per action and do
not prevent other actions from completing. Media creation is not supported
-- uploads and URL imports have their own endpoints.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia bulk-actions create [flags]
```

### Examples

```
  wistia bulk-actions create
```

### Options

```
  -a, --actions string   An array of actions to process, one per record. Maximum 1000 actions per
                         request, and the request body must stay under 2 MB -- whichever limit is
                         reached first. An oversized body is rejected with a '413' and no action
                         in it runs. Each action specifies an operation (create, update, delete,
                         or move), a resource type, and the relevant payload or record ID.
                         
                         Use 'job' instead when every record takes the same payload.
                         (JSON array)
      --body string      Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -h, --help             help for create
  -j, --job string       One change applied to many records, named by a parent ('scope') or listed
                         explicitly ('ids'). The server resolves the target and runs one action per
                         record, so a folder of 400 media takes one job rather than 400 actions.
                         
                         A 'scope' resolves to exactly what the matching list endpoint returns for
                         that parent, including its defaults -- so a 'folder' scope on 'media' reaches
                         media in that folder's subfolders, and includes **archived** media.
                         
                         A job resolves to at most 5000 records. Beyond that it is rejected rather
                         than truncated, so a job never silently acts on part of the set you named --
                         narrow the scope, or send the records as an actions array.
                         
                         Cannot be used with 'create', which has no record to address, and is not
                         available to external contacts.
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

* [wistia bulk-actions](wistia_bulk-actions.md)	 - Operations for bulk-actions

### Machine interface

* `wistia bulk-actions create --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia bulk-actions create --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
