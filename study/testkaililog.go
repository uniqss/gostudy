package main

import (
	"errors"
	"github.com/kiali/kiali/log"
)

func main() {
	log.InitializeLogger()

	namespace := "namespace.xxxxx"
	err := errors.New("yyyyy")
	log.Errorf("Error fetching Services per namespace %s: %s", namespace, err)
}

