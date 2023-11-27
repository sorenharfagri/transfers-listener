package http

import (
	"net/http"
	jsonAddressStats "token-service/internal/delivery/http/addressStats"
	"token-service/pkg/type/context"

	"github.com/gin-gonic/gin"
)

func (d *Delivery) TopFive(c *gin.Context) {

	var ctx = context.New(c)

	addressStats, err := d.ucAddressStats.TopFive(ctx)

	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var list = []*jsonAddressStats.AddressStatsResponse{}
	for _, value := range addressStats {
		list = append(list, jsonAddressStats.ProtoToAddressStatsResponse(&value))
	}

	c.JSON(http.StatusOK, list)
}
