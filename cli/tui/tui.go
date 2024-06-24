package tui

import (
	"fmt"
	// "strings"
	"t_kt/cli/cmd/commands"
	"t_kt/ui/components"
	"t_kt/ui/styles"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	activeMenu int
	cursor int
	menus []string
	loaded bool
	menuOptions map[int][]tea.Model
	keys keyMap
	help help.Model
	selectedStyle lipgloss.Style
	errorStyle lipgloss.Style
	err error
}

func (app Application) Error() string { 
	return app.err.Error()
}

func InitApp() Application {
	opt1 := []tea.Model{
		components.InitialButton("Очистить логи", commands.ClearLogs),
		components.InitialButton("Собрать Логи", commands.ExctractLogs),
	}

	opt2 := []tea.Model{components.InitialText("...")}

	return Application{
		activeMenu: 0,
		cursor: 0,
		menus: []string{"AN", "Cloud"},
		menuOptions: map[int][]tea.Model{0:opt1, 1:opt2,},
		keys: keys,
		selectedStyle: styles.SelectedDefaultStyle(),
		errorStyle: styles.ErrorDefaultStyle(),
	}
}

func (app Application) Init() tea.Cmd {
	return tea.ClearScreen
}

func (app Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.RunResMsg:
		app.sendBack(msg)
		if msg == commands.Unsuccessfully {
			app.err = fmt.Errorf("Ошибка при выполнении команды")
		}
		return app, spinner.Tick
	case tea.KeyMsg:
		app.err = nil
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Batch(tea.ClearScreen, tea.Quit)

		case "enter", " ":
			cmd := app.sendBack(msg)
			return app, cmd

		case "c":
			app.sendBack(msg)

		case "left", "a":
			app.cursor = 0
			status := app.toBool(msg)
			if app.activeMenu != 0 && !status {
				app.activeMenu--
			}
			return app, spinner.Tick

		case "right", "d":
			app.cursor = 0
			status := app.toBool(msg)
			if app.activeMenu != len(app.menus) -1  && !status {
				app.activeMenu++
			}
			return app, spinner.Tick

		case "up", "w":
			status := app.toBool(msg)
			
			if app.cursor != 0 && !status{
				app.cursor--
			}
			return app, spinner.Tick

		case "down", "s":
			status := app.toBool(msg)
			if app.cursor != len(app.menuOptions[app.activeMenu]) - 1 && !status{
				app.cursor++
			}
			return app, spinner.Tick
		}

	default:
		var cmd tea.Cmd
		cmd = app.sendBack(msg)
		return app, cmd

	}

	return app, nil
}

func (app Application) View() string {
	var menusView, optionsView string

	for idx, menu := range app.menus {
		if idx == app.activeMenu {
			menusView += app.selectedStyle.Render(fmt.Sprintf(" %s ", menu))
		} else {
			menusView += fmt.Sprintf(" %s ", menu)
		}
	}
	menusView += "\n\n"

	for idx, option := range app.menuOptions[app.activeMenu] {
		cursor := " "
		if idx == app.cursor {
			cursor = app.selectedStyle.Render(">")
		}
		menusView += fmt.Sprintf("%s %s\n\n", cursor, option.View())
	}
	
	errView := "\n\n"
	if app.err != nil {
		errView = app.errorStyle.Render(app.Error()) + errView
	}

	view := menusView + optionsView + errView

	helpView := app.help.View(app.keys)
	
	// height := 8 - strings.Count(view, "\n") - strings.Count(helpView, "\n")

	return view +  helpView
}


func (app Application) sendBack(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	app.menuOptions[app.activeMenu][app.cursor], cmd = app.menuOptions[app.activeMenu][app.cursor].Update(msg)

	return cmd
}

func (app Application) toBool(msg tea.Msg) bool {
	msg = app.sendBack(msg)()
	s, ok := msg.(bool)
	if !ok {
		return false
	}
	return s
}