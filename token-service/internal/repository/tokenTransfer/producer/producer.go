package events

import (
	"sync"
	"token-service/internal/domain/tokenTransfer"
	log "token-service/pkg/type/logger"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Producer struct {
	subClients map[chan *tokenTransfer.TokenTransfer]bool
	mutex      sync.Mutex
	client     *ethclient.Client
}

func New(client *ethclient.Client) (*Producer, error) {

	p := &Producer{
		client:     client,
		subClients: make(map[chan *tokenTransfer.TokenTransfer]bool),
	}

	err := p.initialize()

	if err != nil {
		log.Fatal("r-tokenTransfer-producer-intialize", zap.Error(err))
	}

	return p, nil
}
