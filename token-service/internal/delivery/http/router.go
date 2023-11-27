package http

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"token-service/pkg/type/logger"
)

func (d *Delivery) initRouter(
	isProduction bool,
	logLevel string,
) *gin.Engine {

	if isProduction {
		switch logLevel {
		case "DEBUG":
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	var router = gin.New()

	router.Use(Tracer())
	// Logs all panic to error log
	// stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(logger.GetLogger(), true))

	router.GET("/top-five", d.TopFive)

	return router
}
