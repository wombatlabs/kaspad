package prefixmanager

import (
	"github.com/wombatlabs/kaspad/infrastructure/logger"
	"github.com/wombatlabs/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
