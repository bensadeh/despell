package nerdfonts

const (
	defaultIcon = ""

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

func getIcons() map[string]string {
	return map[string]string{
		"Python":     python,
		"[tmux]":     tmux,
		"ack":        search,
		"atop":       graph,
		"bash":       shell,
		"bat":        animal,
		"cat":        animal,
		"clx":        yc,
		"curl":       globe,
		"duplicate":  duplicate,
		"fd":         search,
		"find":       search,
		"fish":       shell,
		"fzf":        search,
		"gh":         github,
		"git":        git,
		"go":         golang,
		"grep":       search,
		"htop":       graph,
		"http":       globe,
		"lazygit":    git,
		"less":       text,
		"lf":         directories,
		"lynx":       globe,
		"man":        book,
		"more":       text,
		"mv":         move,
		"nano":       pencil,
		"nnn":        directories,
		"node":       node,
		"nvim":       vim,
		"pico":       pencil,
		"ping":       globe,
		"ranger":     directories,
		"rg":         search,
		"rm":         trash,
		"ruby":       ruby,
		"scp":        computers,
		"sleep":      hourglass,
		"ssh":        computers,
		"sudo":       warning,
		"tail":       down,
		"tig":        git,
		"top":        graph,
		"vi":         vim,
		"vim":        vim,
		"w3m":        globe,
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
