package host

import "context"

type Service interface {
	CreateHost(ctx context.Context, host *Host) (*Host, error)
	QueryHost(ctx context.Context, request *QueryHostRequest) (*HostSet, error)
	DescribeHost(ctx context.Context, request *DescribeHostRequest) (*Host, error)
	UpdateHost(ctx context.Context, request *UpdateHostRequest) (*Host, error)
	DeleteHost(ctx context.Context, request *DeleteHostRequest) (*Host, error)
}
