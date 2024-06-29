package cmd

type RunResMsg struct {
	Info string
}

var (
	Successfully = RunResMsg{Info: "succesfull"}
	DumpDrop     = RunResMsg{Info: "dump drop"}
	NoDumps      = RunResMsg{Info: "no dumps"}
)

func clientAndServerDumpMsg(clientPath, servePath string) RunResMsg {
	info := "client dump drop: " + clientPath
	info += "\nserver dump drop: " + servePath

	return RunResMsg{Info: info}
}

func createErrMsg(err error) RunResMsg {
	return RunResMsg{Info: err.Error()}
}
