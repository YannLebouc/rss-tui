# RSS-TUI

A terminal-based RSS feed reader built in Go.  
Browse your feeds and read articles directly from your terminal.

---

## Features

- Display list of RSS feeds
- Browse articles per feed
- Read full article content
- Keyboard-driven navigation
- Manual refresh

---

## Requirements

- Go 1.25+

---

## Installation

```bash
go install github.com/YannLebouc/rss-tui/cmd/rss-tui@latest
````

---

## Usage

Create a config file:

A config file example is available at the project's root

### Linux

```
$HOME/.config/rss-tui/feeds
```

### Windows

```
%USERPROFILE%\.config\rss-tui\feeds
```

### Example format

```
https://example.com/rss.xml
https://another-site.com/feed
```

Then run:

```bash
rss-tui
```

---

## Controls

* `j / k` → navigate
* `enter` → select
* `esc` → back
* `r` → refresh
* `q` → quit

---

## Technologies

* Go
* Bubble Tea
* Bubbles
* Lipgloss
* html2text

---

## License

MIT

