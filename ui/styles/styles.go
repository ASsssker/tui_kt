package styles

import "github.com/charmbracelet/lipgloss"


func ButtonDefaultStyle() lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFE4C4")).Background(lipgloss.Color("#7FFFD4"))
	return s
}

func SelectedDefaultStyle() lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("#7FFF00"))
	return s
}