package socketcan

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/sys/unix"
)

type ifreq struct {
	Name  [unix.IFNAMSIZ]byte
	Index int
}

func getIfIndex(itf Interface, ifName string) (int, error) {
	ifNameRaw, err := unix.ByteSliceFromString(ifName)
	if err != nil {
		return 0, err
	}
	if len(ifNameRaw) > unix.IFNAMSIZ {
		return 0, fmt.Errorf("Maximum ifname length is %d characters", unix.IFNAMSIZ)
	}

	var ifReq ifreq
	fd := itf.SocketFD()
	copy(ifReq.Name[:], ifNameRaw)
	_, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(fd),
		unix.SIOCGIFINDEX,
		uintptr(unsafe.Pointer(&ifReq)))
	if errno != 0 {
		return 0, fmt.Errorf("ioctl: %s", errno)
	}

	return ifReq.Index, nil
}

func ListCANInterfaces() ([]string, error) {
	matches, err := filepath.Glob("/sys/class/net/*")
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(matches))
	for _, m := range matches {

		typepath := filepath.Join(m, "type")
		typeAsStr, err := os.ReadFile(typepath)
		if err != nil {
			continue
		}
		typeAsInt, err := strconv.Atoi(strings.TrimSpace(string(typeAsStr)))
		if err != nil || typeAsInt != unix.ARPHRD_CAN {
			continue
		}
		res = append(res, filepath.Base(m))
	}

	return res, nil
}
