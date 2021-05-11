package stock

import "despell/core"

const (
	animal      = ""
	book        = ""
	computers   = ""
	directories = "פּ"
	down        = ""
	duplicate   = ""
	git         = ""
	github      = ""
	globe       = ""
	golang      = ""
	graph       = ""
	hourglass   = ""
	move        = ""
	node        = ""
	pencil      = "פֿ"
	python      = ""
	ruby        = ""
	search      = ""
	shell       = ""
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
		"Python":     {Text: python, Color: "blue"},
		"[tmux]":     {Text: tmux, Color: "green"},
		"ack":        {Text: search, Color: "normal"},
		"atop":       {Text: graph, Color: "normal"},
		"bash":       {Text: shell, Color: "normal"},
		"bat":        {Text: animal, Color: "magenta"},
		"cat":        {Text: animal, Color: "red"},
		"clx":        {Text: yc, Color: "color208"},
		"curl":       {Text: globe, Color: "green"},
		"duplicate":  {Text: duplicate, Color: "normal"},
		"fd":         {Text: search, Color: "normal"},
		"find":       {Text: search, Color: "normal"},
		"fish":       {Text: shell, Color: "normal"},
		"fzf":        {Text: search, Color: "normal"},
		"gh":         {Text: github, Color: "black"},
		"git":        {Text: git, Color: "brightred"},
		"go":         {Text: golang, Color: "blue"},
		"grep":       {Text: search, Color: "normal"},
		"htop":       {Text: graph, Color: "yellow"},
		"http":       {Text: globe, Color: "green"},
		"lazygit":    {Text: git, Color: "red"},
		"less":       {Text: text, Color: "normal"},
		"lf":         {Text: directories, Color: "normal"},
		"lynx":       {Text: globe, Color: "green"},
		"man":        {Text: book, Color: "normal"},
		"more":       {Text: text, Color: "normal"},
		"mv":         {Text: move, Color: "normal"},
		"nano":       {Text: pencil, Color: "normal"},
		"nnn":        {Text: directories, Color: "normal"},
		"node":       {Text: node, Color: "normal"},
		"nvim":       {Text: vim, Color: "normal"},
		"pico":       {Text: pencil, Color: "normal"},
		"ping":       {Text: globe, Color: "normal"},
		"ranger":     {Text: directories, Color: "normal"},
		"rg":         {Text: search, Color: "normal"},
		"rm":         {Text: trash, Color: "normal"},
		"ruby":       {Text: ruby, Color: "normal"},
		"scp":        {Text: computers, Color: "normal"},
		"sleep":      {Text: hourglass, Color: "normal"},
		"ssh":        {Text: computers, Color: "normal"},
		"sudo":       {Text: warning, Color: "normal"},
		"tail":       {Text: down, Color: "normal"},
		"tig":        {Text: git, Color: "normal"},
		"top":        {Text: graph, Color: "normal"},
		"vi":         {Text: vim, Color: "normal"},
		"vim":        {Text: vim, Color: "normal"},
		"w3m":        {Text: globe, Color: "normal"},
		"wget":       {Text: globe, Color: "normal"},
		"youtube-dl": {Text: youtube, Color: "normal"},
		"zsh":        {Text: shell, Color: "normal"},
	}
}
