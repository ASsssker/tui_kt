package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)


func (app *Application) updateCursor(msg tea.KeyMsg) tea.Cmd {
	if app.loaded {
		return app.spinner.Tick
	}

	var cmd tea.Cmd

	switch {
	case key.Matches(msg, app.keys.Enter):
		cmd = app.updateSelectedOption(msg)
		app.loaded = true

	case key.Matches(msg, app.keys.Up):
		if app.cursor != 0 {
			app.cursor--
		}
	case key.Matches(msg, app.keys.Down):
		if app.cursor != len(app.menuOptions[app.activeMenu]) - 1 {
			app.cursor++
		}
	case key.Matches(msg, app.keys.Right):
		if app.activeMenu != len(app.menus)-1 {
			app.activeMenu++
			app.cursor = 0
		}
	case key.Matches(msg, app.keys.Left):
		if app.activeMenu != 0 {
			app.activeMenu--
			app.cursor = 0
		}
	}

	return cmd
}

func (app *Application) updateSelectedOption(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	app.menuOptions[app.activeMenu][app.cursor], cmd = app.menuOptions[app.activeMenu][app.cursor].Update(msg)

	return cmd
}