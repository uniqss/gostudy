package common

import (
	log "github.com/sirupsen/logrus"
)

type _TestMsgProc struct {
}

func NewTestMsgProc() IProc{
	mp := &_TestMsgProc{}
	return mp
}

func (proc *_TestMsgProc) OnSessionConnected(session ISession){
	log.Info("OnSessionConnected")
}

func (proc *_TestMsgProc) OnSessionClosed(session ISession){
	log.Info("OnSessionClosed")
}

func (proc *_TestMsgProc) OnSessionMsgHandle(session ISession, msg string){
	log.Info("OnSessionMsgHandle msg:", msg)
	err := session.Send(msg)
	if err != nil {
		log.Error("OnSessionMsgHandle session.Send failed.err:", err)
	}
}

