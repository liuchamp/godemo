package main

type Actor interface {
	OnConnect() error
	OnMessage(data []byte) error
	OnError(err error)
	OnClose()
	GetMsgChan() <-chan []byte
	Done() <-chan struct{}
}
type TCPCtrl struct {
	Actor
}

func (t *TCPCtrl) Scan() {
	defer t.OnClose()
	err := t.OnConnect()
	if err != nil {
		t.OnError(err)
	}
Circle:
	for {
		select {
		case <-t.Done():
			break Circle
		case msg := <-t.GetMsgChan():
			err := t.OnMessage(msg)
			if err != nil {
				t.OnError(err)
			}
		}
	}
}

func main() {

}
