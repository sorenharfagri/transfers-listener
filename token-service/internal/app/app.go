package app

import (
	"os"
	"os/signal"
	"syscall"
	"token-service/config"

	deliveryHttp "token-service/internal/delivery/http"
	transferStorage "token-service/internal/repository/tokenTransfer/cache"
	transferProducer "token-service/internal/repository/tokenTransfer/producer"
	useCaseAddressStats "token-service/internal/useCase/addressStats"
	useCaseTransferHandler "token-service/internal/useCase/transferHandler"
	"token-service/pkg/tracing"
	"token-service/pkg/type/context"

	log "token-service/pkg/type/logger"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	log.Init(*cfg.Service.IsProduction, cfg.Log.Level)
	logger, err := log.New()
	if err != nil {
		panic(err)
	}

	tracingOpts := tracing.Options{
		JaegerHost:  cfg.Jaeger.Host,
		JaegerPort:  cfg.Jaeger.Port,
		ServiceName: cfg.Service.Name,
	}

	closer, err := tracing.New(context.Empty(), tracingOpts)
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			logger.Error(err)
		}
	}()

	handler := gin.New()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	handler.Use(cors.New(corsCfg))

	client, err := ethclient.Dial(cfg.Provider.Url)
	if err != nil {
		log.Fatal(err)
	}

	transferProducer, err := transferProducer.New(client)

	if err != nil {
		log.Fatal(err)
	}

	transferStorage := transferStorage.New()
	_ = useCaseTransferHandler.New(transferProducer, transferStorage)
	addressStatsUc := useCaseAddressStats.New(transferStorage, client)

	listenerHttp := deliveryHttp.New(addressStatsUc, *cfg.Service.IsProduction, cfg.Log.Level)

	go func() {
		log.Info("service started successfully on http", zap.Uint16("port", cfg.HTTP.Port))
		if err = listenerHttp.Run(cfg.HTTP.Port); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
