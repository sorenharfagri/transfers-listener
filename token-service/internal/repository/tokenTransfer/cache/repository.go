package cache

import "token-service/internal/domain/tokenTransfer"

type Repository struct {
	transfers []tokenTransfer.TokenTransfer
}

func New() *Repository {
	return &Repository{
		transfers: make([]tokenTransfer.TokenTransfer, 0),
	}
}
