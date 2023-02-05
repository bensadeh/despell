package unknown

import (
	"github.com/bensadeh/despell/colors"
	"github.com/bensadeh/despell/core"
	"github.com/bensadeh/despell/emoji"
	"github.com/bensadeh/despell/nerdfont"
)

const (
	MissingCommandKey = "unknownCommand"
)

func GetUnknownCommandIcon() core.Icon {
	return core.Icon{Icon: nerdfont.Shell, Color: colors.Normal, Emoji: emoji.TopHat}
}
