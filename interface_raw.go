// +build !linux

package socketcan

type RawInterface struct {
}

func (itf *RawInterface) SocketFD() int {
	return 0
}

func NewRawInterface(ifName string) (*RawInterface, error) {
	res := &RawInterface{}
	return res, nil
}

func (itf *RawInterface) Close() error {
	return nil
}

func (itf *RawInterface) Send(f CanFrame) error {
	return nil
}

func (itf *RawInterface) Receive() (CanFrame, error) {
	f := CanFrame{Dlc: 0, Data: make([]byte, 8)}
	select {}
	return f, nil
}
