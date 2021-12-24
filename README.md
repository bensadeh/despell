<p align="center">
  <img src="assets/despell.png" width="350" />
</p>

<p align="center">
<a href="https://github.com/bensadeh/despell/releases" target="__blank"><img src="https://img.shields.io/github/v/release/bensadeh/despell?style=flat&label=&color=7a5ccc"</a>
<a href="/CHANGELOG.md" target="__blank"><img src="https://img.shields.io/badge/docs-changelog-30365F?style=flat&label=" alt="Changelog"></a>
<a href="/LICENSE" target="__blank"><img src="https://img.shields.io/github/license/bensadeh/despell?style=flat&color=e39400&label=" alt="License"></a>
</p>


#

<p align="center">
  <code>despell</code> puts icons in <code>tmux</code>'s status line
</p>



<p align="center">
  <img src="assets/example2.png" width="700" />
</p>

## Installing

<img align="right" width="300" src="assets/example3.png" alt="Screenshot">

### Homebrew

```console
# Install
brew install bensadeh/despell/despell

# Run
despell zsh
```

### From source

Make sure that `$GOPATH/bin` is in your `PATH` variable.

Then run the following commands: 

```console
# Install
go install

# Run
despell zsh
```

### Dependencies

`despell` requires your terminal to use a [Nerd Fonts](https://www.nerdfonts.com)-patched font.

## How does it work?

At its core, `despell` is a just a glorified [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table. 
It takes a string (process name) as input and returns a string (icon) as output. 

## Enabling despell
### How to use

`despell` was created to add a corresponding Nerd Font icon next to the currently running command as reported by `tmux`
in the status line. To use `despell`, you must edit the following segments in your `~/.tmux.conf`:
- `window-status-current-format` (active window)
- `window-status-format` (inactive window)

Inside these segments, call `#(despell #W)` to map the command name to their respective icons. Have a look at the layouts 
below for an example of the configuration used in the screenshot. Note that the colors may need to be adjusted to 
your current color scheme if you're not using the `palenight` theme.

`despell` is not limited to providing icons only for process names. It can be extended to give icons for **session names** (`#(despell #{session_name})`) and **hostnames** (`#(despell #H)`).

### Settings
#### Update frequency

To configure how often `tmux` refreshes its status line, add the following command to 
your `~/.tmux.conf`:

```tmux
# Set how often to update the status line in seconds
tmux set -g status-interval 5
```

#### Per-icon colors

To let `despell` set the icon color and override your theme settings, run `despell` with the 
`-c` flag:

```tmux
#(despell -c #W)
```

### Examples

#### [Minimal](/examples/minimal.conf)
<p align="center">
  <img src="assets/minimal.png" width="700" />
</p>

#### [Colors](/examples/colors.conf)

<p align="center">
  <img src="assets/colors.png" width="700" />
</p>

#### [Rounded](/examples/rounded.conf)

<p align="center">
  <img src="assets/rounded.png" width="700" />
</p>


## Overriding and adding icons

Override default icons or add new mappings by creating an `overrides.json` and placing it in
`~/.config/despell/overrides.json`. You can either use [this example file](/examples) or the snippet 
below as a starting off point:

```json
{
  "ssh": {"Icon": "◇", "Color": "red"},
  "zsh": {"Icon": "❤", "Color": "blue"},
  "unknownCommand": {"Icon": "?", "Color": "green"}
}
```

## Default mappings

`despell` aims to include the most commonly used commands out of the box. Please let me know if 
you think there is a command that should be included in the default mappings.

For commands that are less common, please use `overrides.json` locally. 


## Under the hood

Screenshots use:

* [iTerm2](https://iterm2.com/) for the terminal
* [Palenight Theme](https://github.com/JonathanSpeek/palenight-iterm2) for the color scheme
* [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) for the font
