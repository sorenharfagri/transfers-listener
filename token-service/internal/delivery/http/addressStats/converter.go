package addressStats

type AddressStatsResponse struct {
	Address  string `json:"address" binding:"required"`
	Activity uint64 `json:"activity" binding:"required"`
}
