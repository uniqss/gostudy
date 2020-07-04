package common

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type IProc interface {
	OnSessionConnected(session ISession)
	OnSessionClosed(session ISession)
	OnSessionMsgHandle(session ISession, msg string)
}

type IService interface {
	Register(handler IProc)
	Init(addr string) error
	Destroy()
}

type _Service struct {
	messageProc IProc
	ctx         context.Context
	cancelCtx   context.CancelFunc
}

var upgrader = websocket.Upgrader{}

func NewService() IService {
	service := &_Service{}
	service.ctx, service.cancelCtx = context.WithCancel(context.Background())
	return service
}

func (service *_Service) Register(handler IProc) {
	service.messageProc = handler
}

func (service *_Service) Init(addr string) error {
	if service.messageProc == nil {
		return errors.New("service init, no messageProc")
	}

	go func() {
		http.HandleFunc("/", service.wsHandler)
		http.ListenAndServe(addr, nil)
	}()

	log.Info("net service listening on :", addr)

	return nil
}

func (service *_Service) wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("wsHandler upgrade failed. err:", err)
		return
	}
	service.handleConn(c)
}

func (service *_Service) handleConn(conn *websocket.Conn) {
	log.Info("handleConn RemoteAddr:", conn.RemoteAddr().String())

	session := NewSession(conn)
	service.messageProc.OnSessionConnected(session)

loop:
	for {
		msg, err := session.Recv()
		if err != nil {
			log.Error("handleConn read data error. err:", err, " session:", session.GetData(), " RemoteAddr:", conn.RemoteAddr().String())
			session.Close()
			break
		}

		service.messageProc.OnSessionMsgHandle(session, msg)

		select {
		case <-service.ctx.Done():
			session.Close()
			break loop
		}
	}

	service.messageProc.OnSessionClosed(session)

	log.Info("handleConn exiting... ", session.GetData())
}

func (service *_Service) Destroy(){
	service.cancelCtx()
	log.Info("service destroying...")
}