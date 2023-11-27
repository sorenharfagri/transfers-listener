package cache

import "token-service/internal/domain/tokenTransfer"

// Можно добавить доп логику кеширования, по типу хранения transfer только за последние 100 блоков
// Динмаически фильтровать от большего к меньшему transfer при добавлении нового и тд
func (r *Repository) Insert(transfer tokenTransfer.TokenTransfer) {
	r.transfers = append(r.transfers, transfer)
}

func (r *Repository) ListAll() []tokenTransfer.TokenTransfer {
	return r.transfers
}
