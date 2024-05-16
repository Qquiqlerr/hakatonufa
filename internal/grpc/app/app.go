package app

import (
	"dronozor/internal/grpc/dronozor"
	dronozor2 "dronozor/protos/gen/go/obb.dronozor.v1"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type App struct {
	gRPC *grpc.Server
	port int
}

func New(port int, botchan chan dronozor2.PhotoRequest) *App {
	grpcServer := grpc.NewServer()
	dronozor.Register(grpcServer, botchan)
	return &App{gRPC: grpcServer, port: port}
}

func (a *App) MustStart() {
	fmt.Println("starting server")
	if err := a.Start(); err != nil {
		panic(err)
	}
}

func (a *App) Start() error {
	const op = "app.Start"
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(a.port))
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if err := a.gRPC.Serve(lis); err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	a.gRPC.GracefulStop()
}
