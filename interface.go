package socketcan

import (
	"time"

	"golang.org/x/sys/unix"
)

type Interface interface {
	SocketFD() int
}

func SetRecvTimeout(i Interface, timeout time.Duration) error {
	tv := unix.NsecToTimeval(timeout.Nanoseconds())
	return unix.SetsockoptTimeval(i.SocketFD(), unix.SOL_SOCKET, unix.SO_RCVTIMEO, &tv)
}

func SetSendTimeout(i Interface, timeout time.Duration) error {
	tv := unix.NsecToTimeval(timeout.Nanoseconds())
	return unix.SetsockoptTimeval(i.SocketFD(), unix.SOL_SOCKET, unix.SO_SNDTIMEO, &tv)
}
