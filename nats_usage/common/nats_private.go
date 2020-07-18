package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

// privates
type _Service struct {
	options    *NatsConfig
	arrProc    []IProc
	nc         *nats.Conn
	sendList   ISafeList
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (service *_Service) Init(options *NatsConfig) error {
	service.options = options
	service.sendList = NewSafeList()
	service.ctx, service.cancelFunc = context.WithCancel(context.Background())
	if err := service.initProducer(service.options.NatsAddr); err != nil {
		return err
	}
	service.ping()
	service.nc.Subscribe(options.ServiceAddr, service.HandleMessage)
	service.nc.Subscribe(options.BroadcastAddr, service.HandleMessage)

	go service.sendLoop()

	return nil
}

func (service *_Service) Destroy() {
	service.cancelFunc()

	if service.nc != nil {
		service.nc.Close()
		service.nc = nil
	}
}
func (service *_Service) AddProc(proc IProc) {
	service.arrProc = append(service.arrProc, proc)
}

var ErrorNotInit = errors.New("Nats not init")

type _SendInfo struct {
	Addr    string
	Message string
}

func (service *_Service) Post(addr string, message string) error {
	if service.nc == nil {
		return ErrorNotInit
	}

	if message == "" {
		return nil
	}

	service.sendList.Push(&_SendInfo{addr, message})

	return nil
}
func (service *_Service) initProducer(addr string) error {
	var err error
	service.nc, err = nats.Connect(addr, nats.Timeout(10*time.Second))
	if err != nil {
		return err
	}
	return nil
}
func (service *_Service) ping() {
	err := service.Post(service.options.ServiceAddr, "ping from:"+service.options.ServiceAddr)
	if err != nil {
		fmt.Println("ping failed. err:", err)
	}
	err = service.Post(service.options.BroadcastAddr, "broadcast ping from:"+service.options.BroadcastAddr)
	if err != nil {
		fmt.Println("broad cast ping failed. err:", err)
	}
}
func (service *_Service) HandleMessage(msg *nats.Msg) {
	var data []byte
	data = msg.Data

}
