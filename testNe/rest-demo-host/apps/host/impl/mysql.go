package impl

import (
	"context"
	"database/sql"
	"fmt"
	"learn.go/testNe/rest-demo-host/apps/host"
	"learn.go/testNe/rest-demo-host/configs"
	"log"
)

const (
	InsertResourceSQL = `
	INSERT INTO resource (
		id,
		vendor,
		region,
		create_at,
		expire_at,
		type,
		name,
		description,
		status,
		update_at,
		sync_at,
		accout,
		public_ip,
		private_ip
	)
	VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`

	// INSERT INTO `host` ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	// VALUES
	// ( "111", 1, 2048, 1, 'n', 'linux', 'centos8', '00000' );
	InsertDescribeSQL = `
	INSERT INTO host ( resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number )
	VALUES
		( ?,?,?,?,?,?,?,? );
	`

	queryHostSQL = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
)

var _ host.Service = &HostServiceImpl{}

type HostServiceImpl struct {
	db *sql.DB
	l  log.Logger
}

func (h *HostServiceImpl) Config() error {
	db, err := configs.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	h.db = db
	return nil
}

func (h *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				h.l.Println("rollback error ,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				h.l.Println("commit error ,%s", err)
			}
		}
	}()

	rstmt ,err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_ err := rstmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
		)
	if err != nil {
		return err
	}
	dstmt, err := tx.Prepare(InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}

	return nil

}
