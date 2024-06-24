package components

import (
	"fmt"
	"t_kt/cli/cmd/commands"
	"t_kt/ui/styles"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CheckBox struct {
	Title  string
	Status bool
	F      func() tea.Msg
}

func InitialCheckBox(title string, f func() tea.Msg) CheckBox {
	return CheckBox{Title: title, F: f}
}

func (c CheckBox) Init() tea.Cmd {
	return nil
}

func (c CheckBox) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			c.Status = !c.Status
			return c, c.getStatus
		}
	}
	return c, func() tea.Msg { return false }
}

func (c CheckBox) View() string {
	s := "%s [%s]"
	checked := " "

	if c.Status {
		checked = "X"
	}

	return fmt.Sprintf(s, c.Title, checked)
}

func (c CheckBox) getStatus() tea.Msg {
	return c.Status
}

type Button struct {
	Title   string
	spinner spinner.Model
	loaded  bool
	F       func() tea.Msg
	style   lipgloss.Style
	err error
}

func InitialButton(title string, f func() tea.Msg) Button {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
	s.Spinner = spinner.Line
	return Button{Title: title,
		spinner: s,
		F:       f,
		style:   styles.ButtonDefaultStyle(),
	}
}

func (b Button) Init() tea.Cmd {
	return b.spinner.Tick
}

func (b Button) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.RunResMsg:
		b.loaded = false
		return b, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			b.loaded = true
			return b, b.F

		case "c":
			b.loaded = false

		case "up", "w", "down", "s", "right", "d", "left", "a":
			return b, b.GetLoadStatus
		}

	default:
		var cmd tea.Cmd
		b.spinner, cmd = b.spinner.Update(msg)
		return b, cmd
	}
	return b, b.spinner.Tick
}

func (b Button) View() string {
	var s string
	if b.loaded {
		s = fmt.Sprintf("%s ", b.Title) + b.spinner.View()
	} else {
		s = b.Title
	}

	return s

}

func (b Button) GetLoadStatus() tea.Msg {
	return b.loaded
}

type Text struct {
	Title string
	Status bool
	
}

func InitialText(text string) Text {
	return Text{Title: text}
}

func (t Text) Init() tea.Cmd {
	return nil
}

func (t Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, func() tea.Msg {return t.Status}
}

func (t Text) View() string {
	return t.Title
}

