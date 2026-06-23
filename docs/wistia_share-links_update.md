## wistia share-links update

Update share link

### Synopsis

Updates the share link for a single media. If the media does not have a
share link yet, one is created with the supplied visibility.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia share-links update [flags]
```

### Examples

```
  wistia share-links update --media-id <id> --visibility unlocked
```

### Options

```
      --body string           Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                  help for update
  -m, --media-id string       The hashed ID of the media. [required]
  -v, --visibility unlocked   Controls who can view the media via this share link.
                              
                              - unlocked: anyone with the link can view the media.
                              - `account`: only signed-in members of the media's account can view.
                              - `locked`: only contacts with access to the media's folder can view.
                              - `domain_verified`: only viewers signed in with an email address at a
                                domain verified on the media's account can view. Requires the account
                                to be enrolled in the domain validation gate; otherwise setting this
                                value returns 400.
                               (options: unlocked, account, locked, domain_verified) [required]
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

* [wistia share-links](wistia_share-links.md)	 - Operations for share-links
