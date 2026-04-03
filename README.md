# wistia

Command-line interface for the *Data* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github-com/wistia/wistia-cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)


<br /><br />
> [!IMPORTANT]
> This CLI is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/wistia/wistia). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Data API: Wistia Data API
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [wistia](#wistia)
  * [CLI Installation](#cli-installation)
  * [Shell Completion](#shell-completion)
  * [CLI Example Usage](#cli-example-usage)
  * [Authentication](#authentication)
  * [Available Commands](#available-commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start CLI Installation [installation] -->
## CLI Installation

### Quick Install (Linux/macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/wistia/wistia-cli/main/scripts/install.sh | bash
```

### Quick Install (Windows PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/wistia/wistia-cli/main/scripts/install.ps1 | iex
```

### Go Install

Alternatively, install directly via Go:

```bash
go install github.com/wistia/wistia-cli/cmd/wistia@latest
```

### Manual Download

Download pre-built binaries for your platform from the [releases page](https://github.com/wistia/wistia-cli/releases).
<!-- End CLI Installation [installation] -->

<!-- Start Shell Completion [completion] -->
## Shell Completion

Shell completions are available for Bash, Zsh, Fish, and PowerShell.

### Bash

```bash
# Add to ~/.bashrc:
source <(wistia completion bash)

# Or install permanently:
wistia completion bash > /etc/bash_completion.d/wistia
```

### Zsh

```zsh
# Add to ~/.zshrc:
source <(wistia completion zsh)

# Or install permanently:
wistia completion zsh > "${fpath[1]}/_wistia"
```

### Fish

```fish
wistia completion fish | source

# Or install permanently:
wistia completion fish > ~/.config/fish/completions/wistia.fish
```

### PowerShell

```powershell
wistia completion powershell | Out-String | Invoke-Expression
```
<!-- End Shell Completion [completion] -->

<!-- Start CLI Example Usage [usage] -->
## CLI Example Usage

### Example

```bash
wistia media upload-form --bearer-auth 'Bearer test_token' --url 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4' --low-priority true

```
<!-- End CLI Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
wistia --bearer-auth <value> <command> [arguments]
```

### 2. Environment variables

Set credentials via environment variables:

| Variable | Description |
|----------|-------------|
| `WISTIA_CLI_BEARER_AUTH` | HTTP Bearer |

### 3. OS Keychain (recommended for workstations)

Credentials are stored securely in your operating system's keychain when you run:

```bash
wistia configure
```

Secret credentials (tokens, API keys, passwords) are automatically stored in:
- **macOS**: Keychain
- **Linux**: GNOME Keyring / KWallet (via D-Bus Secret Service)
- **Windows**: Windows Credential Locker

If no keychain is available (e.g., in CI environments), credentials fall back to the config file.

### 4. Configuration file

Run the interactive `configure` command to store non-secret settings:

```bash
wistia configure
```

Configuration is stored in `~/.config/wistia/config.yaml`.
<!-- End Authentication [security] -->

<!-- Start Available Commands [operations] -->
## Available Commands

<details open>
<summary>Available commands</summary>

### [media](docs/wistia_media.md)

* [`upload-form`](docs/wistia_media_upload-form.md) - Upload or Import Media
* [`upload-multipart`](docs/wistia_media_upload-multipart.md) - Upload or Import Media
* [`list`](docs/wistia_media_list.md) - Media List
* [`get`](docs/wistia_media_get.md) - Media Show
* [`update`](docs/wistia_media_update.md) - Media Update
* [`delete`](docs/wistia_media_delete.md) - Media Delete
* [`copy`](docs/wistia_media_copy.md) - Media Copy
* [`swap`](docs/wistia_media_swap.md) - Media Swap
* [`get-stats`](docs/wistia_media_get-stats.md) - Media Stats
* [`translate`](docs/wistia_media_translate.md) - Media Translate
* [`archive`](docs/wistia_media_archive.md) - Medias Archive
* [`move`](docs/wistia_media_move.md) - Media Move
* [`restore`](docs/wistia_media_restore.md) - Media Restore

### [projects](docs/wistia_projects.md)

* [`list`](docs/wistia_projects_list.md) - Project List
* [`create`](docs/wistia_projects_create.md) - Project Create
* [`get`](docs/wistia_projects_get.md) - Project Show
* [`update`](docs/wistia_projects_update.md) - Project Update
* [`delete`](docs/wistia_projects_delete.md) - Project Delete
* [`copy`](docs/wistia_projects_copy.md) - Project Copy

### [subfolders](docs/wistia_subfolders.md)

* [`list`](docs/wistia_subfolders_list.md) - Subfolder List
* [`create`](docs/wistia_subfolders_create.md) - Create Subfolder
* [`get`](docs/wistia_subfolders_get.md) - Show Subfolder
* [`update`](docs/wistia_subfolders_update.md) - Update Subfolder
* [`delete-subfolder`](docs/wistia_subfolders_delete-subfolder.md) - Delete Subfolder

### [project-sharings](docs/wistia_project-sharings.md)

* [`list`](docs/wistia_project-sharings_list.md) - Project Sharing List
* [`create`](docs/wistia_project-sharings_create.md) - Project Sharing Create
* [`get`](docs/wistia_project-sharings_get.md) - Project Sharing Show
* [`update`](docs/wistia_project-sharings_update.md) - Project Sharing Update
* [`delete`](docs/wistia_project-sharings_delete.md) - Project Sharing Delete

### [account](docs/wistia_account.md)

* [`get`](docs/wistia_account_get.md) - Account Show

### [allowed-domains](docs/wistia_allowed-domains.md)

* [`list`](docs/wistia_allowed-domains_list.md) - Allowed Domains List
* [`create`](docs/wistia_allowed-domains_create.md) - Allowed Domain Create
* [`get`](docs/wistia_allowed-domains_get.md) - Allowed Domain Show
* [`delete`](docs/wistia_allowed-domains_delete.md) - Allowed Domain Delete

### [background-job-status](docs/wistia_background-job-status.md)

* [`get`](docs/wistia_background-job-status_get.md) - Background Job Status Show

### [customizations](docs/wistia_customizations.md)

* [`get`](docs/wistia_customizations_get.md) - Customizations Show
* [`create`](docs/wistia_customizations_create.md) - Customizations Create
* [`update`](docs/wistia_customizations_update.md) - Customizations Update
* [`delete`](docs/wistia_customizations_delete.md) - Customizations Delete

### [captions](docs/wistia_captions.md)

* [`list`](docs/wistia_captions_list.md) - Captions List
* [`create`](docs/wistia_captions_create.md) - Captions Create
* [`create-multipart`](docs/wistia_captions_create-multipart.md) - Captions Create
* [`purchase`](docs/wistia_captions_purchase.md) - Captions Purchase
* [`get`](docs/wistia_captions_get.md) - Captions Show
* [`update`](docs/wistia_captions_update.md) - Captions Update
* [`update-multipart`](docs/wistia_captions_update-multipart.md) - Captions Update
* [`delete`](docs/wistia_captions_delete.md) - Captions Delete

### [trims](docs/wistia_trims.md)

* [`create`](docs/wistia_trims_create.md) - Trims Create

### [localizations](docs/wistia_localizations.md)

* [`list`](docs/wistia_localizations_list.md) - Localizations List
* [`create`](docs/wistia_localizations_create.md) - Localizations Create
* [`get`](docs/wistia_localizations_get.md) - Localizations Show
* [`delete`](docs/wistia_localizations_delete.md) - Localizations Delete

### [tags](docs/wistia_tags.md)

* [`list`](docs/wistia_tags_list.md) - Tags List
* [`create`](docs/wistia_tags_create.md) - Tags Create
* [`delete`](docs/wistia_tags_delete.md) - Tags Delete

### [search](docs/wistia_search.md)

* [`search`](docs/wistia_search_search.md) - Search

### [channels](docs/wistia_channels.md)

* [`list`](docs/wistia_channels_list.md) - Channels List
* [`get`](docs/wistia_channels_get.md) - Channels Show

#### [channels-channel-episodes](docs/wistia_channels_channels-channel-episodes.md)

* [`list`](docs/wistia_channels_channels-channel-episodes_list.md) - Channel Episodes List filtered by channel

### [channel-episodes](docs/wistia_channel-episodes.md)

* [`get`](docs/wistia_channel-episodes_get.md) - Channel Episodes Show
* [`create`](docs/wistia_channel-episodes_create.md) - Channel Episode Create
* [`list`](docs/wistia_channel-episodes_list.md) - Channel Episodes List

### [expiring-access-tokens](docs/wistia_expiring-access-tokens.md)

* [`create`](docs/wistia_expiring-access-tokens_create.md) - Create an expiring access token

### [live-stream-events](docs/wistia_live-stream-events.md)

* [`list`](docs/wistia_live-stream-events_list.md) - Live Stream Events List
* [`create`](docs/wistia_live-stream-events_create.md) - Live Stream Event Create
* [`get`](docs/wistia_live-stream-events_get.md) - Live Stream Event Show
* [`update`](docs/wistia_live-stream-events_update.md) - Live Stream Event Update
* [`delete`](docs/wistia_live-stream-events_delete.md) - Live Stream Event Delete

### [stats-account](docs/wistia_stats-account.md)

* [`get`](docs/wistia_stats-account_get.md) - Stats:Account Show

### [stats-projects](docs/wistia_stats-projects.md)

* [`get`](docs/wistia_stats-projects_get.md) - Stats:Projects Show

### [stats-media](docs/wistia_stats-media.md)

* [`get`](docs/wistia_stats-media_get.md) - Stats:Media Show
* [`get-by-date`](docs/wistia_stats-media_get-by-date.md) - Stats:Media By Date
* [`get-engagement`](docs/wistia_stats-media_get-engagement.md) - Stats:Media Engagement

### [stats-visitors](docs/wistia_stats-visitors.md)

* [`list`](docs/wistia_stats-visitors_list.md) - Stats:Visitors List
* [`get`](docs/wistia_stats-visitors_get.md) - Stats:Visitors Show

### [stats-events](docs/wistia_stats-events.md)

* [`list`](docs/wistia_stats-events_list.md) - Stats:Events List
* [`get`](docs/wistia_stats-events_get.md) - Stats:Events Show

</details>
<!-- End Available Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Operations that accept a request body support three input methods, with a clear priority chain:

### Individual flags (highest priority)

```bash
wistia <command> --name "Jane" --age 30
```

### `--body` flag

Provide the entire request body as a JSON string:

```bash
wistia <command> --body '{"name": "John", "age": 30}'
```

Individual flags override `--body` values:

```bash
# Result: {name: "Jane", age: 30}
wistia <command> --body '{"name": "John", "age": 30}' --name "Jane"
```

### Stdin piping (lowest priority)

Pipe JSON into any command that accepts a request body:

```bash
echo '{"name": "John", "age": 30}' | wistia <command>
```

Individual flags override stdin values:

```bash
# Result: {name: "Jane", age: 30}
echo '{"name": "John", "age": 30}' | wistia <command> --name "Jane"
```

This is useful for chaining commands, reading from files, or scripting:

```bash
# Read body from a file
wistia <command> < request.json

# Pipe from another command
curl -s https://example.com/data.json | wistia <command>
```

### Priority

When multiple input methods are used, the priority is:

| Priority | Source | Description |
|----------|--------|-------------|
| 1 (highest) | Individual flags | `--name "Jane"` always wins |
| 2 | `--body` flag | Whole-body JSON via flag |
| 3 (lowest) | Stdin | Piped JSON input |
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL

Use `--server-url` to override the server URL entirely, bypassing any named or indexed server selection:

```bash
wistia --server-url https://custom-api.example.com <command> [arguments]
```

**Precedence**: `--server-url` > `--server` > default
<!-- End Server Selection [server] -->

<!-- Start Output Formats [output-formats] -->
## Output Formats

Every command supports a `--output-format` flag that controls how the response is rendered to stdout.

### Available formats

| Format | Flag | Description |
|--------|------|-------------|
| Pretty | `--output-format pretty` (default) | Aligned key-value pairs with color, nested indentation. Human-readable at a glance. |
| JSON | `--output-format json` | JSON output. Passthrough when the response is already JSON (preserves original field order and numeric precision). Falls back to typed marshaling otherwise. |
| YAML | `--output-format yaml` | YAML output via standard marshaling. |
| Table | `--output-format table` | Tabular output for array responses. |
| TOON | `--output-format toon` | [Token-Oriented Object Notation](https://github.com/toon-format/spec) — a compact, line-oriented format that typically uses 30–60% fewer tokens than JSON. Well-suited for piping responses into LLM prompts. |

```bash
# Default pretty output
wistia <command>

# Machine-readable JSON
wistia <command> --output-format json

# TOON for LLM-friendly compact output
wistia <command> --output-format toon

# Pipe JSON to jq without using --output-format
wistia <command> --output-format json | jq '.fieldName'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
wistia <command> --jq '.name'

# Filter an array
wistia <command> --jq '.items[] | select(.active == true)'
```

### Color control

Use `--color` to control terminal colors:

| Value | Behavior |
|-------|----------|
| `auto` (default) | Color when stdout is a TTY, plain text otherwise |
| `always` | Always colorize |
| `never` | Never colorize |

The `NO_COLOR` and `FORCE_COLOR` environment variables are also respected.

### Streaming and pagination

When using `--all` (pagination) or streaming operations, output is written incrementally as items arrive:

| Format | Streaming behavior |
|--------|-------------------|
| `json` | One compact JSON object per line ([NDJSON](https://github.com/ndjson/ndjson-spec)) |
| `yaml` | YAML documents separated by `---` |
| `toon` | One TOON-encoded object per block, separated by blank lines |
| `pretty` (default) | Pretty-printed items separated by blank lines |
<!-- End Output Formats [output-formats] -->

<!-- Start Error Handling [errors] -->
## Error Handling

The CLI uses standard exit codes to indicate success or failure:

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success |
| `1` | Error (API error, invalid input, etc.) |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
wistia ... > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
wistia <command> --dry-run
```

Output goes to stderr and includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

The command exits successfully without contacting the API. This is useful for verifying request construction before executing.

### Debug

Log request and response diagnostics while running normally:

```bash
wistia <command> --debug
```

Debug output goes to stderr and includes:
- Request method, URL, headers, and body preview
- Response status, headers, and body preview
- Transport errors (if any)

The command still executes normally and produces its regular output on stdout.

### Flag Precedence

If both `--dry-run` and `--debug` are set, `--dry-run` takes precedence and no network calls are made.

### Security

Sensitive information is automatically redacted in diagnostic output:
- **Headers**: `Authorization`, `Cookie`, `Set-Cookie`, `X-API-Key`, and other security headers show `[REDACTED]`
- **Body**: JSON fields named `password`, `secret`, `token`, `api_key`, `client_secret`, etc. show `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This CLI is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/wistia/wistia-cli&utm_campaign=cli)
