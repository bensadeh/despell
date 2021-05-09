<p align="center">
  <img src="assets/magica.png" width="350" />
</p>

<div align="center">
  
[![Latest release](https://img.shields.io/github/v/release/bensadeh/magica?style=flat&label=stable&color=e1acff&labelColor=292D3E)](https://github.com/bensadeh/magica/releases)
[![Changelog](https://img.shields.io/badge/docs-changelog-9cc4ff?style=flat&labelColor=292D3E)](https://github.com/bensadeh/magica/blob/master/CHANGELOG.md)
[![License](https://img.shields.io/github/license/bensadeh/magica?style=flat&color=c3e88d&labelColor=292D3E)](https://github.com/bensadeh/magica/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bensadeh/magica?style=flat&color=ffe585&labelColor=292D3E)](https://github.com/bensadeh/magica/blob/master/go.mod)
</div>

#

`magica` is a tool for showing window-specific icons in `tmux`'s statusline. 


<p align="center">
  <img src="assets/example2.png" width="700" />
</p>

## Installing

<img align="right" width="300" src="assets/example3.png" alt="Screenshot">

### Homebrew

```console
# Install
brew install bensadeh/magica/magica

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

`magica` relies on [Nerd Fonts](https://www.nerdfonts.com) for providing the icons.

## How does it work?

At its core, `magica` is just a [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table. It takes a 
process name as input and returns an icon as output. 

If no matches are found, a "default response" icon is returned.

## Enabling magica
### How to configure

`magica` was created to add a corresponding Nerd Font icon next to the currently running command as reported by `tmux`
in the status line. To enable the icons, you must redefine the `window-status-current-format` (active window) and 
`window-status-format` (inactive window) segments in `~/.tmux.conf`.

Inside these segments, call `#(despell #W)` to call `magica` and map the command name to an icon. Have a look at the layouts 
below for an example of the configuration used in the screenshot. Note that the colors may need to be adjusted to 
your current color scheme if you're not using the `palenight` theme.

### Examples

#### [Minimal](/examples/minimal.conf)
<p align="center">
  <img src="assets/minimal.png" width="700" />
</p>


#### [Rounded](/examples/rounded.conf)

<p align="center">
  <img src="assets/rounded.png" width="700" />
</p>


## Overriding and adding icons

Override default icons or add new mappings by creating an `overrides.json` and placing it in
`~/.config/magica/overrides.json`. You can either use [this example file](/examples) or the snippet 
below as a starting off point:

```json
{
  "unknownCommand": "?",
  "ssh": "◇",
  "zsh": "❤"
}
```


## On the choice of which commands to include

`magica` aims to include the **most commonly used commands** out of the box. Please let me know if 
you think there is a command that should be included in the default mappings.

For commands that are less common, please use the override JSON on your local system. 

## Under the hood

Screenshots use:

* [iTerm2](https://iterm2.com/) for the terminal
* [Palenight Theme](https://github.com/JonathanSpeek/palenight-iterm2) for the color scheme
* [JetBrains Mono](https://github.com/JetBrains/JetBrainsMono) for the font
