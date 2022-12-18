package http

import (
	"github.com/gin-gonic/gin"
	"learn.go/testNe/rest-demo-host/apps"
	"learn.go/testNe/rest-demo-host/apps/host"
)

// 面向接口, 真正Service的实现, 在服务实例化的时候传递进行
// 也就是(CLI)  Start时候
var handler = &Handler{}

type Handler struct {
	service host.Service
}

func (h *Handler) Config() {
	h.service = apps.GetImpl(host.AppName).(host.Service)
}

func (h *Handler) Name() string {
	return host.AppName
}

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		c.JSON(401, "error")
		return
	}
	// mock data save
	//ins, err := h.service.CreateHost(c.Request.Context(), ins)
	//if err != nil {
	//	c.JSON(400, err)
	//	return
	//}
	// add default value
	ins.InjectDefault()
	c.JSON(200, ins)
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func init() {
	apps.RegistryGin(handler)
}
