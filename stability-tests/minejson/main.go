package main

import (
	"github.com/wombatlabs/kaspad/domain/consensus"
	"github.com/wombatlabs/kaspad/stability-tests/common"
	"github.com/wombatlabs/kaspad/stability-tests/common/mine"
	"github.com/wombatlabs/kaspad/stability-tests/common/rpc"
	"github.com/wombatlabs/kaspad/util/panics"
	"github.com/wombatlabs/kaspad/util/profiling"
	"github.com/pkg/errors"
)

func main() {
	defer panics.HandlePanic(log, "minejson-main", nil)
	err := parseConfig()
	if err != nil {
		panic(errors.Wrap(err, "error parsing configuration"))
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())

	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}
	rpcClient, err := rpc.ConnectToRPC(&cfg.Config, cfg.NetParams())
	if err != nil {
		panic(errors.Wrap(err, "error connecting to JSON-RPC server"))
	}
	defer rpcClient.Disconnect()

	dataDir, err := common.TempDir("minejson")
	if err != nil {
		panic(err)
	}

	consensusConfig := consensus.Config{Params: *cfg.NetParams()}

	err = mine.FromFile(cfg.DAGFile, &consensusConfig, rpcClient, dataDir)
	if err != nil {
		panic(errors.Wrap(err, "error in mine.FromFile"))
	}
}
