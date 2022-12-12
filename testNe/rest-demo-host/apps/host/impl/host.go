package impl

import (
	"context"
	"learn.go/testNe/rest-demo-host/apps/host"
)

var _ host.Service = &HostServiceImpl{}

//func (i *HostServiceImpl) Config() error {
//	return nil
//}

func (h *HostServiceImpl) SaveHost(ctx context.Context, host *host.Host) (*Host, error) {
	return nil, nil
}
func (h *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*HostSet, error) {
	return nil, nil
}
