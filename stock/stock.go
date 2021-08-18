package stock

import "despell/core"

const (
	animal      = ""
	book        = ""
	computers   = ""
	directories = "פּ"
	down        = ""
	duplicate   = ""
	git         = ""
	github      = ""
	globe       = ""
	golang      = ""
	graph       = ""
	hourglass   = ""
	java        = ""
	javascript  = ""
	lua         = ""
	markdown    = ""
	move        = ""
	pencil      = "פֿ"
	python      = ""
	ruby        = ""
	search      = ""
	shell       = ""
	sync        = ""
	text        = ""
	tmux        = "﬿"
	trash       = "﫧"
	vim         = ""
	warning     = ""
	yc          = ""
	youtube     = ""
)

func GetDefaults() map[string]core.Icon {
	return map[string]core.Icon{
		"Python":     {Icon: python, Color: "yellow"},
		"[tmux]":     {Icon: tmux, Color: "green"},
		"ack":        {Icon: search, Color: "cyan"},
		"atop":       {Icon: graph, Color: "yellow"},
		"bash":       {Icon: shell, Color: "blue"},
		"bat":        {Icon: animal, Color: "magenta"},
		"cat":        {Icon: animal, Color: "red"},
		"clx":        {Icon: yc, Color: "color214"},
		"curl":       {Icon: globe, Color: "blue"},
		"duplicate":  {Icon: duplicate, Color: "yellow"},
		"exa":        {Icon: directories, Color: "yellow"},
		"fd":         {Icon: search, Color: "cyan"},
		"find":       {Icon: search, Color: "cyan"},
		"fish":       {Icon: shell, Color: "blue"},
		"fzf":        {Icon: search, Color: "cyan"},
		"gh":         {Icon: github, Color: "blue"},
		"git":        {Icon: git, Color: "red"},
		"glow":       {Icon: markdown, Color: "magenta"},
		"go":         {Icon: golang, Color: "blue"},
		"grep":       {Icon: search, Color: "cyan"},
		"htop":       {Icon: graph, Color: "yellow"},
		"http":       {Icon: globe, Color: "blue"},
		"java":       {Icon: java, Color: "red"},
		"lazygit":    {Icon: git, Color: "red"},
		"less":       {Icon: text, Color: "magenta"},
		"lf":         {Icon: directories, Color: "yellow"},
		"ls":         {Icon: directories, Color: "yellow"},
		"lua":        {Icon: lua, Color: "blue"},
		"lynx":       {Icon: globe, Color: "blue"},
		"man":        {Icon: book, Color: "magenta"},
		"more":       {Icon: text, Color: "magenta"},
		"mv":         {Icon: move, Color: "yellow"},
		"nano":       {Icon: pencil, Color: "magenta"},
		"nnn":        {Icon: directories, Color: "yellow"},
		"node":       {Icon: javascript, Color: "yellow"},
		"nvim":       {Icon: vim, Color: "green"},
		"pico":       {Icon: pencil, Color: "magenta"},
		"ping":       {Icon: globe, Color: "blue"},
		"ranger":     {Icon: directories, Color: "yellow"},
		"rg":         {Icon: search, Color: "cyan"},
		"rm":         {Icon: trash, Color: "red"},
		"rsync":      {Icon: sync, Color: "red"},
		"ruby":       {Icon: ruby, Color: "red"},
		"scp":        {Icon: computers, Color: "cyan"},
		"sleep":      {Icon: hourglass, Color: "cyan"},
		"ssh":        {Icon: computers, Color: "cyan"},
		"sudo":       {Icon: warning, Color: "red"},
		"tail":       {Icon: down, Color: "cyan"},
		"tig":        {Icon: git, Color: "red"},
		"top":        {Icon: graph, Color: "yellow"},
		"vi":         {Icon: vim, Color: "green"},
		"vim":        {Icon: vim, Color: "green"},
		"w3m":        {Icon: globe, Color: "blue"},
		"wget":       {Icon: globe, Color: "blue"},
		"youtube-dl": {Icon: youtube, Color: "red"},
		"zsh":        {Icon: shell, Color: "blue"},
	}
}
