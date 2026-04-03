## wistia project-sharings list

Project Sharing List

### Synopsis

A sharing is an object that links either a contact or a contact group to a project, including information about the contacts' permissions to that project.
Retrieve a list of all sharings for a given project. Supports paging and sorting.

## Requires api token with one of the following permissions
```
Read, update & delete anything
Read all data
```

```
wistia project-sharings list [flags]
```

### Examples

```
  wistia project-sharings list --project-id <id>
```

### Options

```
  -h, --help                help for list
      --page int            Page number to retrieve
      --per-page int        Number of projects per page
      --project-id string   Project Hashed ID [required]
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

* [wistia project-sharings](wistia_project-sharings.md)	 - Operations for project-sharings
