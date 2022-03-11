/**
 * Special ioctl commands for supporting the unix interface
 * to SocketCAN interface.
 *
 */

package socketcan

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Unfortunately - this is not made public in the "unix" library so,
//   making ioctl requests is hard. I've manually recreated it here.
func ioctlPtr(fd int, req uint, arg unsafe.Pointer) (err error) {
	_, _, e1 := syscall.Syscall(unix.SYS_IOCTL, uintptr(fd), uintptr(req), uintptr(arg))
	if e1 != 0 {
		err = error(e1)
	}
	return
}

func IoctlGetIfIndex(fd int, ifNameRaw []byte) (int, error) {
	type ifreq struct {
		Name  [unix.IFNAMSIZ]byte
		Index int
	}
	var ifReq ifreq

	if len(ifNameRaw) > unix.IFNAMSIZ {
		return -1, fmt.Errorf("Maximum ifname length is %d characters", unix.IFNAMSIZ)
	}

	copy(ifReq.Name[:], ifNameRaw)
	err := ioctlPtr(fd, unix.SIOCGIFINDEX, unsafe.Pointer(&ifReq));
	if err != nil {
		return -1, err
	}

	return ifReq.Index, nil
}
