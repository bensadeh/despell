<p align="center">
  <img src="assets/despell.png" width="150" />
</p>

#

<p align="center">
  <code>despell</code> maps commands to icons in <code>tmux</code>'s status line
</p>

<p align="center">
  <img src="assets/full.png" width="700" />
</p>

### Features

- ⚡️ **Fast** — Instant lookup time (`O(1)`)
- 🌈 **Simple** — Use despell with any theme you'd like
- ⚙️ **Customizable** — Change or add any mapping

## Installing

```console
cargo install despell
```

> **Note**
> Make sure that `$HOME/.cargo/bin` is in your `PATH` environment variable


`despell` requires your terminal to use a [Nerd Fonts](https://www.nerdfonts.com)-patched font.

## How does it work?

At its core, `despell` takes a string (process name) as input and returns a string (icon) as output.

## Getting started

To use `despell` in your existing config, replace all occurrences of

```
#W
```

with

```
#(despell #W) #W
```

in your `~/.tmux.conf`.

If you don't have a `~/.tmux.conf` yet, have a look at the example configs below to get started.

## Example config

You can start using `despell` by using the example config from the screenshot. Copy of the config below into your
own `~/.tmux.conf` to get started.

```tmux
# Colors
tmux_active_fg=#a6accd
tmux_active_bg=#414863
tmux_inactive_fg=default
tmux_statusbar_bg=#232235

# Window status separator
set-window-option -g window-status-separator ''

# Status bar
set-option -g status-style bg=$tmux_statusbar_bg
set-option -g status-left ""
set-option -g status-right ""

# Justify status bar
set -g status-justify centre

# Active
set-window-option -g window-status-current-format "\
#[bg=$tmux_active_bg] #(despell -c #W)\
#[fg=$tmux_active_fg bg=$tmux_active_bg] #W "

# Inactive
set-window-option -g window-status-format "\
#[fg=$tmux_inactive_fg,bg=$tmux_statusbar_bg] #(despell -c #W)\
#[fg=$tmux_inactive_fg,dim bg=$tmux_statusbar_bg] #W "

```

## Settings

### Center alignment

To center the status line instead of left aligned (default), add the following command to
your `~/.tmux.conf`:

```tmux
# Set alignment
set -g status-justify centre
```

### Update frequency

To configure how often `tmux` refreshes its status line, add the following command to
your `~/.tmux.conf`:

```tmux
# Update the status line every X seconds
set -g status-interval 5
```

### Per-icon colors

To let `despell` set the icon color and override your theme settings, run `despell` with the
`-c` or `--color` flag:

```tmux
#(despell -c #W)
```

### Emojis

To use emojis instead of Nerd Fonts, run `despell` with the `-e` or `--emoji` flag:

```tmux
#(despell -e #W)
```

## Overriding and adding icons

To override any of the default mappings, run `despell` with the `-u` or `--custom` flag. Place a `config.toml` in
`~/.config/despell/config.toml` with your custom mappings.

All fields are optional, so if you don't use emojis, you can safely omit the fields from the `TOML`.

```toml
[default]
nerdfont = "◒"
color = "none"
emoji = "🐠"

[icons.command1]
nerdfont = "◇"
color = "blue"
emoji = "🌐"

[icons.command2]
nerdfont = "❤"
color = "magenta"
emoji = "💙"
```

## Is a mapping missing?

Let me know by opening an Issue, Discussion or PR.
