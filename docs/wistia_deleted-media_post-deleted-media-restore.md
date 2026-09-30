## wistia deleted-media post-deleted-media-restore

Restore Deleted Media

### Synopsis

Restores one or more soft-deleted media. By default each media returns to the
folder it was deleted from; pass folder_id to restore them into a specific
folder instead. Only media still inside the restore window can be recovered.
The restore runs asynchronously and the response includes a background job
status.


## Requires api token with one of the following permissions
```
Upload and view media
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia deleted-media post-deleted-media-restore [flags]
```

### Examples

```
  wistia deleted-media post-deleted-media-restore --media-hashed-ids abc123
```

### Options

```
      --body string                    Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -f, --folder-id string               Optional hashed id of the folder to restore the media into. If omitted, each media returns to the folder it was deleted from.
  -h, --help                           help for post-deleted-media-restore
  -m, --media-hashed-ids stringArray   The hashed ids of the soft-deleted media to restore. Up to 1000 at a time. [required]
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

* [wistia deleted-media](wistia_deleted-media.md)	 - Operations for deleted-media

### Machine interface

* `wistia deleted-media post-deleted-media-restore --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia deleted-media post-deleted-media-restore --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
