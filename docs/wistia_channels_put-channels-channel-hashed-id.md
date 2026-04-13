## wistia channels put-channels-channel-hashed-id

Update Channel

### Synopsis

Updates a channel.

```
wistia channels put-channels-channel-hashed-id [flags]
```

### Examples

```
  wistia channels put-channels-channel-hashed-id --channel-hashed-id <id>
```

### Options

```
      --author-name string         The name of the author(s) for the channel. This parameter only takes effect if podcasting is enabled for the channel.
      --auto-publish-enabled       Whether the episodes are automatically published when added to the channel. Cannot be enabled if podcasting is on.
      --body string                Request body as JSON (alternative to individual flags). Can also be provided via stdin.
      --category1 string           The primary category for the channel. This parameter only takes effect if podcasting is enabled for the channel. (options: arts > books, arts > design, arts > fashion_and_beauty, arts > food, arts > performing_arts, arts > visual_arts, business > careers, business > entrepreneurship, business > investing, business > management, business > marketing, business > non_profit, comedy > comedy_interviews, comedy > improv, comedy > stand_up, education > courses, education > how_to, education > language_learning, education > self_improvement, fiction > comedy_fiction, fiction > drama, fiction > science_fiction, health_and_fitness > alternative_health, health_and_fitness > fitness, health_and_fitness > medicine, health_and_fitness > mental_health, health_and_fitness > nutrition, health_and_fitness > sexuality, kids_and_family > education_for_kids, kids_and_family > parenting, kids_and_family > pets_and_animals, kids_and_family > stories_for_kids, leisure > animation_and_manga, leisure > automotive, leisure > aviation, leisure > crafts, leisure > games, leisure > hobbies, leisure > home_and_garden, leisure > video_games, music > music_commentary, music > music_history, music > music_interviews, news > business_news, news > daily_news, news > entertainment_news, news > news_commentary, news > politics, news > sports_news, news > tech_news, religion_and_spirituality > buddhism, religion_and_spirituality > christianity, religion_and_spirituality > hinduism, religion_and_spirituality > islam, religion_and_spirituality > judaism, religion_and_spirituality > religion, religion_and_spirituality > spirituality, science > astronomy, science > chemistry, science > earth_sciences, science > life_sciences, science > mathematics, science > natural_sciences, science > nature, science > physics, science > social_sciences, society_and_culture > documentary, society_and_culture > personal_journals, society_and_culture > philosophy, society_and_culture > places_and_travel, society_and_culture > relationships, sports > baseball, sports > basketball, sports > cricket, sports > fantasy_sports, sports > football, sports > golf, sports > hockey, sports > rugby, sports > running, sports > soccer, sports > swimming, sports > tennis, sports > volleyball, sports > wilderness, sports > wrestling, tv_and_film > after_shows, tv_and_film > film_history, tv_and_film > film_interviews, tv_and_film > film_reviews, tv_and_film > tv_reviews)
      --category2 string           The secondary category for the channel. This parameter only takes effect if podcasting is enabled for the channel. (options: arts > books, arts > design, arts > fashion_and_beauty, arts > food, arts > performing_arts, arts > visual_arts, business > careers, business > entrepreneurship, business > investing, business > management, business > marketing, business > non_profit, comedy > comedy_interviews, comedy > improv, comedy > stand_up, education > courses, education > how_to, education > language_learning, education > self_improvement, fiction > comedy_fiction, fiction > drama, fiction > science_fiction, health_and_fitness > alternative_health, health_and_fitness > fitness, health_and_fitness > medicine, health_and_fitness > mental_health, health_and_fitness > nutrition, health_and_fitness > sexuality, kids_and_family > education_for_kids, kids_and_family > parenting, kids_and_family > pets_and_animals, kids_and_family > stories_for_kids, leisure > animation_and_manga, leisure > automotive, leisure > aviation, leisure > crafts, leisure > games, leisure > hobbies, leisure > home_and_garden, leisure > video_games, music > music_commentary, music > music_history, music > music_interviews, news > business_news, news > daily_news, news > entertainment_news, news > news_commentary, news > politics, news > sports_news, news > tech_news, religion_and_spirituality > buddhism, religion_and_spirituality > christianity, religion_and_spirituality > hinduism, religion_and_spirituality > islam, religion_and_spirituality > judaism, religion_and_spirituality > religion, religion_and_spirituality > spirituality, science > astronomy, science > chemistry, science > earth_sciences, science > life_sciences, science > mathematics, science > natural_sciences, science > nature, science > physics, science > social_sciences, society_and_culture > documentary, society_and_culture > personal_journals, society_and_culture > philosophy, society_and_culture > places_and_travel, society_and_culture > relationships, sports > baseball, sports > basketball, sports > cricket, sports > fantasy_sports, sports > football, sports > golf, sports > hockey, sports > rugby, sports > running, sports > soccer, sports > swimming, sports > tennis, sports > volleyball, sports > wilderness, sports > wrestling, tv_and_film > after_shows, tv_and_film > film_history, tv_and_film > film_interviews, tv_and_film > film_reviews, tv_and_film > tv_reviews)
      --category3 string           The third category for the channel. This parameter only takes effect if podcasting is enabled for the channel. (options: arts > books, arts > design, arts > fashion_and_beauty, arts > food, arts > performing_arts, arts > visual_arts, business > careers, business > entrepreneurship, business > investing, business > management, business > marketing, business > non_profit, comedy > comedy_interviews, comedy > improv, comedy > stand_up, education > courses, education > how_to, education > language_learning, education > self_improvement, fiction > comedy_fiction, fiction > drama, fiction > science_fiction, health_and_fitness > alternative_health, health_and_fitness > fitness, health_and_fitness > medicine, health_and_fitness > mental_health, health_and_fitness > nutrition, health_and_fitness > sexuality, kids_and_family > education_for_kids, kids_and_family > parenting, kids_and_family > pets_and_animals, kids_and_family > stories_for_kids, leisure > animation_and_manga, leisure > automotive, leisure > aviation, leisure > crafts, leisure > games, leisure > hobbies, leisure > home_and_garden, leisure > video_games, music > music_commentary, music > music_history, music > music_interviews, news > business_news, news > daily_news, news > entertainment_news, news > news_commentary, news > politics, news > sports_news, news > tech_news, religion_and_spirituality > buddhism, religion_and_spirituality > christianity, religion_and_spirituality > hinduism, religion_and_spirituality > islam, religion_and_spirituality > judaism, religion_and_spirituality > religion, religion_and_spirituality > spirituality, science > astronomy, science > chemistry, science > earth_sciences, science > life_sciences, science > mathematics, science > natural_sciences, science > nature, science > physics, science > social_sciences, society_and_culture > documentary, society_and_culture > personal_journals, society_and_culture > philosophy, society_and_culture > places_and_travel, society_and_culture > relationships, sports > baseball, sports > basketball, sports > cricket, sports > fantasy_sports, sports > football, sports > golf, sports > hockey, sports > rugby, sports > running, sports > soccer, sports > swimming, sports > tennis, sports > volleyball, sports > wilderness, sports > wrestling, tv_and_film > after_shows, tv_and_film > film_history, tv_and_film > film_interviews, tv_and_film > film_reviews, tv_and_film > tv_reviews)
      --channel-hashed-id string   The hashed id of the Channel [required]
      --copyright string           The channel's copyright information. This parameter only takes effect if podcasting is enabled for the channel.
      --custom-url string          Use if embedding the channel on your own site. The custom URL ensures links always direct to your page and not Wistia's.
      --description string         The channel's description.
      --episode-format string      The format for episodes for the podcast channel. This parameter only takes effect if podcasting is enabled for the channel. (options: episodic, episodic_with_seasons, serial)
      --explicit string            Whether the channel contains explicit content. This parameter only takes effect if podcasting is enabled for the channel.
  -h, --help                       help for put-channels-channel-hashed-id
  -l, --language string            The ISO 639-1 language code for the channel. This parameter only takes effect if podcasting is enabled for the channel. (options: af, be, bg, ca, cs, da, de-at, de-ch, de-de, de-li, de-lu, de, el, en-au, en-bz, en-ca, en-gb, en-ie, en-jm, en-nz, en-ph, en-tt, en-us, en-za, en-zw, en, es-ar, es-bo, es-cl, es-co, es-cr, es-do, es-ec, es-es, es-gt, es-hn, es-mx, es-ni, es-pa, es-pe, es-pr, es-py, es-sv, es-uy, es-ve, es, et, eu, fi, fo, fr-be, fr-ca, fr-ch, fr-fr, fr-lu, fr-mc, fr, ga, gd, gl, haw, hr, hu, in, is, it-ch, it-it, it, ja, ko, mk, nl-be, nl-nl, nl, no, pl, pt-br, pt-pt, pt, ro-mo, ro-ro, ro, ru-mo, ru-ru, ru, sk, sl, sq, sr, sv-fi, sv-se, sv, tr, uk, zh-cn, zh-tw)
  -n, --name string                The display name for the channel
      --owner-email string         The email of the owner for the channel. This parameter only takes effect if podcasting is enabled for the channel.
      --owner-name string          The name of the owner for the channel. This parameter only takes effect if podcasting is enabled for the channel.
  -p, --podcast-enabled            Whether podcasting is enabled for this channel.
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

* [wistia channels](wistia_channels.md)	 - Operations for channels
