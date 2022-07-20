package unknown

import "github.com/bensadeh/despell/core"

const (
	UnknownCommandKey = "unknownCommand"
)

func GetUnknownCommandIcon() core.Icon {
	return core.Icon{Icon: "", Color: "gray"}
}
