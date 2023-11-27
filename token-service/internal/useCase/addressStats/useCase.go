package addressStats

import (
	"token-service/internal/useCase/adapters"

	"github.com/ethereum/go-ethereum/ethclient"
)

type UseCase struct {
	adapterReader adapters.TransferReader
	client        *ethclient.Client
}

func New(adapterReader adapters.TransferReader, client *ethclient.Client) *UseCase {
	var uc = &UseCase{
		adapterReader: adapterReader,
		client:        client,
	}
	return uc
}
