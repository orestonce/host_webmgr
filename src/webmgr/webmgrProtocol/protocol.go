package webmgrProtocol

const (
	PTPingRequest   = `PTPingRequest`
	PTPingResponse  = `PTPingResponse`
	PTShellRequest  = `PTShellRequest`
	PTShellResponse = `PTShellResponse`
)

type PingRequest struct {
	GoOs     string
	Pwd      string
}

type ShellRequest struct {
	Cmd string
	Id  int64
}

type ShellResponse struct {
	Result string
	Id     int64
}
