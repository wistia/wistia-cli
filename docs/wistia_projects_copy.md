## wistia projects copy

Project Copy

### Synopsis

This method does not copy the project’s sharing information (i.e. users that could see the old project will not automatically be able to see the new one).
For the request you can specify the owner of a new project by passing an optional parameter. The person you specify must be a Manager in the account.
For the response, the HTTP Location header will be set to the URL where the new project resource resides. The body of the response will contain an object representing the new copy of the project that was just created.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia projects copy [flags]
```

### Examples

```
  wistia projects copy --id <id>
```

### Options

```
  -a, --admin-email string   The email address of the account Manager that will be the owner of the new project. Defaults to the Account Owner if invalid or omitted.
      --body string          Request body as JSON (alternative to individual flags). Can also be provided via stdin.
  -h, --help                 help for copy
  -i, --id string            Project Hashed ID [required]
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

* [wistia projects](wistia_projects.md)	 - Operations for projects
