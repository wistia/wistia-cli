## wistia remix post-remixes-remix-hashed-id-continue

Continue Remix

### Synopsis

Submit a follow-up edit to an existing remix. Creates a new remix version
in the same conversation. The previous remix is preserved and can be
referenced later.

The new remix version is automatically exported (rendered) upon completion.
If `folder_id` is provided, the output is exported to that folder. Otherwise,
it defaults to the same folder as the previous remix version's output.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia remix post-remixes-remix-hashed-id-continue [remix-hashed-id] [flags]
```

### Examples

```
  wistia remix post-remixes-remix-hashed-id-continue --remix-hashed-id <id> --instructions 'Cut the intro and add background music'
```

### Options

```
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -f, --folder-id string         Hashed ID of the destination folder for the exported media. Defaults to the same folder as the previous remix version's output.
  -h, --help                     help for post-remixes-remix-hashed-id-continue
  -i, --instructions string      Natural language instructions for the edit (e.g., "cut the first 10 seconds"). [required]
  -r, --remix-hashed-id string   The hashed ID of the current remix version to edit from. (or pass it as the [remix-hashed-id] argument)
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

* [wistia remix](wistia_remix.md)	 - Operations for remix

### Machine interface

* `wistia remix post-remixes-remix-hashed-id-continue --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia remix post-remixes-remix-hashed-id-continue --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
