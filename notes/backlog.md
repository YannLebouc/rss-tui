# BACKLOG

This backlog tracks the implementation progress of rss-tui.
Tasks are intentionally small to allow steady progress.

---

## Phase 1 — Project foundations

### Repository setup

- [x] Create project structure
- [x] Create `ROADMAP.md`
- [x] Create `BACKLOG.md`
- [x] Write initial `README.md`
- [x] Setup basic project layout

## Phase 2 — Feed configuration

Goal: allow users to define feeds.

### Feed config file

- [x] Decide config file location (`~/.config/rss-tui/feeds`)
- [x] Implement config loader
- [x] Support one feed URL per line
- [x] Ignore empty lines
- [x] Ignore comments (`#`)
- [x] Add optional tag parsing

Example:

`
https://blog.golang.org/feed.xml go programming
`

### Feed model

- [x] Define what should be stored from an rss feed
- [x] Create
- [x] Create `Feed`, `Channel`, `Item`, `ChannelImage`, `FeedUrl` structs
- [x] Store URL
- [x] Store optional tags
- [x] Store mandatory fields for an RSS feed

---

## Phase 3 — Fetch feeds

Goal: download RSS feeds from URLs.

### HTTP fetch

- [x] Fetch feed from URL
- [x] Handle HTTP errors
- [x] Return raw feed content

### Feed refresh

- [ ] Refresh one feed
- [x] Refresh all feeds from config

---

## Phase 4 — Parse RSS / Atom

Goal: extract items from feed content.

### RSS parsing

- [x] Parse RSS XML
- [x] Extract feed title
- [x] Extract item title
- [x] Extract item link
- [x] Extract item publication date

### Atom parsing

- [x] Detect Atom feeds
- [x] Extract Atom entry title
- [x] Extract entry link
- [x] Extract entry date

---

## Phase 5 — Local storage (SQLite)

Goal: persist feeds and items locally.

### Database setup

- [ ] Choose database location
- [ ] Create database on startup
- [ ] Create initial schema
- [ ] Implement simple migrations

### Feeds table

- [ ] Store feed URL
- [ ] Store feed title
- [ ] Store tags
- [ ] Store last fetch timestamp
- [ ] Store ETag / Last-Modified (optional for now)

### Items table

- [ ] Store item title
- [ ] Store item link
- [ ] Store publication date
- [ ] Store GUID
- [ ] Link item to feed

### Read state

- [ ] Add read/unread flag
- [ ] Add read timestamp

---

## Phase 6 — Store fetched items

Goal: save parsed items into database.

- [ ] Insert feed if not present
- [ ] Insert new items
- [ ] Skip duplicates
- [ ] Update feed last fetch time

---

## Phase 7 — CLI reader

Goal: basic reading workflow without full TUI.

### Commands

- [ ] `rss-tui refresh`
- [ ] `rss-tui feeds`
- [ ] `rss-tui list`

### Output

- [ ] List feeds
- [ ] List items for a feed
- [ ] Display unread marker
- [ ] Show item date

---

## Phase 8 — Mark items as read

Goal: basic reading workflow.

- [ ] Add command to mark item as read
- [ ] Update read state in database
- [ ] Display unread count per feed

---

## Phase 9 — Basic TUI

Goal: interactive terminal interface.

### Layout

- [ ] Feed list view
- [ ] Item list view

### Navigation

- [ ] Keyboard navigation
- [ ] Move up/down
- [ ] Select feed
- [ ] Select item

### Actions

- [ ] Open item link in browser
- [ ] Mark item read/unread
- [ ] Refresh feeds
- [ ] Quit application

---

## Phase 10 — Polishing

- [ ] Improve error messages
- [ ] Improve logging
- [ ] Improve code organization
- [ ] Add basic tests
- [ ] Clean up README

---

## Future improvements

- [ ] Global unread view
- [ ] Feed categories
- [ ] OPML import/export
- [ ] Auto refresh
- [ ] Parallel feed fetching
- [ ] Search
- [ ] Item preview
