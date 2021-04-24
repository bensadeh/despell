<p align="center">
  <img src="assets/magica.png" width="350" />
</p>

#

`magica` is a tool for showing icons in your `tmux` statusline. 

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

## How does it work?

### Design



At its core, `magica` acts like a [hash map](https://en.wikipedia.org/wiki/Hash_table) lookup table. It takes a 
process name as input and returns a [nerdfont](https://www.nerdfonts.com) icon as output. 

If no matches are found, a "default response" icon is returned.

### Customizing the tmux status line

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
