package http

import (
	"github.com/gin-gonic/gin"
	"learn.go/testNe/rest-demo-host/apps/host"
	"log"
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

func (h *handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		c.JSON(401, "error")
		return
	}
	ins, err := h.service.CreateHost(c.Request.Context(), ins)
	if err != nil {
		c.JSON(400, "data error")
		return
	}
	c.JSON(200, ins)
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost())
}

func init() {
	
}
