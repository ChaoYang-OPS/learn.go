package impl

import (
	"context"
	"database/sql"
	"fmt"
	"learn.go/testNe/rest-demo-host/apps/host"
)

// 完成对象和数据库直接的转换

// 把Host对象保存到数据内, 数据的一致性

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)
	// 把数据入库到 resource表和host表
	// 一次需要往2个表录入数据, 我们需要2个操作 要么都成功，要么都失败, 事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Println("rollback error,")
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Println("Commit error,")
			}
		}
	}()
	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}
	// 插入Describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}
	return nil
}

func (i *HostServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 开启一个事务 tx
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Printf("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Printf("commit error, %s", err)
			}
		}
	}()

	var (
		resStmt, hostStmt *sql.Stmt
	)

	// 更新 Resource表
	resStmt, err = tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return err
	}

	// 更新 Host表
	hostStmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return err
	}
	_, err = hostStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
