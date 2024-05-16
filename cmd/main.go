package main

import (
	"dronozor/internal/bot"
	"dronozor/internal/config"
	"dronozor/internal/grpc/app"
	dronozor2 "dronozor/protos/gen/go/obb.dronozor.v1"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	botchan := make(chan dronozor2.PhotoRequest, 1)
	application := app.New(cfg.Port, botchan)
	go bot.StartBot(botchan)
	go application.MustStart()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	application.Stop()
	fmt.Println("Application stopped")
}
