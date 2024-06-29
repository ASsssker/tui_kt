package main

import (
	"t_kt/cli/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitApp())
	if _, err := p.Run(); err != nil {
		panic(err)
	}

}
