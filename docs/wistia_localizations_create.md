## wistia localizations create

Localizations Create

### Synopsis

Create a new localization.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia localizations create [flags]
```

### Examples

```
  wistia localizations create --media-hashed-id <id> --output-language <value>
```

### Options

```
  -a, --auto-enable              Whether to automatically enable the localization. (default true)
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                     help for create
  -m, --media-hashed-id string   The hashed ID of the media to create a localization for. [required]
      --output-language string   The language to localize the media to as a 3-character IETF language code. [required]
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

* [wistia localizations](wistia_localizations.md)	 - Operations for localizations
