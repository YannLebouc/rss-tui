# Roadmap

## Vision

`rss-tui` is a keyboard-driven RSS/Atom reader for the terminal.

The goal is to keep it:

- simple
- fast
- local-first
- pleasant to use
- easy to understand and maintain

This project is intentionally built incrementally.

---

## V1 — Core reader

Goal: build a solid, usable, local RSS reader with a simple terminal workflow.

### Feed management

- Manage feeds through a configuration file
- List configured feeds
- Store feeds locally
- Support optional feed tags / categories

### Fetching

- Manually refresh all feeds
- Manually refresh a single feed
- Parse RSS feeds
- Parse Atom feeds
- Store fetched items locally
- Avoid obvious duplicates

### Reading experience

- List items for a selected feed
- Show item title, source and publication date
- Mark item as read / unread
- Open item link in the browser

### Persistence

- Local SQLite database
- Store feeds
- Store items
- Store read/unread state
- Store last fetch metadata

### TUI / CLI

- Keyboard navigation
- Basic feed list view
- Basic item list view
- Basic item details view
- Minimal and clear terminal interface

### Quality

- Clean project structure
- README with usage instructions
- Basic tests
- Proper error handling
- Consistent logging

---

## V2 — Better daily experience

Goal: make the reader more comfortable and pleasant for regular use.

### Feed organization

- Better tag/category support
- Filter feeds by tag
- Global "all unread" view
- Sort items by date

### Fetching improvements

- Parallel fetching
- HTTP cache support (`ETag`, `Last-Modified`)
- Better duplicate detection
- Better refresh feedback and status
- Partial refresh strategies

### Reading workflow

- Mark all items in a feed as read
- Mark all visible items as read
- Better visual distinction between read and unread items

### UX improvements

- Help screen
- Clean status bar
- Confirmation prompts for destructive actions

### Import / export

- Import feeds from a simple config file

### Quality/tests

- More tests
- Improved database schema and migrations
- Better error messages

---

## Might happen

These ideas are interesting, but not required for the project to be successful.

### Features

- Auto-refresh in the background
- Search inside titles / summaries
- Saved / starred items
- Offline article content storage
- Read-it-later integration
- Feed grouping by folders

### TUI enhancements

- Vim-like navigation
- Theme support
- Mouse support

### Technical / backend ideas

- Advanced caching strategy
- Pluggable storage
- Local HTTP API
- Sync between devices
- Metrics / observability
- Packaging for multiple platforms

### Nice extras

- Article deduplication across feeds
- Feed health diagnostics
- Dead feed detection
- CLI scripting helpers
