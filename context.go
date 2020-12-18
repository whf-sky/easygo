package easygo

import (
	"github.com/kataras/iris/v12"
)


type Context struct {
	Ctx  iris.Context

}

type ContextInterface interface {
	Init(ctx  iris.Context)
}

func (c *Context) Init(ctx  iris.Context) {
	c.Ctx = ctx
}
