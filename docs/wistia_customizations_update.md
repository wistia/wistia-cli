## wistia customizations update

Update Customizations

### Synopsis

Allows for partial updates on a video’s customizations. If a value is null, then that key will be deleted from the saved customizations. If it is not null, that value will be set.

## Requires api token with one of the following permissions
```
Read, update & delete anything
```

```
wistia customizations update [flags]
```

### Examples

```
  wistia customizations update --media-id <id>
```

### Options

```
  -a, --auto-play                           If set to true, the video will play as soon as it’s ready. Note that autoplay might not work on some devices and browsers.
      --body string                         Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --controls-visible-on-load            If set to true, controls like the big play button, playbar, volume, etc. will be visible as soon as the video is embedded.
      --copy-link-and-thumbnail-enabled     If set to false, the option to “Copy Link and Thumbnail” will be removed when right-clicking on the video.
      --do-not-track                        If set to true, data for each viewing session will not be tracked.
      --email string                        Associate a specific email address with this video’s viewing sessions.
      --end-video-behavior string           Determines what happens when the video ends. Options are default (stays on the last frame), reset (shows thumbnail and controls), and loop (plays again from the start).
      --fake-fullscreen                     If set to true, the video will try to play in a pseudo-fullscreen mode on certain mobile devices.
      --fit-strategy string                 Resizes the video when there's a discrepancy between its aspect ratio and that of its parent container. Options are contain, cover, fill, and none.
      --fullscreen-button                   If set to true, the fullscreen button will be available as a video control.
      --fullscreen-on-rotate-to-landscape   If set to false, the video will not automatically go to fullscreen mode on mobile when rotated to landscape.
  -h, --help                                help for update
  -k, --key-moments                         If set to false, the key moments feature will be disabled.
      --media-id string                     The hashed ID of the video to be customized. [required]
      --muted                               If set to true, the video will start in a muted state.
      --play-button                         Indicates if the play button is visible.
      --play-pause-notifier                 If set to false, animations for the Pause and Play symbols will be removed.
      --play-suspended-off-screen           If set to false for a muted autoplay video, the video won't pause when out of view.
      --playback-rate-control               If set to false, the playback speed controls in the settings menu will be hidden.
      --playbar                             If set to true, the playbar will be available. If set to false, it will be hidden.
      --player-color string                 Changes the base color of the player. Expects a hexadecimal rgb string.
      --playlist-links                      Enables the use of specially crafted links on the page to associate with a video, turning them into a playlist.
      --playlist-loop                       If set to true and this video has a playlist, it will loop back to the first video after the last one has finished.
      --playsinline                         If set to false, videos will play within the native mobile player.
      --plugin string                       JSON object
      --preload string                      Sets the video’s preload property. Possible values are metadata, auto, none, true, and false.
      --quality-control                     If set to false, the video quality selector in the settings menu will be hidden.
      --quality-max int                     Specifies the maximum quality the video will play at.
      --quality-min int                     Specifies the minimum quality the video will play at.
  -r, --resumable string                    Determines if the video should resume from where the viewer left off. Options are true, false, and auto.
      --seo                                 If set to true, the video’s metadata will be injected into the page’s markup for SEO.
      --settings-control                    If set to true, the settings control will be available.
      --silent-auto-play string             Determines how videos handle autoplay in contexts where normal autoplay might be blocked. Options are true, allow, and false.
      --small-play-button                   boolean flag
      --still-url string                    Overrides the thumbnail image that appears before the video plays.
      --thumbnail-alt-text string           Sets the Thumbnail Alt Text for the media.
      --time string                         Sets the starting time of the video.
      --video-foam string                   JSON value (one of: boolean | { minWidth: integer, maxWidth: integer, minHeight: integer, maxHeight: integer })
      --volume float                        Sets the volume of the video.
      --volume-control                      When set to true, a volume control is available over the video.
  -w, --wmode string                        If set to transparent, the background behind the player will be transparent instead of black.
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

* [wistia customizations](wistia_customizations.md)	 - Operations for customizations
