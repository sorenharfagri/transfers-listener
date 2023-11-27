package tokenTransfer

import (
	"token-service/pkg/type/address"
)

type TokenTransfer struct {
	from        address.Address
	blockNumber uint64
}

func New(from address.Address, blockNumber uint64) *TokenTransfer {
	return &TokenTransfer{from: from, blockNumber: blockNumber}
}

func (tt TokenTransfer) From() address.Address {
	return tt.from
}

func (tt TokenTransfer) BlockNumber() uint64 {
	return tt.blockNumber
}
