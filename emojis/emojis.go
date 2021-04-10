package emojis

const (
	defaultIcon = "ﲵ"

	animal = "🐾"
	clx    = "📰"
	git    = "🌿"
	globe  = "🌎"
	graph  = "📊"
	man    = "📕"
	node   = ""
	pager  = "📑"
	pencil = "✏️"
	python = "🐍"
	ruby   = "♦️"
	search = "🔍"
	shell  = "🐚"
	ssh    = "🛰"
	sudo   = "❗️"
	tmux   = "🔳"
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
		"pico":    pencil,
		"ping":    globe,
		"rg":      search,
		"ruby":    ruby,
		"ssh":     ssh,
		"sudo":    sudo,
		"tig":     git,
		"top":     graph,
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
