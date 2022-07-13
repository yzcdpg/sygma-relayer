package events

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (es EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

const (
	DepositSig           EventSig = "Deposit(uint8,bytes32,uint64,address,bytes,bytes)"
	StartKeygenSig       EventSig = "StartKeygen()"
	KeyRefreshSig        EventSig = "KeyRefresh(string)"
	ProposalExecutionSig EventSig = "ProposalExecution(uint8,uint64,bytes32)"
	FeeChangedSig        EventSig = "FeeChanged(uint256)"
)

// Refresh struct holds key refresh event data
type Refresh struct {
	// SHA1 hash of topology file
	Hash string
}
