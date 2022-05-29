package core

import (
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
)

type ViewSelector int

const (
	DeliveredView ViewSelector = 1
	CheckedView   ViewSelector = 2
	ScreenedView  ViewSelector = 3
)

//
// TxInfo encapsulates information used by mempool to sorting.
//
type TxInfo struct {
	EffectiveGasPrice *big.Int
	Address           common.Address
	Sequence          uint64
}

//
// Ledger defines the interface of the ledger
//
type Ledger interface {
	GetCurrentBlock() *Block
	ScreenTxUnsafe(rawTx common.Bytes) result.Result
	ScreenTx(rawTx common.Bytes) (priority *TxInfo, res result.Result)
	ProposeBlockTxs(block *Block, shouldIncludeValidatorUpdateTxs bool, includeInterChainMessageTxsTillNonceMap map[InterChainMessageEventType]*big.Int) (stateRootHash common.Hash, blockRawTxs []common.Bytes, res result.Result)
	ApplyBlockTxs(block *Block) result.Result
	ApplyBlockTxsForChainCorrection(block *Block) (common.Hash, result.Result)
	//ResetState(height uint64, rootHash common.Hash) result.Result
	ResetState(block *Block) result.Result
	FinalizeState(height uint64, rootHash common.Hash) result.Result
	GetFinalizedValidatorSet(blockHash common.Hash, isNext bool) (*ValidatorSet, error)
	PruneState(endHeight uint64) error
	GetLastProcessedEventNonce(IMCEType InterChainMessageEventType, blockHash common.Hash) (*big.Int, error)
}
