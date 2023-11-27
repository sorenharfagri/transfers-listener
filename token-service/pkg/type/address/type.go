package address

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

type Address struct {
	value string
}

func New(adr string) (*Address, error) {
	isValid := isValidAddress(adr)
	if !isValid {
		return nil, errors.New("invalid address")
	}

	return &Address{value: adr}, nil
}

func (a Address) String() string {
	return a.value
}

func isValidAddress(adr string) bool {
	address := common.HexToAddress(adr)
	return address.Hex() == adr
}
