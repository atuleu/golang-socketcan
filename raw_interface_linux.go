package socketcan

import "golang.org/x/sys/unix"

type rawInterface struct {
	fd   int
	name string
}

func (itf *rawInterface) SocketFD() int {
	return itf.fd
}

func NewRawInterface(ifName string) (RawInterface, error) {
	res := &rawInterface{name: ifName}
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

func (itf *rawInterface) Close() error {
	return unix.Close(itf.fd)
}

func (itf *rawInterface) Send(f CanFrame) error {
	frameBytes := make([]byte, 16)
	f.putID(frameBytes)
	frameBytes[4] = f.Dlc
	copy(frameBytes[8:], f.Data)
	_, err := unix.Write(itf.fd, frameBytes)
	return err
}

func (itf *rawInterface) Receive() (CanFrame, error) {
	f := CanFrame{Data: make([]byte, 8)}
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
