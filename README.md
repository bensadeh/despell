<p align="center">
  <img src="assets/despell.png" width="350" />
</p>

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

# Test
despell zsh
```

### From source

Make sure that `$GOPATH/bin` is in your `PATH` variable.

Then run the following commands: 

```console
# Install
go install

# Test
despell zsh
```

### Dependencies

`despell` requires your terminal to use a [Nerd Fonts](https://www.nerdfonts.com)-patched font.

## How does it work?

At its core, `despell` is a just a [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table. 
It takes a string (process name) as input and returns a string (icon) as output. 

## Enabling despell
### How to use

`despell` was created to add a corresponding Nerd Font icon next to the currently running command in `tmux`'s
status line. To use `despell`, edit the following segments in your `~/.tmux.conf`:
- `window-status-current-format` (active window)
- `window-status-format` (inactive window)

Inside these segments, call `#(despell #W)` to map the command name to their respective icons. Have a look at the layouts 
below for an example of the configuration used in the screenshot. Note that the colors may need to be adjusted to 
your current color scheme if you're not using the `palenight` theme.

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

## Under the hood

Screenshots use:

* [iTerm2](https://iterm2.com/) for the terminal
* [Palenight Theme](https://github.com/JonathanSpeek/palenight-iterm2) for the color scheme
* [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) for the font
