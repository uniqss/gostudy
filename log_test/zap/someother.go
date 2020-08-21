package main

import (
	"go.uber.org/zap"
	"zap/common"
)

func SomeOther() {
	zLog := common.ZLog
	zLog.Info("This is log in SomeOther func", zap.String("somekey", "somevalue"))
}
