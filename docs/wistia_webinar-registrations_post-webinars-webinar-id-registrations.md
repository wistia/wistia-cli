## wistia webinar-registrations post-webinars-webinar-id-registrations

Create Webinar Registration

### Synopsis

Register a person for a webinar by providing their email, first name, and last name.

This endpoint generates a unique visitor key and returns a personalized webinar URL for the registrant.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia webinar-registrations post-webinars-webinar-id-registrations [flags]
```

### Examples

```
  wistia webinar-registrations post-webinars-webinar-id-registrations --webinar-id <id> --email john.doe@example.com --first-name John --last-name Doe
```

### Options

```
      --body string         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -e, --email string        Email address of the registrant [required]
  -f, --first-name string   First name of the registrant [required]
  -h, --help                help for post-webinars-webinar-id-registrations
  -l, --last-name string    Last name of the registrant [required]
  -w, --webinar-id string   Hashed ID of the webinar [required]
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

* [wistia webinar-registrations](wistia_webinar-registrations.md)	 - Operations for webinar-registrations
