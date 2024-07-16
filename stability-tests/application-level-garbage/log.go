package main

import (
	"github.com/wombatlabs/kaspad/infrastructure/logger"
	"github.com/wombatlabs/kaspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
