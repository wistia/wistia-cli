## wistia channels update

Update Channel

### Synopsis

Updates a channel.

```
wistia channels update [channel-hashed-id] [flags]
```

### Examples

```
  wistia channels update --channel-hashed-id <id>
```

### Options

```
  -a, --auto-publish-enabled       Whether the episodes are automatically published when added to the channel. Cannot be enabled if podcasting is on.
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -c, --channel-hashed-id string   The hashed id of the Channel (or pass it as the [channel-hashed-id] argument)
      --custom-url string          Use if embedding the channel on your own site. The custom URL ensures links always direct to your page and not Wistia's.
      --description string         The channel's description.
  -h, --help                       help for update
  -n, --name string                The display name for the channel
      --podcast-enabled            Whether podcasting is enabled for this channel.
      --podcast-settings string    Podcast specific settings for a channel. These settings only take effect if
                                   podcasting is enabled for the channel. These values appear in the channel's
                                   publicly accessible podcast RSS feed.
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

* [wistia channels](wistia_channels.md)	 - Operations for channels

### Machine interface

* `wistia channels update --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia channels update --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
