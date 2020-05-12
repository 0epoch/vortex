package im

type WsConn struct {

}

var imServer *Server

func NewWsConn () *WsConn{
	return new(WsConn)
}

func (ws *WsConn) Run() {
	imServer = NewServer()
	imServer.Start()
}