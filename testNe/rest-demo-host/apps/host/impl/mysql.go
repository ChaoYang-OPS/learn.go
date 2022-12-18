package impl

import (
	"context"
	"database/sql"
	"github.com/infraboard/mcube/sqlbuilder"
	"learn.go/testNe/rest-demo-host/apps"
	"learn.go/testNe/rest-demo-host/apps/host"
	"learn.go/testNe/rest-demo-host/configs"
	"log"
)

var _ host.Service = &HostServiceImpl{}

type HostServiceImpl struct {
	db *sql.DB
	l  *log.Logger
}

var impl = HostServiceImpl{}

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		db: configs.C().MySQL.GetDB(),
		l:  log.Default(),
	}
}
func (h *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	// 默认值填充
	ins.InjectDefault()
	// 有dao模块 负责 把对象入库
	if err := h.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (h *HostServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (*host.HostSet, error) {
	b := sqlbuilder.NewBuilder(QueryHostSQL)
	if request.Keywords != "" {
		// (r.`name`='%' OR r.description='%' OR r.private_ip='%' OR r.public_ip='%')
		//  10.10.1, 接口测试
		b.Where("r.`name`LIKE ? OR r.description LIKE ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+request.Keywords+"%",
			"%"+request.Keywords+"%",
			request.Keywords+"%",
			request.Keywords+"%",
		)
	}
	b.Limit(request.OffSet(), request.GetPageSize())

	querySQL, args := b.Build()
	h.l.Printf("query sql %s, args: %v\n", querySQL, args)
	// query stmt, 构建一个Prepare语句
	stmt, err := h.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	set := host.NewHostSet()
	for rows.Next() {
		ins := host.NewHost()
		if err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
			&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
			&ins.Account, &ins.PublicIP, &ins.PrivateIP,
			&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)
	}
	// total统计
	countSQL, args := b.BuildCount()
	h.l.Printf("count sql: %s, args %v\n", countSQL, args)
	countStmt, err := h.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}
	return set, nil
}

func (h *HostServiceImpl) DescribeHost(ctx context.Context, request *host.DescribeHostRequest) (*host.Host, error) {
	//TODO implement me
	return nil, nil
}

func (h *HostServiceImpl) UpdateHost(ctx context.Context, request *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (h *HostServiceImpl) DeleteHost(ctx context.Context, request *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}

func (h *HostServiceImpl) Config() {
	db := configs.C().MySQL.GetDB()
	h.db = db
}

func (h *HostServiceImpl) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryImpl(&impl)
}
