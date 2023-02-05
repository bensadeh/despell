package stock

import (
	"github.com/bensadeh/despell/color"
	"github.com/bensadeh/despell/core"
	"github.com/bensadeh/despell/emoji"
	"github.com/bensadeh/despell/nerdfont"
)

func GetDefaultIcon() core.Icon {
	return core.Icon{Icon: nerdfont.Shell, Color: color.Normal, Emoji: emoji.TopHat}
}

func GetStockMappings() map[string]core.Icon {
	return map[string]core.Icon{
		"Python":     {Icon: nerdfont.Python, Color: color.Yellow, Emoji: emoji.Snake},
		"[tmux]":     {Icon: nerdfont.Tmux, Color: color.Green, Emoji: emoji.NutAndBold},
		"ack":        {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"atop":       {Icon: nerdfont.Graph, Color: color.Yellow, Emoji: emoji.Microscope},
		"bat":        {Icon: nerdfont.Animal, Color: color.Magenta, Emoji: emoji.Bat},
		"cargo":      {Icon: nerdfont.Rust, Color: color.Red, Emoji: emoji.Crab},
		"cat":        {Icon: nerdfont.Animal, Color: color.Red, Emoji: emoji.Cat},
		"clx":        {Icon: nerdfont.Yc, Color: color.Orange, Emoji: emoji.Newspaper},
		"curl":       {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"docker":     {Icon: nerdfont.Docker, Color: color.Blue, Emoji: emoji.Whale},
		"duplicate":  {Icon: nerdfont.Duplicate, Color: color.Yellow, Emoji: emoji.Leaves},
		"emacs":      {Icon: nerdfont.Emacs, Color: color.Magenta, Emoji: emoji.Pen},
		"exa":        {Icon: nerdfont.Directories, Color: color.Yellow, Emoji: emoji.Folder},
		"fd":         {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"find":       {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"fzf":        {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"gh":         {Icon: nerdfont.Github, Color: color.Blue, Emoji: emoji.TanabataTree},
		"git":        {Icon: nerdfont.Git, Color: color.Red, Emoji: emoji.Wood},
		"glow":       {Icon: nerdfont.Markdown, Color: color.Magenta, Emoji: emoji.Star},
		"go":         {Icon: nerdfont.Golang, Color: color.Cyan, Emoji: emoji.HamsterFace},
		"grep":       {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"htop":       {Icon: nerdfont.Graph, Color: color.Yellow, Emoji: emoji.Microscope},
		"http":       {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"java":       {Icon: nerdfont.Java, Color: color.Red, Emoji: emoji.Coffee},
		"lazygit":    {Icon: nerdfont.Git, Color: color.Red, Emoji: emoji.TanabataTree},
		"less":       {Icon: nerdfont.Text, Color: color.Magenta, Emoji: emoji.BookRed},
		"lf":         {Icon: nerdfont.Directories, Color: color.Yellow, Emoji: emoji.Folder},
		"ls":         {Icon: nerdfont.Directories, Color: color.Yellow, Emoji: emoji.Folder},
		"lua":        {Icon: nerdfont.Lua, Color: color.Blue, Emoji: emoji.Moon},
		"lynx":       {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"man":        {Icon: nerdfont.Book, Color: color.Magenta, Emoji: emoji.BookBlue},
		"more":       {Icon: nerdfont.Text, Color: color.Magenta, Emoji: emoji.BookRed},
		"mv":         {Icon: nerdfont.Move, Color: color.Yellow, Emoji: emoji.MovingTruck},
		"nano":       {Icon: nerdfont.Pencil, Color: color.Magenta, Emoji: emoji.Pen},
		"nnn":        {Icon: nerdfont.Directories, Color: color.Yellow, Emoji: emoji.Folder},
		"node":       {Icon: nerdfont.Node, Color: color.Green, Emoji: emoji.Globe},
		"npm":        {Icon: nerdfont.Package, Color: color.Red, Emoji: emoji.Package},
		"nvim":       {Icon: nerdfont.Vi, Color: color.Green, Emoji: emoji.Pen},
		"pico":       {Icon: nerdfont.Pencil, Color: color.Magenta, Emoji: emoji.Pen},
		"ping":       {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"ranger":     {Icon: nerdfont.Directories, Color: color.Yellow, Emoji: emoji.Folder},
		"rg":         {Icon: nerdfont.Search, Color: color.Cyan, Emoji: emoji.MagnifyingGlass},
		"rm":         {Icon: nerdfont.Trash, Color: color.Red, Emoji: emoji.CrossMark},
		"rsync":      {Icon: nerdfont.Sync, Color: color.Red, Emoji: emoji.AnticlockwiseArrows},
		"ruby":       {Icon: nerdfont.Ruby, Color: color.Red, Emoji: emoji.DiamondsSuit},
		"rustc":      {Icon: nerdfont.Rust, Color: color.Red, Emoji: emoji.Crab},
		"safe-rm":    {Icon: nerdfont.Trash, Color: color.Red, Emoji: emoji.CrossMark},
		"scp":        {Icon: nerdfont.Computers, Color: color.Cyan, Emoji: emoji.Computer},
		"sleep":      {Icon: nerdfont.Hourglass, Color: color.Cyan, Emoji: emoji.Zzz},
		"spin":       {Icon: nerdfont.Down, Color: color.Cyan, Emoji: emoji.Spin},
		"ssh":        {Icon: nerdfont.Computers, Color: color.Cyan, Emoji: emoji.Key},
		"sudo":       {Icon: nerdfont.Warning, Color: color.Red, Emoji: emoji.RedExclamationMark},
		"tail":       {Icon: nerdfont.Down, Color: color.Cyan, Emoji: emoji.DownPointingTriangle},
		"tig":        {Icon: nerdfont.Git, Color: color.Red, Emoji: emoji.TanabataTree},
		"tmux":       {Icon: nerdfont.Tmux, Color: color.Green, Emoji: emoji.NutAndBold},
		"top":        {Icon: nerdfont.Graph, Color: color.Yellow, Emoji: emoji.Microscope},
		"vi":         {Icon: nerdfont.Vi, Color: color.Green, Emoji: emoji.Pen},
		"vim":        {Icon: nerdfont.Vi, Color: color.Green, Emoji: emoji.Pen},
		"w3m":        {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"wget":       {Icon: nerdfont.Globe, Color: color.Blue, Emoji: emoji.Globe},
		"youtube-dl": {Icon: nerdfont.YouTube, Color: color.Red, Emoji: emoji.FilmFrames},
	}
}
