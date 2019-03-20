// +build !linux

package socketcan

type rawInterfaceStub struct {
}

func (itf *rawInterfaceStub) SocketFD() int {
	return 0
}

func NewRawInterface(ifName string) (RawInterface, error) {
	res := &rawInterfaceStub{}
	return res, nil
}

func (itf *rawInterfaceStub) Close() error {
	return nil
}

func (itf *rawInterfaceStub) Send(f CanFrame) error {
	return nil
}

func (itf *rawInterfaceStub) Receive() (CanFrame, error) {
	f := CanFrame{Dlc: 0, Data: make([]byte, 8)}
	select {}
	return f, nil
}

func IsClosedInterfaceError(err error) bool {
	return false
}
