package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/taldoflemis/brain.test/config"
	"github.com/taldoflemis/brain.test/internal/adapters/driven/auth"
	"github.com/taldoflemis/brain.test/internal/adapters/driven/misc"
	"github.com/taldoflemis/brain.test/internal/adapters/driven/postgres"
	"github.com/taldoflemis/brain.test/internal/adapters/drivers/web"
	"github.com/taldoflemis/brain.test/internal/core/services"
)

// @title		Brain.test API
// @version	0.69420
// @BasePath	/
// @host		localhost:42069
func main() {
	// Init configs
	koanf := config.NewKoanfson()
	err := koanf.LoadFromJSON("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = koanf.LoadFromEnv("BRAIN_")
	if err != nil {
		log.Fatal(err)
	}
	fiberCfg, err := config.NewFiberConfig()
	if err != nil {
		log.Fatal(err)
	}
	pgCfg, err := config.NewPostgresConfig()
	if err != nil {
		log.Fatal(err)
	}
	localIDPCfg, err := config.NewLocalIDPConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Init Drivens
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	zapLoggerAdapter := misc.NewZapLogger(sugar)

	pool, err := postgres.NewPool(postgres.GenerateConnectionString(pgCfg))
	if err != nil {
		log.Fatal(err)
	}
	_, err = pool.Exec(context.Background(), "SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	gameStorer := postgres.NewPostgresGameStorer(pool)
	localIDPStorer := postgres.NewLocalIDPPostgresStorer(pool)

	localIDP := auth.NewLocalIdp(*localIDPCfg, zapLoggerAdapter, localIDPStorer)

	// Init Services
	validationService := services.NewValidationService()
	authService := services.NewAuthenticationService(zapLoggerAdapter, localIDP, validationService)
	gameService := services.NewGameService(zapLoggerAdapter, validationService, gameStorer)

	// Init Drivers
	handlers := make([]web.Handler, 0)

	jwtMiddleware := web.NewJWTMiddleware(localIDP)
	authHandler := web.NewAuthHandler(jwtMiddleware, authService, validationService)
	handlers = append(handlers, authHandler)

	gameHandler := web.NewGameHandler(jwtMiddleware, validationService, gameService)
	handlers = append(handlers, gameHandler)

	router := web.NewRouter(*fiberCfg, logger, handlers)
	err = router.Serve()
	if err != nil {
		log.Fatal(err)
	}

	// // Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-quit
}
