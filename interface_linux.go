package socketcan

import (
	"golang.org/x/sys/unix"
)

func getIfIndex(itf Interface, ifName string) (int, error) {
	ifNameRaw, err := unix.ByteSliceFromString(ifName)
	if err != nil {
		return 0, err
	}
	fd := itf.SocketFD()
	return IoctlGetIfIndex(fd, ifNameRaw)
}
