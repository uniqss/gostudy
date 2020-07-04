package common

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type ISession interface {
	Send(message string) error
	Recv() (string, error)
	SetData(data interface{})
	GetData() interface{}
	Close()
	SetValidate()
	IsValidate() bool
}

type _Session struct {
	conn *websocket.Conn

	sendList ISafeList

	data       interface{}
	isValidate bool

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func NewSession(conn *websocket.Conn) ISession {
	session := &_Session{
		conn:     conn,
		sendList: NewSafeList(),
	}

	session.ctx, session.ctxCancel = context.WithCancel(context.Background())

	go session.loop()

	return session
}

func (session *_Session) Send(message string) error {
	if session.conn == nil {
		return errors.New("session Send, not connected")
	}

	if err := session.conn.WriteMessage(1, []byte(message)); err != nil {
		return err
	}
	return nil
}
func (session *_Session) Recv() (string, error) {
	_, p, err := session.conn.ReadMessage()
	//messageType, p, err := session.conn.ReadMessage()
	//log.Info("Recv messageType:", messageType)
	if err != nil {
		return "", err
	}
	return string(p), nil
}

func (session *_Session) loop() {
loop:
	for {
		select {
		case <-session.sendList.Signal():
			for {
				msg, err := session.sendList.Pop()
				if err != nil {
					break
				}
				messageString := msg.(string)
				err = session.Send(messageString)
				if err != nil {
					log.Error("session.Send error. err:", err, " RemoteAddr:", session.conn.RemoteAddr().String())
					session.close()
					break
				}
			}
		case <-session.ctx.Done():
			log.Info("session loop exit ")
			break loop
		}
	}
}

func (session *_Session) SetData(data interface{}) {
	session.data = data
}
func (session *_Session) GetData() interface{} {
	return session.data
}

func (session *_Session) Close() {
	session.close()
}

func (session *_Session) close() {
	session.ctxCancel()
	if session.conn != nil {
		_ = session.conn.Close()
	}
}

func (session *_Session) SetValidate() {
	session.isValidate = true
}

func (session *_Session) IsValidate() bool {
	return session.isValidate
}
