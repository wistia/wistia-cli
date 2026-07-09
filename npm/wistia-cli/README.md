# Wistia CLI

[![npm](https://img.shields.io/npm/v/@wistia/wistia-cli)](https://www.npmjs.com/package/@wistia/wistia-cli)

Command-line interface for the [Wistia Data API](https://docs.wistia.com/) —
manage media, folders, channels, webinars, captions, analytics, and more from
your terminal.

## Install

```bash
npm install -g @wistia/wistia-cli

# or run without installing (agents, CI):
npx -y @wistia/wistia-cli media list

# pin an API-version channel:
npm install -g @wistia/wistia-cli@api-2026-05
```

Requires Node 18+. The package installs a prebuilt binary for your platform;
the command is `wistia`.

## Quickstart

Authenticate with a Wistia API token (create one in your account settings):

```bash
# interactive setup (stores the token in your OS keychain):
wistia configure

# or non-interactive, via environment variable or flag:
export WISTIA_CLI_BEARER_AUTH="your-api-token"
wistia media list
wistia media get --media-hashed-id abc123
```

Run `wistia --help` for the full command tree.

## For agents & CI

The CLI is designed to work well in non-interactive environments:

```bash
# zero-install invocation with env-var auth:
WISTIA_CLI_BEARER_AUTH="$TOKEN" npx -y @wistia/wistia-cli media list --output-format json
```

- Auth via `WISTIA_CLI_BEARER_AUTH` or `--bearer-auth` — no prompts.
- Structured output via `--output-format` (`json`, `yaml`, `table`, and the
  LLM-friendly `toon`), plus built-in `--jq` filtering.
- Pin behavior to an API version with the `api-YYYY-MM` dist-tags.

## How it works

`@wistia/wistia-cli` is a small launcher. The actual CLI is a static Go
binary shipped in per-platform packages
(`@wistia/wistia-cli-<platform>-<arch>`) listed as `optionalDependencies` —
npm installs exactly the one matching your OS/CPU. There are no install
scripts, and no network access happens at install time beyond the registry
itself.

Note: installing with `--omit=optional` (or `--no-optional`) skips the
binary and the CLI will not run. Reinstall without those flags, or use
another install method.

Supported platforms: macOS (arm64, x64), Linux (arm64, x64 — glibc and
musl), Windows (arm64, x64).

## Other install methods

| Channel | Command |
|---|---|
| Homebrew | `brew install wistia/tap/wistia-cli` |
| Shell installer | `curl -fsSL https://raw.githubusercontent.com/wistia/wistia-cli/main/scripts/install.sh \| bash` |
| Manual download | [GitHub Releases](https://github.com/wistia/wistia-cli/releases) |

## Versioning

Versions are dated (`YYYY.M.patch`, e.g. `2026.5.1`) and track Wistia API
versions. Each API version has an `api-YYYY-MM` dist-tag channel; `latest`
points at the newest API version.

## Links

- [Source & issues](https://github.com/wistia/wistia-cli)
- [Wistia API documentation](https://docs.wistia.com/)

MIT © Wistia, Inc.
