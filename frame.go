package socketcan

import (
	"encoding/binary"

	"golang.org/x/sys/unix"
)

type CanFrame struct {
	ID       uint32
	Dlc      byte
	Data     []byte
	Extended bool
	RTR      bool
}

func (f CanFrame) putID(buf []byte) {
	if f.Extended == true {
		f.ID = f.ID & unix.CAN_EFF_MASK
	} else {
		f.ID = f.ID & unix.CAN_SFF_MASK
	}
	if f.RTR {
		f.ID |= unix.CAN_RTR_FLAG
	}

	binary.LittleEndian.PutUint32(buf[0:4], f.ID)
}

func (f CanFrame) getID(buf []byte) {
	f.ID = uint32(binary.LittleEndian.Uint32(buf[0:4]))

	if f.ID&unix.CAN_RTR_FLAG != 0 {
		f.RTR = true
	} else {
		f.RTR = false
	}

	if f.ID&unix.CAN_EFF_FLAG != 0 {
		f.ID &= unix.CAN_EFF_MASK
		f.Extended = true
	} else {
		f.ID &= unix.CAN_SFF_MASK
		f.Extended = false
	}
}
