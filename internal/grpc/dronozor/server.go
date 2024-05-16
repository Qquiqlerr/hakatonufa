package dronozor

import (
	"context"
	dronozor2 "dronozor/protos/gen/go/obb.dronozor.v1"
	"fmt"
	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type ServerAPI struct {
	dronozor2.UnimplementedDronozorServer
	botchan chan dronozor2.PhotoRequest
}

func Register(gRPC *grpc.Server, botchan chan dronozor2.PhotoRequest) {
	dronozor2.RegisterDronozorServer(gRPC, &ServerAPI{botchan: botchan})
}

func (s *ServerAPI) SendPhoto(ctx context.Context, req *dronozor2.PhotoRequest) (*emptypb.Empty, error) {
	if err := ValidateStruct(req); err != nil {
		return nil, err
	}
	req.ImageTS = req.GetImageTS()
	s.botchan <- *req
	return &emptypb.Empty{}, nil
}

func (s *ServerAPI) SendVideo(ctx context.Context, req *dronozor2.VideoRequest) (*emptypb.Empty, error) {
	panic("Implement me")
}

func ValidateStruct(req *dronozor2.PhotoRequest) error {
	if req.GetPhone() == "" {
		return fmt.Errorf("incorrect number")
	}
	if !req.GetImageTS().IsValid() {
		return fmt.Errorf("incorrect timestamp")
	}
	coords := strings.Split(req.GetCords(), " ")
	if !govalidator.IsLatitude(coords[0]) || !govalidator.IsLongitude(coords[1]) {
		return fmt.Errorf("incorrect coordinates")
	}
	return nil
}
