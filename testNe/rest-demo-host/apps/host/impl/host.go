package impl

import (
	"context"
	"learn.go/testNe/rest-demo-host/apps/host"
)

var _ host.Service = &HostServiceImpl{}

//func (i *HostServiceImpl) Config() error {
//	return nil
//}

func (h *HostServiceImpl) SaveHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	return nil, nil
}
