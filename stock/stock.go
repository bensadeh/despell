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
		"Python":     {Text: python, Color: "yellow"},
		"[tmux]":     {Text: tmux, Color: "green"},
		"ack":        {Text: search, Color: "teal"},
		"atop":       {Text: graph, Color: "yellow"},
		"bash":       {Text: shell, Color: "blue"},
		"bat":        {Text: animal, Color: "magenta"},
		"cat":        {Text: animal, Color: "red"},
		"clx":        {Text: yc, Color: "color214"},
		"curl":       {Text: globe, Color: "blue"},
		"duplicate":  {Text: duplicate, Color: "yellow"},
		"fd":         {Text: search, Color: "teal"},
		"find":       {Text: search, Color: "teal"},
		"fish":       {Text: shell, Color: "blue"},
		"fzf":        {Text: search, Color: "teal"},
		"gh":         {Text: github, Color: "blue"},
		"git":        {Text: git, Color: "red"},
		"go":         {Text: golang, Color: "blue"},
		"grep":       {Text: search, Color: "teal"},
		"htop":       {Text: graph, Color: "yellow"},
		"http":       {Text: globe, Color: "blue"},
		"lazygit":    {Text: git, Color: "red"},
		"less":       {Text: text, Color: "magenta"},
		"lf":         {Text: directories, Color: "yellow"},
		"lynx":       {Text: globe, Color: "blue"},
		"man":        {Text: book, Color: "magenta"},
		"more":       {Text: text, Color: "magenta"},
		"mv":         {Text: move, Color: "yellow"},
		"nano":       {Text: pencil, Color: "magenta"},
		"nnn":        {Text: directories, Color: "yellow"},
		"node":       {Text: node, Color: "green"},
		"nvim":       {Text: vim, Color: "green"},
		"pico":       {Text: pencil, Color: "magenta"},
		"ping":       {Text: globe, Color: "blue"},
		"ranger":     {Text: directories, Color: "yellow"},
		"rg":         {Text: search, Color: "teal"},
		"rm":         {Text: trash, Color: "red"},
		"ruby":       {Text: ruby, Color: "red"},
		"scp":        {Text: computers, Color: "teal"},
		"sleep":      {Text: hourglass, Color: "teal"},
		"ssh":        {Text: computers, Color: "teal"},
		"sudo":       {Text: warning, Color: "red"},
		"tail":       {Text: down, Color: "teal"},
		"tig":        {Text: git, Color: "red"},
		"top":        {Text: graph, Color: "yellow"},
		"vi":         {Text: vim, Color: "green"},
		"vim":        {Text: vim, Color: "green"},
		"w3m":        {Text: globe, Color: "blue"},
		"wget":       {Text: globe, Color: "blue"},
		"youtube-dl": {Text: youtube, Color: "red"},
		"zsh":        {Text: shell, Color: "blue"},
	}
}
