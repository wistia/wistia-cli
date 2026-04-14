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
### Homebrew (macOS/Linux)

```bash
brew install wistia/wistia-cli/wistia
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
wistia upload-or-import-media post-form --bearer-auth 'Bearer test_token' --url 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4' --low-priority true

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

### [upload-or-import-media](docs/wistia_upload-or-import-media.md)

* [`post-form`](docs/wistia_upload-or-import-media_post-form.md) - Upload or Import Media
* [`post-multipart`](docs/wistia_upload-or-import-media_post-multipart.md) - Upload or Import Media

### [media-extended-audio-descriptions](docs/wistia_media-extended-audio-descriptions.md)

* [`get`](docs/wistia_media-extended-audio-descriptions_get.md) - List Media Extended Audio Descriptions
* [`get-media-extended-audio-descriptions-id`](docs/wistia_media-extended-audio-descriptions_get-media-extended-audio-descriptions-id.md) - Show Media Extended Audio Description
* [`delete-media-extended-audio-descriptions-id`](docs/wistia_media-extended-audio-descriptions_delete-media-extended-audio-descriptions-id.md) - Delete Media Extended Audio Description
* [`post-media-extended-audio-descriptions-order`](docs/wistia_media-extended-audio-descriptions_post-media-extended-audio-descriptions-order.md) - Order Extended Audio Description
* [`get-media-extended-audio-descriptions-order-status-id`](docs/wistia_media-extended-audio-descriptions_get-media-extended-audio-descriptions-order-status-id.md) - Get Order Status

### [folders](docs/wistia_folders.md)

* [`list`](docs/wistia_folders_list.md) - List Folders
* [`create`](docs/wistia_folders_create.md) - Create Folder
* [`get`](docs/wistia_folders_get.md) - Show Folder
* [`update`](docs/wistia_folders_update.md) - Update Folder
* [`delete`](docs/wistia_folders_delete.md) - Delete Folder
* [`copy`](docs/wistia_folders_copy.md) - Copy Folder

### [subfolders](docs/wistia_subfolders.md)

* [`list`](docs/wistia_subfolders_list.md) - List Subfolders
* [`create`](docs/wistia_subfolders_create.md) - Create Subfolder
* [`get`](docs/wistia_subfolders_get.md) - Show Subfolder
* [`update`](docs/wistia_subfolders_update.md) - Update Subfolder
* [`delete`](docs/wistia_subfolders_delete.md) - Delete Subfolder
* [`bulk-delete`](docs/wistia_subfolders_bulk-delete.md) - Bulk Delete Subfolders

### [folder-sharings](docs/wistia_folder-sharings.md)

* [`list`](docs/wistia_folder-sharings_list.md) - List Folder Sharings
* [`create`](docs/wistia_folder-sharings_create.md) - Create Folder Sharing
* [`get`](docs/wistia_folder-sharings_get.md) - Show Folder Sharing
* [`update`](docs/wistia_folder-sharings_update.md) - Update Folder Sharing
* [`delete`](docs/wistia_folder-sharings_delete.md) - Delete Folder Sharing

### [media](docs/wistia_media.md)

* [`list`](docs/wistia_media_list.md) - List Media
* [`get`](docs/wistia_media_get.md) - Show Media
* [`update`](docs/wistia_media_update.md) - Update Media
* [`delete`](docs/wistia_media_delete.md) - Delete Media
* [`copy`](docs/wistia_media_copy.md) - Copy Media
* [`swap`](docs/wistia_media_swap.md) - Swap Media
* [`get-stats`](docs/wistia_media_get-stats.md) - Show Media Aggregated Stats
* [`translate`](docs/wistia_media_translate.md) - Translate Media
* [`import-url`](docs/wistia_media_import-url.md) - Import Media from URL
* [`archive`](docs/wistia_media_archive.md) - Archive Media
* [`move`](docs/wistia_media_move.md) - Move Media
* [`restore`](docs/wistia_media_restore.md) - Restore Media
* [`bulk-copy`](docs/wistia_media_bulk-copy.md) - Bulk Copy Media

### [taggings](docs/wistia_taggings.md)

* [`bulk-create`](docs/wistia_taggings_bulk-create.md) - Bulk Tag Media

### [account](docs/wistia_account.md)

* [`get`](docs/wistia_account_get.md) - Get Current Account
* [`get-token-details`](docs/wistia_account_get-token-details.md) - Get Current Token

### [allowed-domains](docs/wistia_allowed-domains.md)

* [`list`](docs/wistia_allowed-domains_list.md) - List Allowed Domains
* [`create`](docs/wistia_allowed-domains_create.md) - Create Allowed Domain
* [`get`](docs/wistia_allowed-domains_get.md) - Show Allowed Domain
* [`delete`](docs/wistia_allowed-domains_delete.md) - Delete Allowed Domain

### [background-job-status](docs/wistia_background-job-status.md)

* [`get`](docs/wistia_background-job-status_get.md) - Show Background Job Status

### [customizations](docs/wistia_customizations.md)

* [`get`](docs/wistia_customizations_get.md) - Show Customizations
* [`create`](docs/wistia_customizations_create.md) - Create Customizations
* [`update`](docs/wistia_customizations_update.md) - Update Customizations
* [`delete`](docs/wistia_customizations_delete.md) - Delete Customizations

### [captions](docs/wistia_captions.md)

* [`list`](docs/wistia_captions_list.md) - List Captions by Media
* [`create`](docs/wistia_captions_create.md) - Create Captions
* [`create-multipart`](docs/wistia_captions_create-multipart.md) - Create Captions
* [`list-all`](docs/wistia_captions_list-all.md) - List Captions
* [`purchase`](docs/wistia_captions_purchase.md) - Purchase Captions
* [`get`](docs/wistia_captions_get.md) - Show Captions
* [`update`](docs/wistia_captions_update.md) - Update Captions
* [`update-multipart`](docs/wistia_captions_update-multipart.md) - Update Captions
* [`delete`](docs/wistia_captions_delete.md) - Delete Captions

### [trims](docs/wistia_trims.md)

* [`create`](docs/wistia_trims_create.md) - Create Media from Trims

### [localizations](docs/wistia_localizations.md)

* [`list`](docs/wistia_localizations_list.md) - List Localizations
* [`create`](docs/wistia_localizations_create.md) - Create Localization
* [`get`](docs/wistia_localizations_get.md) - Show Localization
* [`delete`](docs/wistia_localizations_delete.md) - Delete Localization

### [tags](docs/wistia_tags.md)

* [`list`](docs/wistia_tags_list.md) - List Tags
* [`create`](docs/wistia_tags_create.md) - Create Tags
* [`delete`](docs/wistia_tags_delete.md) - Delete Tag

### [search](docs/wistia_search.md)

* [`search`](docs/wistia_search_search.md) - Search

### [channels](docs/wistia_channels.md)

* [`list`](docs/wistia_channels_list.md) - List Channels
* [`create`](docs/wistia_channels_create.md) - Create Channel
* [`get`](docs/wistia_channels_get.md) - Show Channel
* [`update`](docs/wistia_channels_update.md) - Update Channel
* [`delete`](docs/wistia_channels_delete.md) - Delete Channel

#### [channels-channel-episodes](docs/wistia_channels_channels-channel-episodes.md)

* [`list`](docs/wistia_channels_channels-channel-episodes_list.md) - List Channel Episodes by Channel

### [channel-episodes](docs/wistia_channel-episodes.md)

* [`get`](docs/wistia_channel-episodes_get.md) - Show Channel Episode
* [`create`](docs/wistia_channel-episodes_create.md) - Create Channel Episode
* [`list`](docs/wistia_channel-episodes_list.md) - List Channel Episodes
* [`update`](docs/wistia_channel-episodes_update.md) - Update Channel Episode
* [`delete`](docs/wistia_channel-episodes_delete.md) - Delete Channel Episode
* [`publish`](docs/wistia_channel-episodes_publish.md) - Publish Channel Episode
* [`unpublish`](docs/wistia_channel-episodes_unpublish.md) - Un-publish Channel Episode

### [expiring-access-tokens](docs/wistia_expiring-access-tokens.md)

* [`create`](docs/wistia_expiring-access-tokens_create.md) - Create Expiring Access Token

### [webinars](docs/wistia_webinars.md)

* [`list`](docs/wistia_webinars_list.md) - List Webinars
* [`create`](docs/wistia_webinars_create.md) - Create Webinar
* [`get`](docs/wistia_webinars_get.md) - Show Webinar
* [`update`](docs/wistia_webinars_update.md) - Update Webinar
* [`delete`](docs/wistia_webinars_delete.md) - Delete Webinar

### [webinar-registrations](docs/wistia_webinar-registrations.md)

* [`get-webinars-webinar-id-registrations`](docs/wistia_webinar-registrations_get-webinars-webinar-id-registrations.md) - List Webinar Registrations
* [`create`](docs/wistia_webinar-registrations_create.md) - Create Webinar Registration

### [remix](docs/wistia_remix.md)

* [`post-remixes`](docs/wistia_remix_post-remixes.md) - Create Remix
* [`get-remixes-remix-hashed-id`](docs/wistia_remix_get-remixes-remix-hashed-id.md) - Get Remix
* [`post-remixes-remix-hashed-id-continue`](docs/wistia_remix_post-remixes-remix-hashed-id-continue.md) - Continue Remix
* [`post-remixes-remix-hashed-id-export`](docs/wistia_remix_post-remixes-remix-hashed-id-export.md) - Export Remix
* [`get-remix-account-status`](docs/wistia_remix_get-remix-account-status.md) - Get Remix Account Status

### [stats-account](docs/wistia_stats-account.md)

* [`get`](docs/wistia_stats-account_get.md) - Show Current Account Stats
* [`get-stats-account-by-date`](docs/wistia_stats-account_get-stats-account-by-date.md) - Show Account Stats by Date

### [stats-projects](docs/wistia_stats-projects.md)

* [`get`](docs/wistia_stats-projects_get.md) - Show Project Stats

### [stats-media](docs/wistia_stats-media.md)

* [`get`](docs/wistia_stats-media_get.md) - Show Media Stats
* [`get-by-date`](docs/wistia_stats-media_get-by-date.md) - Show Media Stats by Date
* [`get-engagement`](docs/wistia_stats-media_get-engagement.md) - Show Media Engagement

### [stats-visitors](docs/wistia_stats-visitors.md)

* [`list`](docs/wistia_stats-visitors_list.md) - List Visitors
* [`get`](docs/wistia_stats-visitors_get.md) - Show Visitor

### [stats-events](docs/wistia_stats-events.md)

* [`list`](docs/wistia_stats-events_list.md) - List Events
* [`get`](docs/wistia_stats-events_get.md) - Show Event

### [analytics-media](docs/wistia_analytics-media.md)

* [`get`](docs/wistia_analytics-media_get.md) - Show Media Analytics
* [`get-timeseries`](docs/wistia_analytics-media_get-timeseries.md) - Show Media Analytics Timeseries
* [`get-embed-locations`](docs/wistia_analytics-media_get-embed-locations.md) - Show Media Embed Locations
* [`get-traffic`](docs/wistia_analytics-media_get-traffic.md) - Show Media Traffic Breakdown
* [`get-conversions`](docs/wistia_analytics-media_get-conversions.md) - Show Media Form Conversions
* [`get-languages`](docs/wistia_analytics-media_get-languages.md) - Show Media Languages

### [analytics-webinar](docs/wistia_analytics-webinar.md)

* [`get`](docs/wistia_analytics-webinar_get.md) - Show Webinar Analytics
* [`get-registration`](docs/wistia_analytics-webinar_get-registration.md) - Show Webinar Registration Timeseries
* [`get-traffic`](docs/wistia_analytics-webinar_get-traffic.md) - Show Webinar Traffic Breakdown
* [`get-audience`](docs/wistia_analytics-webinar_get-audience.md) - Show Webinar Audience
* [`get-histograms`](docs/wistia_analytics-webinar_get-histograms.md) - Show Webinar Histograms

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
