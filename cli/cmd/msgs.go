package cmd

type RunResMsg struct {
	Info string
}

var (
	Successfully = RunResMsg{Info: "succesfull"}
	DumpDrop     = RunResMsg{Info: "dump drop"}
	NoDumps      = RunResMsg{Info: "no dumps"}
)

func createErrMsg(err error) RunResMsg {
	return RunResMsg{Info: err.Error()}
}
