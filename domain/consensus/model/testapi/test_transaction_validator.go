package testapi

import (
	"github.com/wombatlabs/kaspad/domain/consensus/model"
	"github.com/wombatlabs/kaspad/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
