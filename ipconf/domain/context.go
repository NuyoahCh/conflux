package domain

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// IpConfContext Ip 配置上下文
type IpConfContext struct {
	Ctx       *context.Context
	AppCtx    *app.RequestContext
	ClientCtx *ClientContext
}

// ClientContext 客户端上下文
type ClientContext struct {
	// TODO 根据 IP 进行分类
	IP string `json:"ip"`
}

// BuildIpConfContext 构建 Ip 配置上下文
func BuildIpConfContext(c *context.Context, ctx *app.RequestContext) *IpConfContext {
	ipConfContext := &IpConfContext{
		Ctx:       c,
		AppCtx:    ctx,
		ClientCtx: &ClientContext{},
	}
	return ipConfContext
}
