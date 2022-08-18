package main

import (
	"github.com/Jac0bDeal/sugoku/internal/cli"
	"go.uber.org/zap"
	"log"
)

var (
	Version = "dev"
)

func setupLogger(level string) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()

	var err error
	config.Level, err = zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	return config.Build()
}

func main() {
	app := cli.NewApp(Version)

	logger, err := setupLogger(app.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			log.Fatalf("failed to sync logger: %s", err.Error())
		}
	}(logger)
	zap.ReplaceGlobals(logger)

	if err := app.Run(); err != nil {
		zap.S().Fatal(err)
	}
}
