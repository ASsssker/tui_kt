package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var dumpsList []string

var ErrInitDumpCheckerErr = errors.New("init dump checker error")
var ErrDumpCkeckerErr = errors.New("dump checker error")

func InitChecker() tea.Msg {
	if len(dumpsList) > 0 {
		dumpsList = []string{}
	}

	clientDumps, err := getDumps(clientLogPath)
	if err != nil {
		return createErrMsg(errors.Join(ErrInitDumpCheckerErr, err))
	}
	if len(clientDumps) > 0 {
		dumpsList = append(dumpsList, clientDumps...)
	}

	serverDumps, err := getDumps(serverLogPath)
	if err != nil {
		return createErrMsg(errors.Join(ErrInitDumpCheckerErr, err))
	}
	if len(serverDumps) > 0 {
		dumpsList = append(dumpsList, serverDumps...)
	}

	return Successfully
}

func CheckDump() tea.Msg {
	time.Sleep(time.Second * 5)

	clientDumps, err := getDumps(clientLogPath)
	if err != nil {
		return createErrMsg(errors.Join(ErrDumpCkeckerErr, err))
	}

	serverDumps, err := getDumps(serverLogPath)
	if err != nil {
		return createErrMsg(errors.Join(ErrDumpCkeckerErr, err))
	}

	totalDumps := len(clientDumps) + len(serverDumps)

	if totalDumps > len(dumpsList) {
		return DumpDrop
	}

	return NoDumps
}

func getDumps(path string) ([]string, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dumps []string
	for _, f := range dir {
		if strings.Contains(f.Name(), "dump") {
			dumps = append(dumps, fmt.Sprintf("%s/%s", path, f.Name()))
		}
	}

	return dumps, nil
}
