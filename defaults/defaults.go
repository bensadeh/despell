package defaults

const (
	clx    = ""
	git    = ""
	graph  = ""
	man    = ""
	pager  = ""
	pencil = "פֿ"
	ruby   = ""
	shell  = ""
	sudo   = ""
	vim    = ""
)

func GetMap() map[string]string {
	return map[string]string{
		"bash":    shell,
		"clx":     clx,
		"git":     git,
		"htop":    graph,
		"lazygit": git,
		"less":    pager,
		"man":     man,
		"more":    pager,
		"nano":    pencil,
		"nvim":    vim,
		"pico":    pencil,
		"ruby":    ruby,
		"sudo":    sudo,
		"tig":     git,
		"top":     graph,
		"vi":      vim,
		"vim":     vim,
		"zsh":     shell,
	}
}
