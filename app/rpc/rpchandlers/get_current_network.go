package rpchandlers

import (
	"github.com/wombatlabs/kaspad/app/appmessage"
	"github.com/wombatlabs/kaspad/app/rpc/rpccontext"
	"github.com/wombatlabs/kaspad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
