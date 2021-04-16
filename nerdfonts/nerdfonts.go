package nerdfonts

const (
	defaultIcon = ""

	animal    = ""
	clx       = ""
	computers = ""
	duplicate = ""
	remove    = "﫧"
	directory = "פּ"
	git       = ""
	github    = ""
	globe     = ""
	graph     = ""
	hourglass = ""
	man       = ""
	move      = ""
	node      = ""
	pager     = ""
	pencil    = "פֿ"
	python    = ""
	ruby      = ""
	search    = ""
	shell     = ""
	sudo      = ""
	tmux      = "﬿"
	vim       = ""
	youtube   = ""
)

func getIcons() map[string]string {
	return map[string]string{
		"Python":     python,
		"[tmux]":     tmux,
		"ack":        search,
		"bash":       shell,
		"bat":        animal,
		"cat":        animal,
		"clx":        clx,
		"duplicate":  duplicate,
		"curl":       globe,
		"fd":         search,
		"find":       search,
		"fish":       shell,
		"fzf":        search,
		"git":        git,
		"gh":         github,
		"grep":       search,
		"htop":       graph,
		"lazygit":    git,
		"less":       pager,
		"lf":         directory,
		"man":        man,
		"more":       pager,
		"mv":         move,
		"nano":       pencil,
		"nnn":        directory,
		"node":       node,
		"nvim":       vim,
		"pico":       pencil,
		"ping":       globe,
		"ranger":     directory,
		"rg":         search,
		"rm":         remove,
		"ruby":       ruby,
		"scp":        computers,
		"sleep":      hourglass,
		"ssh":        computers,
		"sudo":       sudo,
		"tail":       animal,
		"tig":        git,
		"top":        graph,
		"vi":         vim,
		"vim":        vim,
		"wget":       globe,
		"youtube-dl": youtube,
		"zsh":        shell,
	}
}

func GetIcon(key string) string {
	icons := getIcons()
	icon := icons[key]

	if icon == "" {
		return defaultIcon
	}

	return icon
}
