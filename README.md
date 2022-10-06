<p align="center">
  <img src="assets/despell.png" width="350" />
</p>

#

<p align="center">
  <code>despell</code> puts icons in <code>tmux</code>'s status line
</p>

<p align="center">
  <img src="assets/full.png" width="700" />
</p>

## Installing

| Method       | Command                                         |   
|--------------|-------------------------------------------------|
| `homebrew`   | `brew install bensadeh/despell/despell`         |     
| `go install` | `go install github.com/bensadeh/despell@latest` |  
| From source  | `go install`                                    |     

`despell` requires your terminal to use a [Nerd Fonts](https://www.nerdfonts.com)-patched font.

When using `go install`, make sure that `$GOPATH/bin` is in your `PATH` environment variable.

## How does it work?

At its core, `despell` is a just a [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table.
It takes a string (process name) as input and returns a string (icon) as output.

## Getting started

To use `despell`, replace all occurrences of

```
#W
```

with

```
#(despell #W) #W
```

in your `~/.tmux.conf`.

If you don't have a `~/.tmux.conf` yet, have a look at the example configs below to get started.

## Example configs

You can start using `despell` by copying one of the examples below into your own `~/.tmux.conf`.

#### [Colors](/examples/colors.tmux)

<p align="center">
  <img src="assets/example1.png" width="700" />
</p>

#### [No Colors](/examples/no-colors.tmux)

<p align="center">
  <img src="assets/example2.png" width="700" />
</p>

#### [Emojis](/examples/emoji.tmux)

<p align="center">
  <img src="assets/example3.png" width="700" />
</p>

### Settings

#### Update frequency

To configure how often `tmux` refreshes its status line, add the following command to
your `~/.tmux.conf`:

```tmux
# Update the status line every X seconds
tmux set -g status-interval 5
```

#### Per-icon colors

To let `despell` set the icon color and override your theme settings, run `despell` with the
`-c` flag:

```tmux
#(despell -c #W)
```

#### Emojis

To use emojis instead of Nerd Fonts, run `despell` with the `-e` flag:

```tmux
#(despell -e #W)
```

## Overriding and adding icons

Override default icons or add new mappings by creating an `overrides.json` and placing it in
`~/.config/despell/overrides.json`. You can either use [this example file](/examples) or the snippet
below as a starting off point:

`unknownCommand` is a reserved keyword for commands without mappings.

```json
{
  "unknownCommand": {
    "Icon": "?",
    "Color": "green",
    "Emoji": "❔"
  },
  "ssh": {
    "Icon": "◇",
    "Color": "red",
    "Emoji": "🌐"
  },
  "zsh": {
    "Icon": "❤",
    "Color": "blue",
    "Emoji": "💙"
  }
}
```

## Is a mapping missing?

Let me know by opening an Issue, Discussion or PR.

## Under the hood

Screenshots use:

* [iTerm2](https://iterm2.com/) for the terminal
* [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) for the font
