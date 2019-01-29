package socketcan

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Interface interface {
	SocketFD() int
}

func getIfIndex(itf Interface, ifName string) (int, error) {
	ifNameRaw, err := unix.ByteSliceFromString(ifName)
	if err != nil {
		return 0, err
	}
	if len(ifNameRaw) > unix.IFNAMSIZ {
		return 0, fmt.Errorf("Maximum ifname length is %d characters", unix.IFNAMSIZ)
	}

	type ifreq struct {
		Name  [unix.IFNAMSIZ]byte
		Index int
	}
	var ifReq ifreq
	fd := itf.SocketFD()
	copy(ifReq.Name[:], ifNameRaw)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(fd),
		unix.SIOCGIFINDEX,
		uintptr(unsafe.Pointer(&ifReq)))
	if errno != 0 {
		return 0, fmt.Errorf("ioctl: %v", errno)
	}

	return ifReq.Index, nil
}

func SetRecvTimeout(i Interface, timeout time.Duration) error {
	tv := unix.NsecToTimeval(timeout.Nanoseconds())
	return unix.SetsockoptTimeval(i.SocketFD(), unix.SOL_SOCKET, unix.SO_RCVTIMEO, &tv)
}

func SetSendTimeout(i Interface, timeout time.Duration) error {
	tv := unix.NsecToTimeval(timeout.Nanoseconds())
	return unix.SetsockoptTimeval(i.SocketFD(), unix.SOL_SOCKET, unix.SO_SNDTIMEO, &tv)
}
