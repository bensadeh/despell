package unknown

import "despell/core"

const (
	UnknownCommandKey = "unknownCommand"
)

func GetUnknownCommandIcon() core.Icon {
	return core.Icon{Icon: "", Color: "gray"}
}
