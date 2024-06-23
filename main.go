package main

import (
	"t_kt/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(models.InitialModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}

}