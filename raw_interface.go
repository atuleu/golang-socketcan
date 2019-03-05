package socketcan

type RawInterface interface {
	Send(CanFrame) error
	Receive() (CanFrame, error)
	Close() error
}
