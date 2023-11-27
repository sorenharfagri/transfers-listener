package useCase

import (
	"token-service/internal/domain/addressStats"
	"token-service/pkg/type/context"
)

type AddressStats interface {
	TopFive(context.Context) ([]addressStats.AddressStats, error)
}
