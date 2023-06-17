package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/muratsat/cstimer/app"
	"github.com/muratsat/cstimer/app/cube"
)

func test() {
	c := cube.NewCube3x3()
	moves := cube.GenerateMoves()
	c.ApplyMoves(moves)
	fmt.Println(cube.ScrambleString(moves))
	fmt.Println(c.String())
}

func main() {
	p := tea.NewProgram(app.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
