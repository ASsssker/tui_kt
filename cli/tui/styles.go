package tui

import "github.com/charmbracelet/lipgloss"

func ButtonDefaultStyle() lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFE4C4"))
	return s
}

func SelectedDefaultStyle() lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("#7FFF00"))
	return s
}

func ErrorDefaultStyle() lipgloss.Style {
	s := lipgloss.NewStyle().Foreground(lipgloss.Color("#B2222"))
	return s
}
