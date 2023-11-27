package transferHandler

import (
	"token-service/internal/useCase/adapters"
)

type UseCase struct {
	adapterProducer adapters.TransferProducer
	adapterStorage  adapters.TransferStorage
}

func New(adapterProducer adapters.TransferProducer, adapterStorage adapters.TransferStorage) *UseCase {
	var uc = &UseCase{
		adapterProducer: adapterProducer,
		adapterStorage:  adapterStorage,
	}
	uc.initialize()
	return uc
}
