package models

import (
	"fmt"
	"t_kt/ui/components"
	"t_kt/ui/styles"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	activeNav int
	cursor int
	nav []string
	Loaded bool
	navOpt map[int][]tea.Model
	selectedStyle lipgloss.Style
}

func InitialModel() Model {
	opt1 := []tea.Model{
		components.InitialButton("opt11", func() tea.Msg {return "opt11"}),
		components.InitialButton("opt12", func() tea.Msg {return "opt12"}),
		components.InitialCheckBox("opt13", func() tea.Msg {return "opt13"}),
	}
	opt2 := []tea.Model{
		components.InitialButton("opt21", func() tea.Msg {return "opt21"}),
		components.InitialCheckBox("opt22", func() tea.Msg {return "opt22"}),
		components.InitialButton("opt23", func() tea.Msg {return "opt23"}),
	}
	opt3 := []tea.Model{
		components.InitialButton("opt31", func() tea.Msg {return "opt31"}),
		components.InitialButton("opt32", func() tea.Msg {return "opt32"}),
		components.InitialButton("opt33", func() tea.Msg {return "opt33"}),
	}

	return Model{
		activeNav: 0,
		cursor: 0,
		nav: []string{"Opt1", "Opt2", "Opt3"},
		navOpt: map[int][]tea.Model{0:opt1, 1:opt2, 2:opt3},
		selectedStyle: styles.SelectedDefaultStyle(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		spinner.Tick,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Batch(tea.ClearScreen, tea.Quit)

		case "enter", " ":
			m.sendMsg(msg)
			return m, spinner.Tick

		case "c":
			m.sendMsg(msg)

		case "left", "a":
			status := m.toBool(msg)
			if m.activeNav != 0 && !status {
				m.activeNav--
			}
			return m, spinner.Tick

		case "right", "d":
			status := m.toBool(msg)
			if m.activeNav != len(m.nav) -1  && !status {
				m.activeNav++
			}
			return m, spinner.Tick

		case "up", "w":
			status := m.toBool(msg)
			
			if m.cursor != 0 && !status{
				m.cursor--
			}
			return m, spinner.Tick

		case "down", "s":
			status := m.toBool(msg)
			if m.cursor != len(m.navOpt[m.activeNav]) - 1 && !status{
				m.cursor++
			}
			return m, spinner.Tick
		}

	default:
		var cmd tea.Cmd
		cmd = m.sendMsg(msg)
		return m, cmd

	}

	return m, nil
}

func (m Model) View() string {
	var nav string
	for idx, n := range m.nav {
		if idx == m.activeNav {
			nav += m.selectedStyle.Render(fmt.Sprintf(" %s ", n))
		} else {
			nav += fmt.Sprintf(" %s ", n)
		}
	}
	nav += "|\n\n\n\n"

	var options string
	for idx, option := range m.navOpt[m.activeNav] {
		cursor := " "
		if idx == m.cursor {
			cursor = m.selectedStyle.Render(">")
		}
		options += fmt.Sprintf("%s %s\n\n", cursor, option.View())
	}

	return nav + options
}

func (m Model) sendMsg(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.navOpt[m.activeNav][m.cursor], cmd = m.navOpt[m.activeNav][m.cursor].Update(msg)

	return cmd
}

func (m Model) toBool(msg tea.Msg) bool {
	msg = m.sendMsg(msg)()
	s, ok := msg.(bool)
	if !ok {
		return false
	}
	return s
}