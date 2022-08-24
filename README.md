# nhl-scoreboard

I like hockey, but don't like browsing sports websites.

This is a work in progress TUI app built with [bubbletea](https://github.com/charmbracelet/bubbletea) that aims to provide meaningful information about upcoming, in-progress, and completed games.

## Installation

You need a working installation of Go. You can install it [here](https://go.dev/doc/install)

## Build

This can be built with `go build -o <dir/executable>`

For simplicity, just `go build`

and run with: `./nhl-scoreboard`.

## To Do

Most of the functionality I want still needs to be built. In no particular order of importance:

- Handle in-progress games; time left in game and score
- Don't show times for games in the past, just scores
- Standings page
- Refresh functionality (for in-progress games)
- Statistics for:
  - In progress games (shots, scoring, etc.)
  - Finished games (shots, scoring, three stars, link to highlights)
- Better styling across the board
