package nerdfonts

const (
	defaultIcon = ""

	animal    = ""
	clx       = ""
	computers = ""
	copy      = ""
	delete    = "﫧"
	git       = ""
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
)

func getIcons() map[string]string {
	return map[string]string{
		"Python":  python,
		"[tmux]":  tmux,
		"ack":     search,
		"bash":    shell,
		"bat":     animal,
		"cat":     animal,
		"clx":     clx,
		"cp":      copy,
		"curl":    globe,
		"fd":      search,
		"find":    search,
		"fish":    shell,
		"fzf":     search,
		"git":     git,
		"grep":    search,
		"htop":    graph,
		"lazygit": git,
		"less":    pager,
		"man":     man,
		"more":    pager,
		"mv":      move,
		"nano":    pencil,
		"node":    node,
		"nvim":    vim,
		"pico":    pencil,
		"ping":    globe,
		"rg":      search,
		"rm":      delete,
		"ruby":    ruby,
		"scp":     computers,
		"sleep":   hourglass,
		"ssh":     computers,
		"sudo":    sudo,
		"tail":    animal,
		"tig":     git,
		"top":     graph,
		"vi":      vim,
		"vim":     vim,
		"wget":    globe,
		"zsh":     shell,
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
