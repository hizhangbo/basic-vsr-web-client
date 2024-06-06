package service

import (
	"context"

	v1 "basic-vsr-web-client/api/basicvsr/v1"
	"basic-vsr-web-client/internal/biz"
)

// BasicVSRService basic vsr client service
type BasicVSRService struct {
	v1.UnimplementedBasicVSRServer

	uc *biz.BasicVSRUsecase
}

// NewGreeterService new a greeter service.
func NewBasicVSRService(uc *biz.BasicVSRUsecase) *BasicVSRService {
	return &BasicVSRService{uc: uc}
}

// GetStatus implements helloworld.GreeterServer.
func (s *BasicVSRService) GetStatus(ctx context.Context, in *v1.GPURequest) (*v1.GPUReply, error) {
	g, err := s.uc.GetStatus(ctx, &biz.BasicVSR{ProductName: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.GPUReply{
		ProductName: g.ProductName,
		Power:       g.PowerReadings.PowerDraw,
		Fan:         g.FanSpeed,
		Memory:      g.FbMemoryUsage.Used,
	}, nil
}

func (s *BasicVSRService) ExecBasicVsr(ctx context.Context, in *v1.GPURequest) (*v1.ExecReply, error) {
	vsr, err := s.uc.ExecBasicVsr(ctx, &biz.GPURequest{Name: in.Name})
	if err != nil {
		return nil, err
	}

	return &v1.ExecReply{Message: vsr.Name}, nil
}
