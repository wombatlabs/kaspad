package standalone

import (
	"github.com/wombatlabs/kaspad/infrastructure/logger"
	"github.com/wombatlabs/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
