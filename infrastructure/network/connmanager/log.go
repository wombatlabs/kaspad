package connmanager

import (
	"github.com/wombatlabs/kaspad/infrastructure/logger"
	"github.com/wombatlabs/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
