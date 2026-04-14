## wistia captions get

Show Captions

### Synopsis

Returns a video's captions in the specified language.
Supports multiple formats: JSON (default), SRT, VTT, and TXT.
Use file extensions (.srt, .vtt, .txt) or Accept headers to specify format.

## Requires api token with one of the following permissions
```
Read all folder and media data
```

```
wistia captions get [flags]
```

### Examples

```
  wistia captions get --media-hashed-id <id> --language-code <value>
```

### Options

```
  -h, --help                     help for get
  -l, --language-code eng        The 3-character ISO 639-2 language code of the captions to be retrieved (e.g., eng, `fra`, `spa`). Some languages use extended IETF subtags (e.g., `zh-Hant`). [required]
  -m, --media-hashed-id string   The hashed ID of the media from which captions are to be retrieved. [required]
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

* [wistia captions](wistia_captions.md)	 - Operations for captions
