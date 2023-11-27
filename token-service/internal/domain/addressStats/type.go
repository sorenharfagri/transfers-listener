package addressStats

import adr "token-service/pkg/type/address"

type AddressStats struct {
	address          adr.Address
	transferActivity uint64
}

func New(address adr.Address, transferActivity uint64) *AddressStats {
	return &AddressStats{address: address, transferActivity: transferActivity}
}

func (as AddressStats) Address() adr.Address {
	return as.address
}

func (as AddressStats) TransferActivity() uint64 {
	return as.transferActivity
}
