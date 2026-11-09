## wistia customizations update-playback

Update Playback Customizations

### Synopsis

Applies a partial update to a video's playback customizations. Only the
fields supplied are changed; sending a field as null deletes it (reverting to
the default).

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

An [expiring access token](https://docs.wistia.com/reference/post_expiring-token)
created with the `all:delegate_to_contact_permissions` scope and an
authorization granting the `edit` permission on this media can also be used.

```
wistia customizations update-playback [media-id] [flags]
```

### Examples

```
  wistia customizations update-playback --media-id <id>
```

### Options

```
  -a, --auto-play                           If set to true, the video will play as soon as it’s ready. Note that autoplay might not work on some devices and browsers.
      --body string                         Request body as JSON (alternative to individual flags). Can also be provided via stdin; @path reads a file, @- reads stdin to EOF.
  -b, --bpb-time string                     Controls when the big play button appears, expressed as a string.
      --click-for-sound                     If set to true, viewers can click to enable sound on a muted video.
      --controls-visible-on-load            If set to true, controls like the big play button, playbar, volume, etc. will be visible as soon as the video is embedded.
      --copy-link-and-thumbnail-enabled     If set to false, the option to “Copy Link and Thumbnail” will be removed when right-clicking on the video.
      --do-not-track                        If set to true, data for each viewing session will not be tracked.
      --email string                        Associate a specific email address with this video’s viewing sessions.
      --end-video-behavior string           Determines what happens when the video ends. Options are "default" (stays on the last frame), "reset" (shows thumbnail and controls), and "loop" (plays again from the start).
      --fake-full-screen                    If set to true, the video will try to play in a pseudo-fullscreen mode on certain mobile devices.
      --fullscreen-button                   If set to true, the fullscreen button will be available as a video control.
      --fullscreen-on-rotate-to-landscape   If set to false, the video will not automatically go to fullscreen mode on mobile when rotated to landscape.
  -g, --google-analytics string             Google Analytics tracking configuration to associate with this video’s viewing sessions.
  -h, --help                                help for update-playback
      --hls                                 If set to true, HLS adaptive bitrate streaming is enabled.
  -k, --key-moments                         If set to false, the key moments feature will be disabled.
  -m, --media-id string                     The hashed ID of the video to be customized. (or pass it as the [media-id] argument)
      --muted                               If set to true, the video will start in a muted state.
      --play-button                         Indicates if the play button is visible.
      --play-pause-notifier                 If set to false, animations for the Pause and Play symbols will be removed.
      --play-suspended-off-screen           If set to false for a muted autoplay video, the video won’t pause when out of view.
      --playback-rate-control               If set to false, the playback speed controls in the settings menu will be hidden.
      --playbar                             If set to true, the playbar will be available. If set to false, it will be hidden.
      --playlist-links                      Enables the use of specially crafted links on the page to associate with a video, turning them into a playlist.
      --playlist-loop                       If set to true and this video has a playlist, it will loop back to the first video after the last one has finished.
      --playsinline                         If set to false, videos will play within the native mobile player.
      --preload string                      Sets the video’s preload property. Possible values are metadata, auto, none, true, and false.
      --quality-control                     If set to false, the video quality selector in the settings menu will be hidden.
      --quality-max int                     Specifies the maximum quality the video will play at.
      --quality-min int                     Specifies the minimum quality the video will play at.
  -r, --resumable string                    Determines if the video should resume from where the viewer left off. Options are "true", "false", and "auto".
      --seo                                 If set to true, the video’s metadata will be injected into the page’s markup for SEO.
      --settings-control                    If set to true, the settings control will be available.
      --silent-auto-play string             Determines how videos handle autoplay in contexts where normal autoplay might be blocked. Options are "true", "allow", and "false".
      --small-play-button                   If set to true, the small play button control is shown.
      --spherical                           If set to true, the video is rendered as a spherical (360-degree) video.
  -t, --time string                         Sets the starting time of the video.
      --video-foam string                   JSON value (one of: boolean | { "minWidth": integer, "maxWidth": integer, "minHeight": integer, "maxHeight": integer })
      --video-quality string                Sets the default video quality the video will play at.
      --volume float                        Sets the volume of the video.
      --volume-control                      When set to true, a volume control is available over the video.
  -w, --wmode string                        If set to transparent, the background behind the player will be transparent instead of black.
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

* [wistia customizations](wistia_customizations.md)	 - Operations for customizations

### Machine interface

* `wistia customizations update-playback --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia customizations update-playback --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
