package common

// interfaces
type IProc interface {
	MsgProc(srcAddr string, message string)
}
type INatsService interface {
	Init(options *NatsConfig) error
	AddProc(proc IProc)
	Destroy()
	Post(addr string, message string) error
}

func NewNatsService() INatsService {
	return &_Service{}
}
