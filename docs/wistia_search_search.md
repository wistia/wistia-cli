## wistia search search

Search

### Synopsis

Searches media, folders, subfolders, channels, channel episodes, and webinars
with filters, sorting, and paging, ranking every type in one list. Matches names,
descriptions, and transcripts; `total_count` and `counts_by_type` count every
match, not just the returned page, and `per_page=0` returns only the counts.

Releases through 2026-09 serve the previous search, which takes `q` and nests
results by type under `data`. Accounts whose content isn't indexed for this
search receive a 403 and can use the previous search by sending
`X-Wistia-API-Version: 2026-09`.

## Requires api token with one of the following permissions
```
Read all data
```

A token with `Read all folder and media data` is enough when you pass
`content_types` without `webinar`.

Tokens with the "Act with a team member's permissions" permission
(`all:delegate_to_contact_permissions` scope) can also be used. Requests
made with such a token are authorized using the permissions of the
contact assigned to the token.

[Expiring access tokens](https://docs.wistia.com/reference/post_expiring-token)
created with authorizations cannot be used: this search responds with a 403,
whatever the authorizations grant.

```
wistia search search [flags]
```

### Examples

```
  wistia search search
```

### Options

```
      --archived                               Only archived (true) or only non-archived (false) content.
      --assignees stringArray                  Assignees, as 'contact_<hashed ID>' for a contact or 'group_<hashed ID>' for a contact group.
      --containers stringArray                 Folder or channel names to search within. Doesn't match subfolder names; use 'in_container' for those.
      --content-types stringArray              Record types to include. 'folder' matches folders and subfolders together. For a generic request for media, files, or assets, use 'media' to leave out folders and channels without assuming a media subtype. (options: media, folder, subfolder, channel, channel_episode, webinar)
      --created-after string                   Only content created at or after this time: ISO 8601 at any precision ("2026", "2026-08", "2026-08-01", or "2026-08-01T09:30:00Z") or epoch seconds. A truncated value starts at the beginning of its year, month, or day, in UTC.
      --created-before string                  Only content created at or before this time, in the same forms as 'created_after'. A truncated value runs to the end of its year, month, or day, so 'created_after=2026-08' with 'created_before=2026-08' covers all of August 2026. Both bounds are inclusive, so "before 2026" is 'created_before=2025', and "before March 2026" is 'created_before=2026-02'.
      --custom-metadata string                 Filter media by custom metadata field value, keyed by field key. Select and text fields take a value ('custom_metadata[region]=emea') or an array of values matched as OR. Boolean fields take 'true' or 'false'. Number and time fields take an exact number or a range ('custom_metadata[budget][min]=100&custom_metadata[budget][max]=500'), and date fields a 'YYYY-MM-DD' date or a range with 'after' and 'before'. Any field accepts 'custom_metadata[region][exists]=false' to find media missing the field, or 'true' for media with any value. Only media can match, so other content types are excluded. Only available on accounts with access to custom metadata.
      --dubbed-languages stringArray           Language codes the media is dubbed into.
      --duration-group stringArray             Duration buckets, in seconds. Pass every bucket a range covers (for example, under 5 minutes is '0-59', '60-179', and '180-299'). (options: 0-59, 60-179, 180-299, 300-1799, 1800-3599, 3600+)
  -e, --exact-match                            Match 'query' as an exact phrase only, with no prefix or fuzzy expansion, so "power" doesn't match "powerful" (default false).
      --has-transcript                         Only media with at least one published transcript (true) or none (false), which restricts results to media. Pair it with 'media_types', and leave out 'Other': images, documents, and logos never have captions. This checks whether a transcript exists, not whether captions are on in the player.
  -h, --help                                   help for search
      --in-container string                    Hashed ID of a folder, subfolder, or channel to search inside; you don't need to know which of the three it is. Cannot be combined with 'containers'. A folder scope also matches that folder's subfolder records, so pair it with 'media_types' or 'content_types' of 'media' when counting or listing media. A channel scope matches that channel's episodes, not media.
      --include string                         Pass 'custom_metadata' to add each media result's 'custom_metadata_field_values', in the same shape as the Get Custom Metadata Field Values endpoint. Only available on accounts with access to custom metadata; other accounts receive a 403. (options: custom_metadata)
      --include-highlights                     Attach transcript matches to each record: timestamped snippets of where the query matched spoken words (default false; implied by 'transcript_only').
      --include-language-counts                Attach 'transcript_language_counts' and 'dubbed_language_counts': the 100 most common languages in the matched set, each with a label and media count (default false).
  -m, --media-types stringArray                Media subtypes to include. Set this only when the request is limited to video, audio, or another media type; "recording", "media", "file", and "asset" don't imply 'Video'. (options: Video, Audio, Other)
      --not-dubbed-languages stringArray       Leave out media already dubbed into any of these language codes.
      --not-transcript-languages stringArray   Leave out media with a transcript in any of these language codes. Includes media with no published transcripts; leave 'has_transcript' unset unless the request is limited to media with or without captions.
      --origin stringArray                     How the media was created. (options: Editor, Live, Record)
      --page int                               The page of results to return (default 1). Pages reach only the first 10,000 matches, so 'page' × 'per_page' can't pass 10,000. To reach later records, change the filters or sort and start again at page 1.
      --per-page int                           Records per page (default 10, maximum 100). Pass 0 to return only 'total_count' and 'counts_by_type'.
      --play-count-group stringArray           Play-count buckets. Pass every bucket a range covers (for example, more than 100 plays is '101-1000', '1001-10000', '10001-50000', '50001-500000', and '500001+'). (options: 0, 1-100, 101-1000, 1001-10000, 10001-50000, 50001-500000, 500001+)
      --query string                           Free-text search over names, descriptions, and transcripts. Omit to match everything.
      --share-link-visibility stringArray      Share-link visibility. (options: locked, account, unlocked, domain_verified)
      --sort string                            Ordering. Omit for relevance ranking, which is what you want alongside a 'query'; set it for content by recency or name ("the latest video" is 'sort=newest' with 'per_page=1'). (options: newest, oldest, recently_updated, name)
      --speaker-profile-ids stringArray        Speaker profile IDs, from List Speakers. Pass every matching profile; this filters assigned speakers, not people mentioned by name. A list with no ID returns 400.
      --tags stringArray                       Tag names the content must have.
      --transcript-languages stringArray       Language codes present in the transcript.
      --transcript-only                        Search only spoken words in transcripts. Results are media ranked by transcript relevance, each with timestamped transcript matches (default false).
      --unresolved-comments                    Only content with unresolved comments.
      --uploaded-by stringArray                Hashed IDs of the contacts who uploaded the media.
  -w, --workflow-status stringArray            Workflow status values.
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

* [wistia search](wistia_search.md)	 - Operations for search

### Machine interface

* `wistia search search --usage` — this command's flags, defaults and env vars as machine-readable KDL
* `wistia search search --dry-run` — preview the request without OS-keychain access or a network call (human preview on stderr)
* `--dry-run --output-format json` (or a caller-explicit `--jq`) writes one preview object per request as NDJSON on stdout; jq is not applied to previews
* `--output-format json` or `--jq <expr>` for machine-readable live output; in agent mode errors are a JSON envelope on stderr

Exit codes: 0 ok · 1 runtime · 2 usage · 3 authentication/authorization
