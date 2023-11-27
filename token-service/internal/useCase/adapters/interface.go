package adapters

import (
	"token-service/internal/domain/tokenTransfer"
)

type TransferProducer interface {
	Subscribe() <-chan *tokenTransfer.TokenTransfer
}

type TransferStorage interface {
	TransferWriter
	TransferReader
}

type TransferWriter interface {
	Insert(tt tokenTransfer.TokenTransfer)
}

type TransferReader interface {
	ListAll() []tokenTransfer.TokenTransfer
}
