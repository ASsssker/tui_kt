package tui

import (
	"fmt"
	"t_kt/cli/cmd"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CheckBox struct {
	Title  string
	Status bool
	Action func() tea.Msg
}

func InitialCheckBox(title string, f func() tea.Msg) CheckBox {
	return CheckBox{Title: title, Action: f}
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
	Title  string
	Action func() tea.Msg
	style  lipgloss.Style
}

func InitialButton(title string, f func() tea.Msg) Button {
	return Button{Title: title,
		Action: f,
		style:  ButtonDefaultStyle(),
	}
}

func (b Button) Init() tea.Cmd {
	return nil
}

func (b Button) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmd.RunResMsg:
		return b, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			return b, b.Action
		}
	}
	return b, nil
}

func (b Button) View() string {
	return b.style.Render(b.Title)

}

type Text struct {
	Title string
}

func InitialText(text string) Text {
	return Text{Title: text}
}

func (t Text) Init() tea.Cmd {
	return nil
}

func (t Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, func() tea.Msg {
		time.Sleep(time.Second * 3)
		return cmd.Successfully
	}
}

func (t Text) View() string {
	return t.Title
}
