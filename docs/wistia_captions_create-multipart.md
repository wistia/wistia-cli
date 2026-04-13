## wistia captions create-multipart

Create Captions

### Synopsis

Adds captions to a specified media by providing an SRT file or its contents directly.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia captions create-multipart [flags]
```

### Examples

```
  wistia captions create-multipart --media-hashed-id <id>
```

### Options

```
  -c, --caption-file string      Either an attached SRT file or a string parameter with the contents of an SRT file. [required]
  -h, --help                     help for create-multipart
  -l, --language string          An optional parameter that denotes which language this file represents. Should conform to ISO-639–2. If left unspecified, the language code will be detected automatically.
  -m, --media-hashed-id string   The hashed ID of the media for which captions are to be added. [required]
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
