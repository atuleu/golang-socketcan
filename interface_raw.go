package socketcan

import "golang.org/x/sys/unix"

type RawInterface struct {
	fd   int
	Name string
}

func (itf *RawInterface) SocketFD() int {
	return itf.fd
}

func NewRawInterface(ifName string) (*RawInterface, error) {
	res := &RawInterface{Name: ifName}
	var err error
	res.fd, err = unix.Socket(unix.AF_CAN, unix.SOCK_RAW, unix.CAN_RAW)
	if err != nil {
		return nil, err
	}

	ifIndex, err := getIfIndex(res, ifName)
	if err != nil {
		return res, err
	}

	addr := &unix.SockaddrCAN{Ifindex: ifIndex}
	err = unix.Bind(res.fd, addr)

	return res, err
}

func (itf *RawInterface) Close() error {
	return unix.Close(itf.fd)
}

func (itf *RawInterface) Send(f CanFrame) error {
	frameBytes := make([]byte, 16)
	f.putID(frameBytes)
	frameBytes[4] = f.Dlc
	copy(frameBytes[8:], f.Data)
	_, err := unix.Write(itf.fd, frameBytes)
	return err
}

func (itf *RawInterface) Receive() (CanFrame, error) {
	f := CanFrame{Data:make([]byte,8)}
	frameBytes := make([]byte, 16)
	_, err := unix.Read(itf.fd, frameBytes)
	if err != nil {
		return f, err
	}

	f.getID(frameBytes)
	f.Dlc = frameBytes[4]
	copy(f.Data, frameBytes[8:])

	return f, nil
}
