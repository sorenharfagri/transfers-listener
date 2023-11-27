package transferHandler

import (
	"errors"
	log "token-service/pkg/type/logger"

	"go.uber.org/zap"
)

func (uc *UseCase) initialize() {
	go func() {

		transferCh := uc.adapterProducer.Subscribe()

		for transfer := range transferCh {
			uc.adapterStorage.Insert(*transfer)
		}

		log.Fatal("uc-transferHandler-intiialize", zap.Error(errors.New("transfers chan was closed")))
	}()
}
