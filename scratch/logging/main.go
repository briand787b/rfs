package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	sl := zap.NewExample().Sugar()
	defer sl.Sync()

	sl.Infow("logging new thing",
		"url", "https://google.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	dl, _ := zap.NewDevelopment()
	dl.Debug("logging at the debug level")

	pl, _ := zap.NewProduction()
	pl.Debug("logging at the debug level")
}
