## wistia account update-transcript-settings

Update Transcript Settings

### Synopsis

Updates the account's keyword glossary. Provide keywords to replace the
complete list, or add_keywords and/or remove_keywords to edit it.
An empty keywords array clears the glossary. Omitted, null, and empty
add_keywords or remove_keywords mean no action. keywords cannot be combined
with a nonempty add_keywords or remove_keywords, and a term cannot appear in
both. Whitespace is squished and duplicate terms are processed once.
Matching is case-sensitive.

The resulting glossary must fit 100 terms and 500 characters, excluding
separators. Invalid requests return 400 without saving any changes.
Additions and removals preserve unrelated terms. A full replacement can
overwrite changes made since you last read the glossary.
Only the glossary is writable. Returns the same
representation as Get Transcript Settings.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token. That contact must have permission to update
the account's settings.

An [expiring access token](https://docs.wistia.com/reference/post_expiring-token)
created with the `all:delegate_to_contact_permissions` scope and an
`account` authorization granting the `manage-transcript-settings`
permission can also be used.

```
wistia account update-transcript-settings [flags]
```

### Examples

```
  wistia account update-transcript-settings
```

### Options

```
  -a, --add-keywords string      Terms to add while keeping existing entries. Terms already present are
                                 unchanged. Omitted, null, and empty all mean no action.
                                 (JSON array)
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF. Use --schema to print the exact JSON Schema.
  -h, --help                     help for update-transcript-settings
  -k, --keywords string          Replaces the entire glossary. To clear it, send 'keywords: []'.
                                 To edit individual terms, omit 'keywords' or set it to 'null' and
                                 provide terms in 'add_keywords', 'remove_keywords', or both.
                                 (JSON array)
  -r, --remove-keywords string   Terms to remove while keeping other entries. Absent terms are
                                 unchanged. Omitted, null, and empty all mean no action.
                                 (JSON array)
      --schema                   Print the exact JSON Schema of the request body and exit
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

* [wistia account](wistia_account.md)	 - Operations for account

### Machine interface

* `wistia account update-transcript-settings --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia account update-transcript-settings --schema` — the exact JSON Schema of the request body (all `$ref`s bundled)
* `wistia account update-transcript-settings --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
