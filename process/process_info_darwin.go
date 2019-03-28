package process

type darwinProcessInfo struct {
	Id              int    `json:"id" statusFileKey:"Pid"`
	Command         string `json:"command"  statusFileKey:"Name"`
	Executable      string `json:"executable"`
	ParentProcessId int    `json:"parentProcessId  statusFileKey:"Name"`
}

func (dpi darwinProcessInfo) GetId() int {
	return dpi.Id
}

func (dpi darwinProcessInfo) GetCommand() string {
	return dpi.Command
}

func (dpi darwinProcessInfo) GetParentProccessId() int {
	return dpi.ParentProcessId
}

func (dpi darwinProcessInfo) GetExecutable() string {
	return dpi.Executable
}

func processInfo(pid int) (darwinProcessInfo, error) {
	dpi := darwinProcessInfo{}
	return dpi, nil
}

func processExe(pid int) (string, error) {
	return "", nil
}
