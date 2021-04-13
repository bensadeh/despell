package nerdfonts

const (
	defaultIcon = ""

	animal    = ""
	clx       = ""
	git       = ""
	globe     = ""
	graph     = ""
	hourglass = ""
	man       = ""
	node      = ""
	pager     = ""
	pencil    = "פֿ"
	python    = ""
	ruby      = ""
	search    = ""
	shell     = ""
	computers = ""
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
		"nano":    pencil,
		"node":    node,
		"nvim":    vim,
		"pico":    pencil,
		"ping":    globe,
		"rg":      search,
		"ruby":    ruby,
		"scp":     computers,
		"ssh":     computers,
		"sleep":   hourglass,
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
