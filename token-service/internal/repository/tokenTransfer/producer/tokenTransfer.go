package events

import (
	"context"
	"errors"
	log "token-service/pkg/type/logger"

	"token-service/internal/domain/tokenTransfer"
	"token-service/pkg/type/address"

	"github.com/ethereum/go-ethereum"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (p *Producer) initialize() error {
	ctx := context.Background()

	query := ethereum.FilterQuery{ // Хеш сигнатуры transfer
		Topics: [][]ethCommon.Hash{{ethCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}

	transfersCh := make(chan *tokenTransfer.TokenTransfer)
	logsCh := make(chan types.Log)

	sub, err := p.client.SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logsCh:

				from, blockNum, err := decodeLog(vLog)

				if err != nil {
					continue
				}

				// Исключаем генерацию
				if from.String() == "0x0000000000000000000000000000000000000000" {
					continue
				}

				tokenTransfer, err := logToDomain(*from, *blockNum)

				if err != nil {
					continue
				}

				transfersCh <- tokenTransfer
			}
		}
	}()

	go func() {
		for event := range transfersCh {
			p.mutex.Lock()
			for eventChan := range p.subClients {
				eventChan <- event
			}
			p.mutex.Unlock()
		}
	}()

	log.Info("Listening for transfer events...")

	return nil
}

func decodeLog(log types.Log) (*ethCommon.Address, *uint64, error) {

	if len(log.Topics) != 3 {
		return nil, nil, errors.New("invalid topics length")
	}

	from := ethCommon.HexToAddress(log.Topics[1].Hex())
	blockNum := log.BlockNumber

	return &from, &blockNum, nil
}

func logToDomain(from ethCommon.Address, blockNum uint64) (*tokenTransfer.TokenTransfer, error) {

	fromD, err := address.New(from.String())

	if err != nil {
		return nil, err
	}

	tokenTransfer := tokenTransfer.New(*fromD, blockNum)

	return tokenTransfer, nil
}

func (p *Producer) Subscribe() <-chan *tokenTransfer.TokenTransfer {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	transfersCh := make(chan *tokenTransfer.TokenTransfer)
	p.subClients[transfersCh] = true

	return transfersCh
}

// TODO
// func (p *Producer) Unsubscribe(eventChan chan *tokenTransfer.TokenTransfer) {
// }
