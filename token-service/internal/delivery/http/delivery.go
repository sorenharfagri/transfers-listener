package http

import (
	"fmt"
	"token-service/internal/useCase"

	"github.com/gin-gonic/gin"
)

type Delivery struct {
	ucAddressStats useCase.AddressStats
	router         *gin.Engine
}

func New(ucAddressStats useCase.AddressStats, isProduction bool, logLevel string) *Delivery {
	var d = &Delivery{
		ucAddressStats: ucAddressStats,
	}

	d.router = d.initRouter(isProduction, logLevel)
	return d
}

func (d *Delivery) Run(port uint16) error {
	return d.router.Run(fmt.Sprintf(":%d", port))
}
