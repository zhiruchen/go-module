package common

import (
	"time"

	"go.uber.org/zap"
)

func Log() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "https://google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
