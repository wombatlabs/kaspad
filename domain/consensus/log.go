package consensus

import (
	"github.com/wombatlabs/kaspad/infrastructure/logger"
	"github.com/wombatlabs/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
