package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func main() {
	m := map[string]string{"strkey1":"strval1", "strkey2":"1234", "strkey3":"5678"}
	b, err := json.Marshal(m)
	if err != nil {
		logrus.Error("json.Marshal error. err:", err)
		return
	}

	mstring := string(b)
	logrus.Info("mstring:", mstring)
}
