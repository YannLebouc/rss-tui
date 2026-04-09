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

- [x] Refresh all feeds from config
- [ ] Refresh one feed

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

## Phase 7 — CLI reader

Goal: basic reading workflow without full TUI.

### Commands

- [ ] `run app`

### Output

- [x] List feeds
- [x] List items for a feed
- [ ] Show item date
- [x] Convert HTML content to text
- [ ] Handle text wrapping
- [ ] Mouse scrolling available on article details
      
---

## Phase 9 — Basic TUI

Goal: interactive terminal interface.

### Layout

- [x] Feed list view
- [x] Item list view
- [x] Item details view

### Navigation

- [x] Keyboard navigation
- [x] Move up/down
- [x] Select feed
- [x] Select item
- [x] Go back to previous view

### Actions

- [ ] Open item link in browser
- [x] Refresh feeds
- [x] Quit application

---

## Phase 10 — Polishing

- [ ] Improve error messages
- [ ] Improve logging
- [ ] Improve code organization
- [ ] Add basic tests
- [ ] Clean up README

---

## Future improvements

- [ ] Parallel feed fetching
- [ ] Local storage with SQLite
- [ ] Mark item read/unread
- [ ] Global unread view
- [ ] Auto refresh
- [ ] Feed categories
- [ ] Improve UI
