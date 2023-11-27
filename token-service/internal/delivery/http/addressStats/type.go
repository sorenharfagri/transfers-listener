package addressStats

import "token-service/internal/domain/addressStats"

func ProtoToAddressStatsResponse(response *addressStats.AddressStats) *AddressStatsResponse {
	return &AddressStatsResponse{
		Address:  response.Address().String(),
		Activity: response.TransferActivity(),
	}

}
