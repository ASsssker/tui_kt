package cmd

type RunResMsg struct {
	Info string
}

var (
	Successfully   = RunResMsg{Info: "succesfull"}
	DumpDrop       = RunResMsg{Info: "dump drop"}
	PluginDumpDrop = RunResMsg{Info: "plugin dump drop"}
	NoDumps        = RunResMsg{Info: "no dumps"}
	NoPluginDumps  = RunResMsg{Info: "no plugin dumps"}
)

func createErrMsg(err error) RunResMsg {
	return RunResMsg{Info: err.Error()}
}
