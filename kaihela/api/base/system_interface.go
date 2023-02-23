package base

type SystemInterface interface {
	ReqGateWay() (error, string)
	ConnectWebsocket(gateway string) error
	SendData(data []byte) error
	ReceiveData(data []byte) error
	SaveSessionId(sessionId string) error
}
