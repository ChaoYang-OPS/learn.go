package impl

import (
	"database/sql"
	"learn.go/testNe/rest-demo-host/apps/host"
	"learn.go/testNe/rest-demo-host/configs"
	"log"
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

func (h *HostServiceImpl) Name() string {
	return host.AppName
}
