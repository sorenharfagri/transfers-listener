package addressStats

import (
	"sort"
	"token-service/internal/domain/addressStats"
	"token-service/pkg/type/address"
	"token-service/pkg/type/context"
)

// Логику query фильтрования можно перенести на репозиторий
func (uc *UseCase) TopFive(ctx context.Context) ([]addressStats.AddressStats, error) {
	transfers := uc.adapterReader.ListAll()

	// Можно вынести в отдельный сервис и кешировать
	lastBlockNumber, err := uc.client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	addressStatsMap := make(map[address.Address]uint64)
	minBlockNumber := uint64(lastBlockNumber - 100)

	for _, transfer := range transfers {
		if transfer.BlockNumber() >= minBlockNumber {
			addressStatsMap[transfer.From()]++
		}
	}

	var sortedStats []addressStats.AddressStats
	for address, count := range addressStatsMap {
		sortedStats = append(sortedStats, *addressStats.New(address, count))
	}

	sort.Slice(sortedStats, func(i, j int) bool {
		return sortedStats[i].TransferActivity() > sortedStats[j].TransferActivity()
	})

	var topStats []addressStats.AddressStats
	for i := 0; i < 5 && i < len(sortedStats); i++ {
		topStats = append(topStats, sortedStats[i])
	}

	// Можно выдавать ошибку если нет данных/недостаточно трансферов для выдачи топ 5
	return topStats, nil
}
