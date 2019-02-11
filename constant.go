// +build !linux

package socketcan

const (
	CAN_EFF_MASK uint32 = 0x01ffffff
	CAN_SFF_MASK uint32 = 0x000007ff
	CAN_RTR_FLAG uint32 = 1 << 30
	CAN_EFF_FLAG uint32 = 1 << 31
)
