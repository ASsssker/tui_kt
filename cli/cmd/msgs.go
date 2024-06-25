package cmd

type RunResMsg struct {
	Info string
}

var (
	Successfully = RunResMsg{Info: "succesfull"}
	DumpDrop =  RunResMsg{Info: "damp drop"}
)

func createErrMsg(err error) RunResMsg {
	return RunResMsg{Info: err.Error()}
}
 