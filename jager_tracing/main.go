package main

import (
	"fmt"
	"jager_tracing/tracing_wrapper"
	"reflect"
	"time"
)

func doSomeOperation(operation string) {
	fmt.Println(operation)
}

func traceGateway(msgTracingValue reflect.Value) {
	operation := "Gateway msg operation"
	span, err := tracing_wrapper.TracingRootSpan(msgTracingValue, "root", "msgName", operation)
	if err != nil {
		fmt.Println("traceGateway tracing_wrapper.TracingRootSpan error. err:", err)
	}

	// some important operation
	doSomeOperation(operation)

	tracing_wrapper.TracingEndSpan(span)
}

func traceGameServer(msgTracingValue reflect.Value) {
	operation := "GameServer msg operation"
	span, err := tracing_wrapper.TracingChildSpan(msgTracingValue, "root", "msgName", operation)
	if err != nil {
		fmt.Println("traceGameServer tracing_wrapper.TracingChildSpan error. err:", err)
	}

	// some important operation
	doSomeOperation(operation)

	tracing_wrapper.TracingEndSpan(span)
}

type ClientValidateReq struct {
	Version int32  `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	ID      string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	MsgSNo  uint32 `protobuf:"varint,4,opt,name=MsgSNo,proto3" json:"MsgSNo,omitempty"`
	Tracing []byte `protobuf:"bytes,5,opt,name=Tracing,proto3" json:"Tracing,omitempty"`
}
func (m *ClientValidateReq) Reset()         { *m = ClientValidateReq{} }

func main() {
	tracing_wrapper.InitJaeger("jager_tracing_test")

	// suppose this is some message that should be traced
	msg := &ClientValidateReq{}

	val := reflect.ValueOf(msg).Elem().FieldByName("Tracing")
	traceGateway(val)

	traceGameServer(val)

	time.Sleep(time.Second * 3)

	tracing_wrapper.CloseJaeger()
}
