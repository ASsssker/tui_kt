package cmd

type RunResMsg struct {
	Info string
}

var (
	Successfully = RunResMsg{Info: "succesfull"}
	ClientDumpDrop =  RunResMsg{Info: "client dump drop"}
	ServerDumpDrop = RunResMsg{Info: "server dump drop"}

)

func createErrMsg(err error) RunResMsg {
	return RunResMsg{Info: err.Error()}
}
 