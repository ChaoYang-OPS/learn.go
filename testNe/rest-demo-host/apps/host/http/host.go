package http

import (
	"learn.go/testNe/rest-demo-host/apps/host"
	"log"
	"net/http"
)

var (
	Host host.Service
)

var (
	api = &handler{}
)

type handler struct {
	service host.Service
	log     log.Logger
}

func (h *handler) Config() error {
	h.service = Host
	return nil
}

func (h *handler) QueryHost(w http.ResponseWriter)
