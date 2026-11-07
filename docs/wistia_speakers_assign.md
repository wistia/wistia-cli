## wistia speakers assign

Assign Speaker to Media

### Synopsis

Assigns a reusable speaker profile to a media. With `detected_speaker_id`,
every turn by that detected speaker is attributed to the profile, and
assigning a second detected speaker to the same profile merges them.
Without it, the speaker is credited on the media without naming turns.

Naming a detected speaker requires the current `speaker_data_version` as
`expected_version`, and isn't available while the account has speaker
identification turned off.


## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token, who must be able to edit the media's
transcript and view the account's speaker profiles.

An [expiring access token](https://docs.wistia.com/reference/post_expiring-token)
created with the `all:delegate_to_contact_permissions` scope and an
authorization granting the `edit-transcripts` permission on this media can
also be used.

```
wistia speakers assign [media-hashed-id] [flags]
```

### Examples

```
  wistia speakers assign --media-hashed-id <id> --speaker-profile-id abc123def4
```

### Options

```
      --body string                  Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
      --detected-speaker-id string   Only when the user identifies which voice is this person (e.g. "Speaker 1 is Annie"): that speaker's 'detected_speaker_id', such as 'default_speaker_0', from the media's diarized transcript segments. Every turn by that voice is attributed to the profile, and assigning a second detected speaker to the same profile merges them. Omit it to credit the person on the media without naming any turns; don't guess which voice is theirs.
  -e, --expected-version int         The media's current 'speaker_data_version' from its diarized transcript segments. Required with 'detected_speaker_id'; a stale value returns 409.
  -h, --help                         help for assign
  -m, --media-hashed-id string       The hashed ID of the media to assign the speaker to. (or pass it as the [media-hashed-id] argument)
  -s, --speaker-profile-id string    The reusable speaker profile to assign, as returned by List Speakers. Look the person up there first; create a profile only when they aren't listed. [required]
```

### Options inherited from parent commands

```
      --agent-mode             Enable structured errors and default TOON output for AI coding agents. Automatically enabled when a known agent environment is detected (CLAUDECODE, CURSOR_AGENT, etc.). Use --agent-mode=false to disable.
      --bearer-auth string     HTTP Bearer
      --color string           Control colored output: auto (color when output is a TTY), always, or never. Respects NO_COLOR and FORCE_COLOR env vars. (default "auto")
  -d, --debug                  Log request and response diagnostics to stderr
      --dry-run                Preview API requests without sending them (no network, no OS keychain). Human preview on stderr; with -o json or --jq, one JSON object per request on stdout. Local mutation commands (auth login, auth logout and configure) make no request: they skip prompts and writes and report a no-op (stderr, or one JSON object on stdout in the machine form)
  -H, --header stringArray     Set a custom HTTP request header (format: "Key: Value"). Can be specified multiple times.
      --include-headers        Include HTTP response headers in the output
      --interactive            Prompt for missing inputs and open guided configure/auth forms (forms fall back to line prompts on stdin off-TTY) (default true)
  -q, --jq string              Filter and transform output using a jq expression (e.g., '.name', '.items[] | .id')
      --no-interactive         Disable all interactive features (auto-prompting, explorer auto-launch, TUI forms)
      --no-keyring             Never read or write the OS keychain; store secrets in the config file instead (env: WISTIA_CLI_NO_KEYRING)
  -o, --output-format string   Specify the output format. Options: pretty, json, yaml, table, toon. (default "pretty")
      --raw-output             Write --jq string results as raw text instead of JSON strings (like jq -r); non-string results stay JSON
      --server string          Select a server by index (for indexed servers) or name (for named servers)
      --server-url string      Override the default server URL
      --timeout string         HTTP request timeout (e.g., 30s, 5m, 100ms)
      --usage                  Print the CLI Usage schema in KDL format
```

### SEE ALSO

* [wistia speakers](wistia_speakers.md)	 - Operations for speakers

### Machine interface

* `wistia speakers assign --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia speakers assign --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
