## wistia media-extended-audio-descriptions post-media-extended-audio-descriptions-order

Order Extended Audio Description

### Synopsis

Orders an extended audio description for a media. The request will charge the credit card on the account when the order is ready.
Only accounts on paid plans with the `order_audio_descriptions` feature can use this endpoint.

```
wistia media-extended-audio-descriptions post-media-extended-audio-descriptions-order [flags]
```

### Examples

```
  wistia media-extended-audio-descriptions post-media-extended-audio-descriptions-order --media-id <id>
```

### Options

```
  -a, --ai-enabled                  Whether to use AI-generated audio descriptions (cheaper) or human-generated (higher quality). AI is only available for English orders. (default true)
      --body string                 Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -e, --enabled                     Whether the extended audio description should be automatically enabled once the order is complete. (default true)
  -h, --help                        help for post-media-extended-audio-descriptions-order
  -i, --ietf-language-tag string    IETF language tag for the audio description. Defaults to 'eng' (English).
                                    Non-English orders must set 'ai_enabled: false' — AI-generated audio
                                    descriptions are only available in English.
                                    
                                    Spanish ('es-419') orders are only accepted when the source media is
                                    tagged as a Spanish-language variant or has no detected language
                                    (e.g. silent videos). Spanish orders against a media in another
                                    language return '400'.
                                    (options: eng, es-419) (default "eng")
  -m, --media-id string             The hashed id of the media to order the extended audio description for. [required]
      --order-instructions string   Optional instructions for the audio description provider.
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

* [wistia media-extended-audio-descriptions](wistia_media-extended-audio-descriptions.md)	 - Operations for media-extended-audio-descriptions

### Machine interface

* `wistia media-extended-audio-descriptions post-media-extended-audio-descriptions-order --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia media-extended-audio-descriptions post-media-extended-audio-descriptions-order --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
