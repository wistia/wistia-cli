## wistia channel-episodes update

Update Channel Episode

### Synopsis

Updates an existing channel episode in a channel.
## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

```
wistia channel-episodes update [channel-episode-hashed-id] [flags]
```

### Examples

```
  wistia channel-episodes update --channel-episode-hashed-id <id>
```

### Options

```
      --body string                          Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -c, --channel-episode-hashed-id string     The hashed id of the Channel Episode (or pass it as the [channel-episode-hashed-id] argument)
      --description string                   The episode's description or episode notes.
  -e, --episode-notes string                 Additional notes for the episode.
  -h, --help                                 help for update
  -l, --live-stream-event-hashed-id string   The unique alphanumeric identifier for the live stream event associated with this channel episode.
  -m, --media-hashed-id string               The unique alphanumeric identifier for the media associated with this channel episode.
      --podcast-settings string              Podcast specific settings for a channel episode. These settings only take effect
                                             if podcasting is enabled for the channel.
      --publish-at string                    The date and time when the episode is scheduled to be published in UTC timezone.
      --publish-status string                The status of whether or not the episode has been published to your channel. (options: draft, published, scheduled)
  -s, --summary string                       A short summary of the episode that is displayed when space is limited.
  -t, --title string                         The episode's title. If not provided, the channel episode uses the title of the media used to create it.
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

* [wistia channel-episodes](wistia_channel-episodes.md)	 - Operations for channel-episodes

### Machine interface

* `wistia channel-episodes update --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia channel-episodes update --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
