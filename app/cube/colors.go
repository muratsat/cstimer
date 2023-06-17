package cube

import "fmt"

const (
	Green        int = 0
	Red          int = 1
	Blue         int = 2
	Orange       int = 3
	White        int = 4
	Yellow       int = 5
	CubeSides    int = 6
	sideCount3x3 int = 9
)

var Colors [6]string = [6]string{
	"Green",
	"Red",
	"Blue",
	"Orange",
	"White",
	"Yellow",
}

var colorEmojis [6]string = [6]string{
	"🟢",
	"🔴",
	"🔵",
	"🟠",
	"⚪️",
	"🟡",
}

func ColoredSymbol(color int) string {
	if color < 0 || color >= len(colorEmojis) {
		return " "
	}
	return fmt.Sprint(color)
}
