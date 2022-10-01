package defaults

import (
	"github.com/bensadeh/despell/colors"
	"github.com/bensadeh/despell/core"
	"github.com/bensadeh/despell/icons"
)

func GetDefaults() map[string]core.Icon {
	return map[string]core.Icon{
		"Python":     {Icon: icons.Python, Color: colors.Yellow, Emoji: "🐍"},
		"[tmux]":     {Icon: icons.Tmux, Color: colors.Green, Emoji: "🔩"},
		"ack":        {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"atop":       {Icon: icons.Graph, Color: colors.Yellow, Emoji: "📈"},
		"bash":       {Icon: icons.Shell, Color: colors.Normal, Emoji: "🎩"},
		"bat":        {Icon: icons.Animal, Color: colors.Magenta, Emoji: "🦇"},
		"cat":        {Icon: icons.Animal, Color: colors.Red, Emoji: "🐱"},
		"clx":        {Icon: icons.Yc, Color: colors.Orange, Emoji: "🗞"},
		"curl":       {Icon: icons.Globe, Color: colors.Blue, Emoji: "🌎"},
		"docker":     {Icon: icons.Docker, Color: colors.Blue, Emoji: "🐳"},
		"duplicate":  {Icon: icons.Duplicate, Color: colors.Yellow, Emoji: "🍃"},
		"exa":        {Icon: icons.Directories, Color: colors.Yellow, Emoji: "📁"},
		"fd":         {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"find":       {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"fish":       {Icon: icons.Shell, Color: colors.Normal, Emoji: "🐠"},
		"fzf":        {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"gh":         {Icon: icons.Github, Color: colors.Blue, Emoji: "🎋"},
		"git":        {Icon: icons.Git, Color: colors.Red, Emoji: "🎋"},
		"glow":       {Icon: icons.Markdown, Color: colors.Magenta, Emoji: "🌟"},
		"go":         {Icon: icons.Golang, Color: colors.Cyan, Emoji: "🐹"},
		"grep":       {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"htop":       {Icon: icons.Graph, Color: colors.Yellow, Emoji: "📈"},
		"http":       {Icon: icons.Globe, Color: colors.Blue, Emoji: "🌎"},
		"java":       {Icon: icons.Java, Color: colors.Red, Emoji: "☕️"},
		"lazygit":    {Icon: icons.Git, Color: colors.Red, Emoji: "🎋"},
		"less":       {Icon: icons.Text, Color: colors.Magenta, Emoji: "📕"},
		"lf":         {Icon: icons.Directories, Color: colors.Yellow, Emoji: "📁"},
		"ls":         {Icon: icons.Directories, Color: colors.Yellow, Emoji: "📁"},
		"lua":        {Icon: icons.Lua, Color: colors.Blue, Emoji: "🌙"},
		"lynx":       {Icon: icons.Globe, Color: colors.Blue, Emoji: "🌎"},
		"man":        {Icon: icons.Book, Color: colors.Magenta, Emoji: "📘"},
		"more":       {Icon: icons.Text, Color: colors.Magenta, Emoji: "📕"},
		"mv":         {Icon: icons.Move, Color: colors.Yellow, Emoji: "🚚"},
		"nano":       {Icon: icons.Pencil, Color: colors.Magenta, Emoji: "🖋"},
		"nnn":        {Icon: icons.Directories, Color: colors.Yellow, Emoji: "📁"},
		"node":       {Icon: icons.Javascript, Color: colors.Yellow, Emoji: "🌎"},
		"nvim":       {Icon: icons.Vim, Color: colors.Green, Emoji: "🖋 "},
		"pico":       {Icon: icons.Pencil, Color: colors.Magenta, Emoji: "🖋 "},
		"ping":       {Icon: icons.Globe, Color: colors.Blue, Emoji: "📡"},
		"ranger":     {Icon: icons.Directories, Color: colors.Yellow, Emoji: "📁"},
		"rg":         {Icon: icons.Search, Color: colors.Cyan, Emoji: "🔦"},
		"rm":         {Icon: icons.Trash, Color: colors.Red, Emoji: "❌"},
		"rsync":      {Icon: icons.Sync, Color: colors.Red, Emoji: "🔄"},
		"ruby":       {Icon: icons.Ruby, Color: colors.Red, Emoji: "♦️"},
		"safe-rm":    {Icon: icons.Trash, Color: colors.Red, Emoji: "❌"},
		"scp":        {Icon: icons.Computers, Color: colors.Cyan, Emoji: "🖥"},
		"sleep":      {Icon: icons.Hourglass, Color: colors.Cyan, Emoji: "💤"},
		"spin":       {Icon: icons.Down, Color: colors.Cyan, Emoji: "🌀"},
		"ssh":        {Icon: icons.Computers, Color: colors.Cyan, Emoji: "🛰 "},
		"sudo":       {Icon: icons.Warning, Color: colors.Red, Emoji: "❗️"},
		"tail":       {Icon: icons.Down, Color: colors.Cyan, Emoji: "🪵"},
		"tig":        {Icon: icons.Git, Color: colors.Red, Emoji: "🎋"},
		"tmux":       {Icon: icons.Tmux, Color: colors.Green, Emoji: "🔩"},
		"top":        {Icon: icons.Graph, Color: colors.Yellow, Emoji: "📈"},
		"vi":         {Icon: icons.Vim, Color: colors.Green, Emoji: "🖋 "},
		"vim":        {Icon: icons.Vim, Color: colors.Green, Emoji: "🖋 "},
		"w3m":        {Icon: icons.Globe, Color: colors.Blue, Emoji: "🌎"},
		"wget":       {Icon: icons.Globe, Color: colors.Blue, Emoji: "🌎"},
		"youtube-dl": {Icon: icons.YouTube, Color: colors.Red, Emoji: "🎞"},
		"zsh":        {Icon: icons.Shell, Color: colors.Normal, Emoji: "🎩"},
	}
}