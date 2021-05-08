<p align="center">
  <img src="assets/magica.png" width="350" />
</p>

#

<div align="center">

[![Latest release](https://img.shields.io/github/v/release/bensadeh/magica?style=flat&label=stable&color=e1acff&labelColor=292D3E)](https://github.com/bensadeh/circumflex/releases)
[![Changelog](https://img.shields.io/badge/docs-changelog-9cc4ff?style=flat&labelColor=292D3E)](https://github.com/bensadeh/magica/blob/master/CHANGELOG.md)
[![License](https://img.shields.io/github/license/bensadeh/magica?style=flat&color=c3e88d&labelColor=292D3E)](https://github.com/bensadeh/magica/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bensadeh/circumflex?style=flat&color=ffe585&labelColor=292D3E)](https://github.com/bensadeh/circumflex/blob/master/go.mod)
</div>

`magica` is a tool for showing icons in `tmux`'s statusline. 

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

Make sure you have Go installed and that `$GOPATH/bin` is in your `PATH` variable.

Then do the following: 

```console
# Install
go install

# Run
despell zsh
```

### Dependencies

`magica` relies [Nerd Fonts](https://www.nerdfonts.com) for providing the icons.

## How does it work?

At its core, `magica` is just a [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table. It takes a 
process name as input and returns an icon as output. 

If no matches are found, a "default response" icon is returned.

## Customizing the tmux status line

To enable icons in `tmux`'s statusline, you must redefine the `window-status-current-format` (active window) and 
`window-status-format` (inactive window) segments in `.tmux.conf`.

To replace the window number with an icon of the currently running process, change:
```
#I
```
to:

```
#(despell #W)
```
## Overriding and adding icons

Override default icons or add new mappings by creating an `overrides.json` and placing it in
`~/.config/magica/overrides.json`. You can either use [this example file](/example) or the snippet 
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
