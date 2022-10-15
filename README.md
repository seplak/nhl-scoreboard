# nhl-scoreboard

I like hockey, but don't like browsing sports websites.

This is a small TUI app built with [bubbletea](https://github.com/charmbracelet/bubbletea) that displays NHL scores in your terminal.

![](https://github.com/seplak/nhl-scoreboard/demo.gif)

## Requirements

You need a working installation of Go. You can install it [here](https://go.dev/doc/install)

## Build

This can be built with `go build -o <dir/executable>`

For simplicity, just `go build`

and run with: `./nhl-scoreboard`.

## To Do

Most of the functionality I want still needs to be built. In no particular order of importance:

- Don't show times for games in the past, just scores
- Standings page
- Statistics for:
  - In progress games (shots, scoring, etc.)
  - Finished games (shots, scoring, three stars, link to highlights)
- Better styling across the board
- Bold the team abbreviation and score of the winning team in final games
