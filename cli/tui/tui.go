package tui

import (
	"fmt"
	"time"

	"t_kt/cli/cmd"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	activeMenu    int
	cursor        int
	menus         []string
	loaded        bool
	menuOptions   map[int][]tea.Model
	keys          keyMap
	help          help.Model
	spinner       spinner.Model
	dumpDropStyle lipgloss.Style
	selectedStyle lipgloss.Style
	errorStyle    lipgloss.Style
	info []string
	err           error
}

func (app Application) Error() string {
	return app.err.Error()
}

func InitApp() Application {
	opt1 := []tea.Model{
		InitialButton("Очистить логи", cmd.ClearLogs),
		InitialButton("Собрать логи", cmd.ExctractLogs),
		InitialButton("Закрыть клиент", cmd.KillUI),
		InitialButton("Включить дебаг", cmd.SwitchToDebug),
		InitialButton("Перезагрузить сервер", cmd.RestartSrv),
		InitialButton("Отключить сервер", cmd.StopSrv),
		InitialButton("Запустить сервер", cmd.StartSrv),
	}

	opt2 := []tea.Model{InitialText("...")}

	s := spinner.New()
	s.Spinner = spinner.Line

	return Application{
		activeMenu:    0,
		cursor:        0,
		menus:         []string{"AN", "..."},
		menuOptions:   map[int][]tea.Model{0: opt1, 1: opt2},
		keys:          keys,
		spinner:       s,
		dumpDropStyle: DumpDefaultStyle(),
		selectedStyle: SelectedDefaultStyle(),
		errorStyle:    ErrorDefaultStyle(),
	}
}

func (app Application) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, app.spinner.Tick, cmd.InitChecker, cmd.CheckDump, cmd.IninPluginChecker, cmd.CheckPluginDump)
}

func (app Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case cmd.RunResMsg:
		switch msg {
		case cmd.DumpDrop:
			app.info = append(app.info, fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), "Падение дампа"))
			return app, app.checkDump
		
		case cmd.PluginDumpDrop:
			app.info = append(app.info, fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), "Падние дампа плагина"))
			return app, app.checkPluginDump
		
		case cmd.NoDumps:
			return app, app.checkDump

		case cmd.NoPluginDumps:
			return app, app.checkPluginDump

		default:
			if msg != cmd.Successfully {
				app.err = fmt.Errorf("ошибка: %s", msg.Info)
			}
			app.loaded = false
			return app, nil
	}
	case tea.KeyMsg:
		app.err = nil
		switch msg.String() {
		case "ctrl+c", "q":
			return app, tea.Batch(tea.ClearScreen, tea.Quit)
		default:
			cmds = append(cmds, app.updateCursor(msg)...)
			return app, tea.Batch(cmds...)

		}
	case spinner.TickMsg:
		var c tea.Cmd
		app.spinner, c = app.spinner.Update(msg)
		cmds = append(cmds, c)

		return app, tea.Batch(cmds...)
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
		var spinner string
		cursor := " "
		if idx == app.cursor {
			cursor = app.selectedStyle.Render(">")

			if app.loaded {
				spinner = app.spinner.View()
			}
		}

		menusView += fmt.Sprintf("%s %s %s\n\n", cursor, option.View(), spinner)
	}

	infoView := "\n\n"
	if app.info != nil {
		for i:=len(app.info) - 1; i>=0; i-- {
			infoView += app.dumpDropStyle.Render(app.info[i]) + "\n"
		}
	}


	errView := "\n\n"
	if app.err != nil {
		errView = app.errorStyle.Render(app.Error()) + errView
	}

	view := menusView + optionsView + infoView + errView

	helpView := app.help.View(app.keys)

	return view + helpView
}
