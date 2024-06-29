package tui

import (
	"t_kt/cli/cmd"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (app *Application) updateCursor(msg tea.KeyMsg) []tea.Cmd {
	var cmds []tea.Cmd
	if app.loaded {
		return append(cmds, app.spinner.Tick)
	}

	switch {
	case key.Matches(msg, app.keys.Enter):
		c := app.updateSelectedOption(msg)
		app.loaded = true
		opt := app.getSelectedOption()
		switch opt := opt.(type) {
		case Button:
			if opt.Title == "Очистить логи" {
				cmds = append(cmds, c, cmd.InitChecker)

				return cmds
			}
			cmds = append(cmds, c)
		}

	case key.Matches(msg, app.keys.Up):
		if app.cursor != 0 {
			app.cursor--
		}
	case key.Matches(msg, app.keys.Down):
		if app.cursor != len(app.menuOptions[app.activeMenu])-1 {
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

	return cmds
}

func (app *Application) updateSelectedOption(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	app.menuOptions[app.activeMenu][app.cursor], cmd = app.menuOptions[app.activeMenu][app.cursor].Update(msg)

	return cmd
}

func (app *Application) getSelectedOption() tea.Model {
	return app.menuOptions[app.activeMenu][app.cursor]
}

func (app *Application) checkDump() tea.Msg {
	msg := cmd.CheckDump()
	if msg == cmd.DumpDrop {
		cmd.InitChecker()
	}
	return msg
}
