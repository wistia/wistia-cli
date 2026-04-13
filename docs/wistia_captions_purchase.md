## wistia captions purchase

Purchase Captions

### Synopsis

This method is for purchasing English captions for a media. The request will charge the credit card on the account if successful. A saved credit card is required to use this endpoint.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia captions purchase [flags]
```

### Examples

```
  wistia captions purchase --media-hashed-id <id>
```

### Options

```
      --automated                Order computer-generated captions (free) or human-generated captions ($2.50/minute).
      --automatically-enable     Automatically enable captions for the video once the order is ready or hold the captions for review before manually enabling. (default true)
      --body string              Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                     help for purchase
  -m, --media-hashed-id string   Unique identifier for the video. [required]
  -r, --rush                     Enable rush order for one business day turnaround ($4.00/minute) or standard four business day turnaround for human-generated captions ($2.50/minute). Rush can only be used for human-generated captions. (default true)
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
