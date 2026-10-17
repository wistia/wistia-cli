# wistia

Command-line interface for the *Data* API.

[![Built by Speakeasy](https://img.shields.io/badge/Built_by-SPEAKEASY-374151?style=for-the-badge&labelColor=f3f4f6)](https://www.speakeasy.com/?utm_source=github-com/wistia/wistia-cli&utm_campaign=cli)
[![License: MIT](https://img.shields.io/badge/LICENSE_//_MIT-3b5bdb?style=for-the-badge&labelColor=eff6ff)](https://opensource.org/licenses/MIT)

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
  * [For AI agents](#for-ai-agents)
  * [Authentication](#authentication)
  * [Commands](#commands)
  * [Request Body Input](#request-body-input)
  * [Server Selection](#server-selection)
  * [Output Formats](#output-formats)
  * [Error Handling](#error-handling)
  * [Diagnostics](#diagnostics)
* [Development](#development)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

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
brew install wistia/tap/wistia-cli
```

### Manual Download

Download pre-built binaries for your platform from the [releases page](https://github.com/wistia/wistia-cli/releases).
<!-- No CLI Installation [installation] -->

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
wistia upload-or-import-media post-form --bearer-auth 'Bearer test_token' --url 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4' --low-priority=true

```
<!-- End CLI Example Usage [usage] -->

<!-- Start For AI agents [agents] -->
## For AI agents

This CLI is built to be driven by AI coding agents as well as people: everything an agent needs is discoverable from the binary itself, and every command can be validated without credentials. Work down this ladder:

| Run | You get |
|-----|---------|
| `wistia --help`, `wistia account get --help` | Commands by category, runnable examples, flags |
| `wistia --usage`, `wistia account get --usage` | The command surface as machine-readable [KDL](https://kdl.dev): commands, aliases, flags, defaults, env vars, config keys |
| `wistia account update-brand-preload --schema` | The exact JSON Schema of the command's request body (all `$ref`s bundled) — build a valid `--body` from it |
| `wistia account get --dry-run` | The exact HTTP request (method, URL, headers, body), with no credentials or network call |
| `wistia account get --output-format json` (or `--jq`) | Machine-readable output |

### Discover the command surface

```bash
# Every command, flag, default, env var and config key, as KDL
wistia --usage

# One command's subtree only
wistia account get --usage
```

### Read the exact request schema

`--schema` is available on every command that accepts a request body (`--body`, stdin, or a whole-body flag where the command has one), including intent commands. It prints the JSON Schema the request is validated against and exits without calling the API.

```bash
# JSON Schema (draft 2020-12) of the request body, with every $ref bundled under $defs
wistia account update-brand-preload --schema
```

### Probe before you spend

Start quota-spending commands with `--dry-run`. It validates inputs, resolves the request, redacts secrets and binary payloads, makes no network call, and exits 0. It never reads the OS keychain; credentials supplied by flag, environment, or config file are included only as `[REDACTED]`.

```bash
# Human preview: the [DRY-RUN] block is on stderr and stdout is empty
wistia account get --dry-run

# Machine preview: compact JSON on stdout and silent stderr
wistia account get --dry-run --output-format json
```

The machine form writes one object per would-be request, one per line (NDJSON for multi-request commands), with exactly this shape:

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

`body` is a parsed JSON value when the body is JSON, a string for text, `"<bytes:N>"` for binary data, and `null` when absent. An explicit caller `--jq` also selects this JSON preview protocol, but the filter is not applied to preview objects. Command-declared jq presets do not select or filter the preview.

Local mutation commands make no request under `--dry-run`: instead of a preview they emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object. `select(.request)` keeps only would-be requests; `select(.local)` keeps the local no-ops.

### Machine-readable output

```bash
# JSON on stdout
wistia account get --output-format json

# Filter or reshape with a jq expression (always emits JSON, overrides --output-format)
wistia account get --jq '.'

# Print jq string results as plain text instead of JSON strings (like jq -r)
wistia account get --jq '.' --raw-output
```

`--output-format toon` emits [TOON](https://github.com/toon-format/spec), a compact line-oriented format that uses fewer tokens than JSON; it is the default in agent mode.

### Interactive mode
Required-input prompts and guided `configure` / `auth login` forms are enabled by default. Required-input prompts require an interactive terminal; off-TTY forms read line input from stdin. Use `--no-interactive` to force flag-only execution.

```bash
# Prompt for missing command inputs
wistia account get --interactive

# Open the guided configuration form
wistia configure --interactive

# Explicitly launch the terminal command explorer
wistia explore
```

### Agent mode and structured errors
Agent mode turns on automatically when a known agent environment is detected (`CLAUDECODE`, `CURSOR_AGENT`, `CODEX`, `AIDER`, `CLINE`, `WINDSURF_AGENT`, `GITHUB_COPILOT`, `AMAZON_Q`, `GEMINI_CODE_ASSIST`, `SRC_CODY`) or with `--agent-mode` (`--agent-mode=false` disables detection).
In agent mode interactive prompts never launch, output defaults to TOON, and every failure — API errors and CLI usage errors alike — is one JSON envelope on stderr:
Outside agent mode, explicit JSON and `--jq` preserve the compatibility envelope without classification; enable agent mode to request the classified contract.

```json
{
  "error": "...",
  "error_type": "validation_error",
  "error_reason": "CLI_VALIDATION",
  "exit_code": 2,
  "message": "human-readable message",
  "hints": ["what to try next"]
}
```

`error_type` is one of `authentication_error`, `authorization_error`, `not_found`, `validation_error`, `rate_limit_error`, `server_error`, `api_error`, `connection_error`, `protocol_error`, `runtime_error`, `unsupported_error`, `async_failed`, `async_timeout`, `async_unknown_state`. Classification derives from the HTTP status and transport evidence; `error_reason` is absent for API errors. Status-less local failures may use `CLI_VALIDATION`, `CLI_CONNECTION`, `CLI_PROTOCOL`, `CLI_RUNTIME`, `CLI_UNAVAILABLE`, `CLI_AUTHENTICATION`, or the async polling reasons `CLI_ASYNC_FAILED`, `CLI_ASYNC_TIMEOUT`, and `CLI_ASYNC_UNKNOWN_STATE`. `hints` preserves server guidance first, adds the most specific local taxonomy guidance, then typed CLI and command-specific guidance, removing exact duplicates. `exit_code` is always the code for the final `error_type` shown in the envelope: 1 runtime, 2 usage, or 3 authentication/authorization.
<!-- End For AI agents [agents] -->

<!-- Start Authentication [security] -->
## Authentication

Authentication credentials can be configured in four ways (in order of priority):

### 1. Command-line flags

Pass credentials directly as flags to any command:

```bash
wistia --bearer-auth "$WISTIA_CLI_BEARER_AUTH" account get
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

Where the keychain cannot be unlocked (e.g., headless Linux or SSH sessions with a locked GNOME Keyring), skip it entirely with `--no-keyring`, `WISTIA_CLI_NO_KEYRING=true`, or `no_keyring: true` in the config file. The keychain is then never read or written, and secrets are stored in the config file instead; flags and environment variables still take precedence.

### 4. Configuration file

Run the interactive `configure` command to store non-secret settings:

```bash
wistia configure
```

Configuration is stored in `~/.config/wistia/config.yaml`.
<!-- End Authentication [security] -->

<!-- Start Commands [operations] -->
## Commands

<details open>
<summary>Available commands</summary>

* [`upload-or-import-media`](docs/wistia_upload-or-import-media.md) - Operations for upload-or-import-media
  * [`post-form`](docs/wistia_upload-or-import-media_post-form.md) - Upload or Import Media
  * [`post-multipart`](docs/wistia_upload-or-import-media_post-multipart.md) - Upload or Import Media
* [`push-devices`](docs/wistia_push-devices.md) - Operations for push-devices
  * [`create`](docs/wistia_push-devices_create.md) - Register Push Device
  * [`delete`](docs/wistia_push-devices_delete.md) - Unregister Push Device
* [`review-bundles`](docs/wistia_review-bundles.md) - Operations for review-bundles
  * [`list`](docs/wistia_review-bundles_list.md) - List Review Bundles
  * [`create`](docs/wistia_review-bundles_create.md) - Create Review Bundle
  * [`delete`](docs/wistia_review-bundles_delete.md) - Delete Review Bundle
* [`custom-metadata-field-definitions`](docs/wistia_custom-metadata-field-definitions.md) - Operations for custom-metadata-field-definitions
  * [`get`](docs/wistia_custom-metadata-field-definitions_get.md) - List Custom Metadata Field Definitions
  * [`post`](docs/wistia_custom-metadata-field-definitions_post.md) - Create Custom Metadata Field Definition
  * [`get-custom-metadata-field-definitions-key`](docs/wistia_custom-metadata-field-definitions_get-custom-metadata-field-definitions-key.md) - Show Custom Metadata Field Definition
  * [`put-custom-metadata-field-definitions-key`](docs/wistia_custom-metadata-field-definitions_put-custom-metadata-field-definitions-key.md) - Update Custom Metadata Field Definition
  * [`delete-custom-metadata-field-definitions-key`](docs/wistia_custom-metadata-field-definitions_delete-custom-metadata-field-definitions-key.md) - Archive Custom Metadata Field Definition
  * [`post-custom-metadata-field-definitions-key-restore`](docs/wistia_custom-metadata-field-definitions_post-custom-metadata-field-definitions-key-restore.md) - Restore Custom Metadata Field Definition
* [`deleted-media`](docs/wistia_deleted-media.md) - Operations for deleted-media
  * [`list`](docs/wistia_deleted-media_list.md) - List Deleted Media
  * [`restore`](docs/wistia_deleted-media_restore.md) - Restore Deleted Media
* [`media`](docs/wistia_media.md) - Operations for media
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
* [`customizations`](docs/wistia_customizations.md) - Operations for customizations
  * [`get`](docs/wistia_customizations_get.md) - Show Customizations
  * [`create`](docs/wistia_customizations_create.md) - Create Customizations
  * [`update`](docs/wistia_customizations_update.md) - Update Customizations
  * [`delete`](docs/wistia_customizations_delete.md) - Delete Customizations
  * [`get-appearance`](docs/wistia_customizations_get-appearance.md) - Show Appearance Customizations
  * [`update-appearance`](docs/wistia_customizations_update-appearance.md) - Update Appearance Customizations
  * [`get-playback`](docs/wistia_customizations_get-playback.md) - Show Playback Customizations
  * [`update-playback`](docs/wistia_customizations_update-playback.md) - Update Playback Customizations
  * [`get-thumbnail`](docs/wistia_customizations_get-thumbnail.md) - Show Thumbnail Customizations
  * [`update-thumbnail`](docs/wistia_customizations_update-thumbnail.md) - Update Thumbnail Customizations
  * [`get-accessibility`](docs/wistia_customizations_get-accessibility.md) - Show Accessibility Customizations
  * [`update-accessibility`](docs/wistia_customizations_update-accessibility.md) - Update Accessibility Customizations
  * [`get-chapters`](docs/wistia_customizations_get-chapters.md) - Show Chapters Customizations
  * [`update-chapters`](docs/wistia_customizations_update-chapters.md) - Update Chapters Customizations
  * [`get-engagement`](docs/wistia_customizations_get-engagement.md) - Show Engagement Customizations
  * [`update-engagement`](docs/wistia_customizations_update-engagement.md) - Update Engagement Customizations
  * [`get-related-media`](docs/wistia_customizations_get-related-media.md) - Show Related Media Customizations
  * [`update-related-media`](docs/wistia_customizations_update-related-media.md) - Update Related Media Customizations
  * [`get-sharing`](docs/wistia_customizations_get-sharing.md) - Show Sharing Customizations
  * [`update-sharing`](docs/wistia_customizations_update-sharing.md) - Update Sharing Customizations
  * [`get-lead-capture`](docs/wistia_customizations_get-lead-capture.md) - Show Lead Capture Customizations
  * [`update-lead-capture`](docs/wistia_customizations_update-lead-capture.md) - Update Lead Capture Customizations
  * [`get-access`](docs/wistia_customizations_get-access.md) - Show Access Customizations
  * [`update-access`](docs/wistia_customizations_update-access.md) - Update Access Customizations
* [`share-links`](docs/wistia_share-links.md) - Operations for share-links
  * [`resolve`](docs/wistia_share-links_resolve.md) - Resolve share link
  * [`get`](docs/wistia_share-links_get.md) - Show share link
  * [`update`](docs/wistia_share-links_update.md) - Update share link
  * [`delete`](docs/wistia_share-links_delete.md) - Delete share link
* [`captions`](docs/wistia_captions.md) - Operations for captions
  * [`list`](docs/wistia_captions_list.md) - List Captions by Media
  * [`create`](docs/wistia_captions_create.md) - Create Captions
  * [`create-multipart`](docs/wistia_captions_create-multipart.md) - Create Captions
  * [`list-all`](docs/wistia_captions_list-all.md) - List Captions
  * [`find-matches`](docs/wistia_captions_find-matches.md) - Find Caption Matches
  * [`purchase`](docs/wistia_captions_purchase.md) - Purchase Captions
  * [`get`](docs/wistia_captions_get.md) - Show Captions
  * [`update`](docs/wistia_captions_update.md) - Update Captions
  * [`update-multipart`](docs/wistia_captions_update-multipart.md) - Update Captions
  * [`delete`](docs/wistia_captions_delete.md) - Delete Captions
  * [`edit`](docs/wistia_captions_edit.md) - Edit Captions Text
* [`speakers`](docs/wistia_speakers.md) - Operations for speakers
  * [`assign`](docs/wistia_speakers_assign.md) - Assign Speaker to Media
  * [`remove`](docs/wistia_speakers_remove.md) - Remove Speaker from Media
  * [`list`](docs/wistia_speakers_list.md) - List Speakers
  * [`create`](docs/wistia_speakers_create.md) - Create Speaker
  * [`update`](docs/wistia_speakers_update.md) - Update Speaker
  * [`delete`](docs/wistia_speakers_delete.md) - Delete Speaker
* [`localizations`](docs/wistia_localizations.md) - Operations for localizations
  * [`list`](docs/wistia_localizations_list.md) - List Localizations
  * [`create`](docs/wistia_localizations_create.md) - Create Localization
  * [`get`](docs/wistia_localizations_get.md) - Show Localization
  * [`delete`](docs/wistia_localizations_delete.md) - Delete Localization
* [`custom-metadata-field-values`](docs/wistia_custom-metadata-field-values.md) - Operations for custom-metadata-field-values
  * [`get-medias-media-hashed-id`](docs/wistia_custom-metadata-field-values_get-medias-media-hashed-id.md) - List Custom Metadata Field Values
  * [`put-medias-media-hashed-id-custom-metadata-field-values-key`](docs/wistia_custom-metadata-field-values_put-medias-media-hashed-id-custom-metadata-field-values-key.md) - Set Custom Metadata Field Value
  * [`delete-medias-media-hashed-id-custom-metadata-field-values-key`](docs/wistia_custom-metadata-field-values_delete-medias-media-hashed-id-custom-metadata-field-values-key.md) - Clear Custom Metadata Field Value
  * [`get-medias-media-hashed-id-custom-metadata-field-values-key`](docs/wistia_custom-metadata-field-values_get-medias-media-hashed-id-custom-metadata-field-values-key.md) - Show Custom Metadata Field Value
* [`trims`](docs/wistia_trims.md) - Operations for trims
  * [`create`](docs/wistia_trims_create.md) - Create Media from Trims
* [`media-extended-audio-descriptions`](docs/wistia_media-extended-audio-descriptions.md) - Operations for media-extended-audio-descriptions
  * [`get`](docs/wistia_media-extended-audio-descriptions_get.md) - List Media Extended Audio Descriptions
  * [`get-media-extended-audio-descriptions-id`](docs/wistia_media-extended-audio-descriptions_get-media-extended-audio-descriptions-id.md) - Show Media Extended Audio Description
  * [`delete-media-extended-audio-descriptions-id`](docs/wistia_media-extended-audio-descriptions_delete-media-extended-audio-descriptions-id.md) - Delete Media Extended Audio Description
  * [`post-media-extended-audio-descriptions-order`](docs/wistia_media-extended-audio-descriptions_post-media-extended-audio-descriptions-order.md) - Order Extended Audio Description
  * [`get-media-extended-audio-descriptions-order-status-id`](docs/wistia_media-extended-audio-descriptions_get-media-extended-audio-descriptions-order-status-id.md) - Get Order Status
* [`brands`](docs/wistia_brands.md) - Operations for brands
  * [`list`](docs/wistia_brands_list.md) - List Brands
  * [`create`](docs/wistia_brands_create.md) - Create Brand
  * [`get`](docs/wistia_brands_get.md) - Show Brand
  * [`update`](docs/wistia_brands_update.md) - Update Brand
  * [`delete`](docs/wistia_brands_delete.md) - Delete Brand
  * [`apply`](docs/wistia_brands_apply.md) - Apply Brand
* [`tags`](docs/wistia_tags.md) - Operations for tags
  * [`list`](docs/wistia_tags_list.md) - List Tags
  * [`create`](docs/wistia_tags_create.md) - Create Tags
  * [`delete`](docs/wistia_tags_delete.md) - Delete Tag
* [`bulk-actions`](docs/wistia_bulk-actions.md) - Operations for bulk-actions
  * [`create`](docs/wistia_bulk-actions_create.md) - Create Bulk Actions
* [`bulk`](docs/wistia_bulk.md) - Operations for bulk
  * [`purchase`](docs/wistia_bulk_purchase.md) - Create Bulk Purchase
* [`taggings`](docs/wistia_taggings.md) - Operations for taggings
  * [`bulk-create`](docs/wistia_taggings_bulk-create.md) - Bulk Tag Media
* [`folders`](docs/wistia_folders.md) - Operations for folders
  * [`list`](docs/wistia_folders_list.md) - List Folders
  * [`create`](docs/wistia_folders_create.md) - Create Folder
  * [`get`](docs/wistia_folders_get.md) - Show Folder
  * [`update`](docs/wistia_folders_update.md) - Update Folder
  * [`delete`](docs/wistia_folders_delete.md) - Delete Folder
  * [`copy`](docs/wistia_folders_copy.md) - Copy Folder
* [`folder-sharings`](docs/wistia_folder-sharings.md) - Operations for folder-sharings
  * [`list`](docs/wistia_folder-sharings_list.md) - List Folder Sharings
  * [`create`](docs/wistia_folder-sharings_create.md) - Create Folder Sharing
  * [`get`](docs/wistia_folder-sharings_get.md) - Show Folder Sharing
  * [`update`](docs/wistia_folder-sharings_update.md) - Update Folder Sharing
  * [`delete`](docs/wistia_folder-sharings_delete.md) - Delete Folder Sharing
* [`subfolders`](docs/wistia_subfolders.md) - Operations for subfolders
  * [`list`](docs/wistia_subfolders_list.md) - List Subfolders
  * [`create`](docs/wistia_subfolders_create.md) - Create Subfolder
  * [`get`](docs/wistia_subfolders_get.md) - Show Subfolder
  * [`update`](docs/wistia_subfolders_update.md) - Update Subfolder
  * [`delete`](docs/wistia_subfolders_delete.md) - Delete Subfolder
  * [`bulk-delete`](docs/wistia_subfolders_bulk-delete.md) - Bulk Delete Subfolders
* [`channels`](docs/wistia_channels.md) - Operations for channels
  * [`list`](docs/wistia_channels_list.md) - List Channels
  * [`create`](docs/wistia_channels_create.md) - Create Channel
  * [`get`](docs/wistia_channels_get.md) - Show Channel
  * [`update`](docs/wistia_channels_update.md) - Update Channel
  * [`delete`](docs/wistia_channels_delete.md) - Delete Channel
  * [`channel-episodes`](docs/wistia_channels_channel-episodes.md) - Operations for channel-episodes
    * [`list`](docs/wistia_channels_channel-episodes_list.md) - List Channel Episodes by Channel
* [`channel-episodes`](docs/wistia_channel-episodes.md) - Operations for channel-episodes
  * [`get`](docs/wistia_channel-episodes_get.md) - Show Channel Episode
  * [`create`](docs/wistia_channel-episodes_create.md) - Create Channel Episode
  * [`list`](docs/wistia_channel-episodes_list.md) - List Channel Episodes
  * [`update`](docs/wistia_channel-episodes_update.md) - Update Channel Episode
  * [`delete`](docs/wistia_channel-episodes_delete.md) - Delete Channel Episode
  * [`publish`](docs/wistia_channel-episodes_publish.md) - Publish Channel Episode
  * [`unpublish`](docs/wistia_channel-episodes_unpublish.md) - Un-publish Channel Episode
* [`channel-collaborators`](docs/wistia_channel-collaborators.md) - Operations for channel-collaborators
  * [`list`](docs/wistia_channel-collaborators_list.md) - List Channel Collaborators
  * [`create`](docs/wistia_channel-collaborators_create.md) - Create Channel Collaborator
  * [`delete`](docs/wistia_channel-collaborators_delete.md) - Delete Channel Collaborator
* [`webinars`](docs/wistia_webinars.md) - Operations for webinars
  * [`list`](docs/wistia_webinars_list.md) - List Webinars
  * [`create`](docs/wistia_webinars_create.md) - Create Webinar
  * [`get`](docs/wistia_webinars_get.md) - Show Webinar
  * [`update`](docs/wistia_webinars_update.md) - Update Webinar
  * [`delete`](docs/wistia_webinars_delete.md) - Delete Webinar
* [`webinar-registrations`](docs/wistia_webinar-registrations.md) - Operations for webinar-registrations
  * [`get-webinars-webinar-id-registrations`](docs/wistia_webinar-registrations_get-webinars-webinar-id-registrations.md) - List Webinar Registrations
  * [`create`](docs/wistia_webinar-registrations_create.md) - Create Webinar Registration
* [`webinar-collaborators`](docs/wistia_webinar-collaborators.md) - Operations for webinar-collaborators
  * [`list`](docs/wistia_webinar-collaborators_list.md) - List Webinar Collaborators
  * [`create`](docs/wistia_webinar-collaborators_create.md) - Create Webinar Collaborator
  * [`delete`](docs/wistia_webinar-collaborators_delete.md) - Delete Webinar Collaborator
* [`account`](docs/wistia_account.md) - Operations for account
  * [`get`](docs/wistia_account_get.md) - Get Current Account
  * [`get-usage`](docs/wistia_account_get-usage.md) - Get Account Usage
  * [`get-credit-balance`](docs/wistia_account_get-credit-balance.md) - Get Credit Balance
  * [`get-brand-preload`](docs/wistia_account_get-brand-preload.md) - Get Brand Preload
  * [`update-brand-preload`](docs/wistia_account_update-brand-preload.md) - Update Brand Preload
  * [`get-brand-kit-colors`](docs/wistia_account_get-brand-kit-colors.md) - Get Brand Kit Colors
  * [`get-token-details`](docs/wistia_account_get-token-details.md) - Get Current Token
* [`contacts`](docs/wistia_contacts.md) - Operations for contacts
  * [`create`](docs/wistia_contacts_create.md) - Invite Contacts
* [`contact`](docs/wistia_contact.md) - Operations for contact
  * [`dismiss-desktop-install-prompt`](docs/wistia_contact_dismiss-desktop-install-prompt.md) - Dismiss Desktop Install Prompt
* [`account-trials`](docs/wistia_account-trials.md) - Operations for account-trials
  * [`create`](docs/wistia_account-trials_create.md) - Start Account Trial
* [`search`](docs/wistia_search.md) - Operations for search
  * [`search`](docs/wistia_search_search.md) - Search
* [`resource-urls`](docs/wistia_resource-urls.md) - Operations for resource-urls
  * [`resolve`](docs/wistia_resource-urls_resolve.md) - Resolve Resource URLs
* [`expiring-access-tokens`](docs/wistia_expiring-access-tokens.md) - Operations for expiring-access-tokens
  * [`create`](docs/wistia_expiring-access-tokens_create.md) - Create Expiring Access Token
* [`background-job-status`](docs/wistia_background-job-status.md) - Operations for background-job-status
  * [`get`](docs/wistia_background-job-status_get.md) - Show Background Job Status
* [`allowed-domains`](docs/wistia_allowed-domains.md) - Operations for allowed-domains
  * [`list`](docs/wistia_allowed-domains_list.md) - List Allowed Domains
  * [`create`](docs/wistia_allowed-domains_create.md) - Create Allowed Domain
  * [`get`](docs/wistia_allowed-domains_get.md) - Show Allowed Domain
  * [`delete`](docs/wistia_allowed-domains_delete.md) - Delete Allowed Domain
* [`remix`](docs/wistia_remix.md) - Operations for remix
  * [`post-remixes`](docs/wistia_remix_post-remixes.md) - Create Remix
  * [`get-remixes-remix-hashed-id`](docs/wistia_remix_get-remixes-remix-hashed-id.md) - Get Remix
  * [`post-remixes-remix-hashed-id-continue`](docs/wistia_remix_post-remixes-remix-hashed-id-continue.md) - Continue Remix
  * [`post-remixes-remix-hashed-id-export`](docs/wistia_remix_post-remixes-remix-hashed-id-export.md) - Export Remix
  * [`get-remix-account-status`](docs/wistia_remix_get-remix-account-status.md) - Get Remix Account Status
* [`stats-account`](docs/wistia_stats-account.md) - Operations for stats-account
  * [`get`](docs/wistia_stats-account_get.md) - Show Current Account Stats
  * [`get-stats-account-by-date`](docs/wistia_stats-account_get-stats-account-by-date.md) - Show Account Stats by Date
* [`stats-projects`](docs/wistia_stats-projects.md) - Operations for stats-projects
  * [`get`](docs/wistia_stats-projects_get.md) - Show Project Stats
* [`stats-media`](docs/wistia_stats-media.md) - Operations for stats-media
  * [`get`](docs/wistia_stats-media_get.md) - Show Media Stats
  * [`get-by-date`](docs/wistia_stats-media_get-by-date.md) - Show Media Stats by Date
  * [`get-engagement`](docs/wistia_stats-media_get-engagement.md) - Show Media Engagement
* [`stats-visitors`](docs/wistia_stats-visitors.md) - Operations for stats-visitors
  * [`list`](docs/wistia_stats-visitors_list.md) - List Visitors
  * [`get`](docs/wistia_stats-visitors_get.md) - Show Visitor
* [`stats-events`](docs/wistia_stats-events.md) - Operations for stats-events
  * [`list`](docs/wistia_stats-events_list.md) - List Events
  * [`get`](docs/wistia_stats-events_get.md) - Show Event
* [`analytics-account`](docs/wistia_analytics-account.md) - Operations for analytics-account
  * [`get`](docs/wistia_analytics-account_get.md) - Show Account Analytics
  * [`get-timeseries`](docs/wistia_analytics-account_get-timeseries.md) - Show Account Analytics Timeseries
  * [`get-top-content`](docs/wistia_analytics-account_get-top-content.md) - Show Account Top Content
  * [`get-embed-locations`](docs/wistia_analytics-account_get-embed-locations.md) - Show Account Embed Locations
  * [`find-media-by-embed-location`](docs/wistia_analytics-account_find-media-by-embed-location.md) - Find Media By Embed Location
* [`analytics-media`](docs/wistia_analytics-media.md) - Operations for analytics-media
  * [`get`](docs/wistia_analytics-media_get.md) - Show Media Analytics
  * [`get-timeseries`](docs/wistia_analytics-media_get-timeseries.md) - Show Media Analytics Timeseries
  * [`get-embed-locations`](docs/wistia_analytics-media_get-embed-locations.md) - Show Media Embed Locations
  * [`get-embed-locations-timeseries`](docs/wistia_analytics-media_get-embed-locations-timeseries.md) - Show Media Embed Locations Timeseries
  * [`get-traffic`](docs/wistia_analytics-media_get-traffic.md) - Show Media Traffic Breakdown
  * [`get-conversions`](docs/wistia_analytics-media_get-conversions.md) - Show Media Form Conversions
  * [`get-languages`](docs/wistia_analytics-media_get-languages.md) - Show Media Languages
* [`analytics-webinar`](docs/wistia_analytics-webinar.md) - Operations for analytics-webinar
  * [`get`](docs/wistia_analytics-webinar_get.md) - Show Webinar Analytics
  * [`get-registration`](docs/wistia_analytics-webinar_get-registration.md) - Show Webinar Registration Timeseries
  * [`get-traffic`](docs/wistia_analytics-webinar_get-traffic.md) - Show Webinar Traffic Breakdown
  * [`get-audience`](docs/wistia_analytics-webinar_get-audience.md) - Show Webinar Audience
  * [`get-histograms`](docs/wistia_analytics-webinar_get-histograms.md) - Show Webinar Histograms

</details>
<!-- End Commands [operations] -->

<!-- Start Request Body Input [stdinpiping] -->
## Request Body Input

Commands that accept a request body take it three ways, with a clear priority chain. The examples use `wistia folders create`; every body-bearing command works the same way.

### Individual flags (highest priority)

Each top-level body field is a flag:

```bash
wistia folders create --name 'My New Folder' --admin-email 'admin@example.com'
```

### `--body` flag

Provide the entire request body as a JSON string:

```bash
wistia folders create --body '{"name":"My New Folder","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}'
```

Individual flags override `--body` values:

```bash
# Sends {"name":"My New Folder (updated)","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}
wistia folders create --body '{"name":"My New Folder","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}' --name 'My New Folder (updated)'
```

### Stdin piping (lowest priority)

Pipe JSON into any command that accepts a request body:

```bash
echo '{"name":"My New Folder","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}' | wistia folders create
```

Individual flags override stdin values:

```bash
# Sends {"name":"My New Folder (updated)","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}
echo '{"name":"My New Folder","adminEmail":"admin@example.com","description":"My New Folder Description","public":false,"personalLibrary":false}' | wistia folders create --name 'My New Folder (updated)'
```

This is useful for chaining commands, reading from files, or scripting:

```bash
# Read body from a file
wistia folders create < request.json

# Pipe from another command
curl -s https://example.com/request.json | wistia folders create
```

### Priority

When multiple input methods are used, the priority is:

| Priority | Source | Description |
|----------|--------|-------------|
| 1 (highest) | Individual flags | `--name ...` always wins |
| 2 | `--body` flag | Whole-body JSON via flag |
| 3 (lowest) | Stdin | Piped JSON input |
<!-- End Request Body Input [stdinpiping] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL

Use `--server-url` to override the server URL entirely, bypassing any named or indexed server selection:

```bash
wistia --server-url https://custom-api.example.com account get
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
wistia account get

# Machine-readable JSON
wistia account get --output-format json

# TOON for LLM-friendly compact output
wistia account get --output-format toon

# Pipe JSON to jq without using --output-format
wistia account get --output-format json | jq '.'
```

### jq filtering

Use `--jq` to filter or transform the response inline using a [jq](https://jqlang.org) expression. This always outputs JSON and overrides `--output-format`:

```bash
# Extract a single field
wistia account get --jq '.'

# Reshape with any jq program; --raw-output prints string results as plain text (like jq -r)
wistia account get --jq '.' --raw-output
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
| `1` | Runtime/API failure |
| `2` | Usage or input failure |
| `3` | Authentication or authorization failure |

On success, the response data is printed to **stdout** as JSON. On failure, error details are printed to **stderr**.

```bash
# Capture output and handle errors
wistia account get --output-format json > output.json 2> error.log
if [ $? -ne 0 ]; then
  echo "Error occurred, see error.log"
fi
```
This CLI uses unclassified error rendering outside agent mode: pretty and TOON print the API error text as received, while `--output-format json` and `--jq` emit the unclassified envelope (including the configure `_hint` for HTTP 401/403) plus `exit_code`. Agent mode always emits the classified JSON envelope with `exit_code`, `error_type`, optional `error_reason`, `message`, `hints`, and optional `status_code` — see [For AI agents](#for-ai-agents).
<!-- End Error Handling [errors] -->

<!-- Start Diagnostics [diagnostics] -->
## Diagnostics

The CLI includes two diagnostic flags available on all commands:

### Dry Run

Preview what would be sent without making any network calls:

```bash
wistia account get --dry-run
```

In human output modes, stdout is empty and the `[DRY-RUN]` block goes to stderr. It includes:
- HTTP method and URL
- Request headers (sensitive values redacted)
- Request body preview (sensitive fields redacted)

With `--output-format json`, or with a caller-explicit `--jq`, stderr is silent and stdout is NDJSON: one compact preview object per would-be request. The jq filter is not applied, and command-declared jq presets do not select the JSON protocol.

```json
{"dry_run":true,"request":{"method":"POST","url":"https://…","headers":{"Accept":["application/json"],…},"body":<JSON value | string | null>}}
```

JSON bodies remain structured; text bodies are strings; binary bodies are `"<bytes:N>"`; absent bodies are `null`. Headers retain all values as arrays, with credentials replaced by `[REDACTED]`. Dry-run never reads the OS keychain, but credentials supplied by flag, environment, or config file still appear redacted. The command exits successfully without contacting the API.

Local mutation commands emit one `{"dry_run":true,"local":true,"command":"…","message":"…"}` object in place of a preview; filter with `select(.request)` or `select(.local)`.

### Debug

Log request and response diagnostics while running normally:

```bash
wistia account get --debug
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
- **Binary data**: binary media and canonical base64 strings are replaced with `<bytes:N>`
- **URL query**: credential-like query parameters are replaced with `[REDACTED]`

Diagnostic output should still be treated as potentially sensitive operational data.
<!-- End Diagnostics [diagnostics] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Contributions

While we value open-source contributions to this CLI, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### CLI Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/wistia/wistia-cli&utm_campaign=cli)
