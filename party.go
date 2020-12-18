package easygo

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"reflect"
)

type partyStruct struct {
	party iris.Party
	ctx reflect.Value
	methods map[string]func(path string, handlers ...context.Handler) *router.Route
}

func partyHandler(ctl reflect.Value, name string) func(c iris.Context) {
	return func(c iris.Context) {
		Init := ctl.MethodByName("Init")
		Init.Call([]reflect.Value{reflect.ValueOf(c)})
		method := ctl.MethodByName(name)
		method.Call([]reflect.Value{})
	}
}

func Party(relativePath string, handlers ...context.Handler) *partyStruct {
	party := app.Party(relativePath, handlers...)
	return &partyStruct{
		party: party,
		methods: map[string]func(path string, handlers ...context.Handler) *router.Route {
			"None":party.None,
			"Get":party.Get,
			"Post":party.Post,
			"Put":party.Put,
			"Delete":party.Delete,
			"Connect":party.Connect,
			"Head":party.Head,
			"Options":party.Options,
			"Patch":party.Patch,
			"Trace":party.Trace,
		},
	}
}

func (p *partyStruct) Request(path string, context ContextInterface) *partyStruct {
	p.ctx = reflect.ValueOf(context)
	for i:=0;i<p.ctx.Type().NumMethod();i++  {
		name := p.ctx.Type().Method(i).Name
		if  fun, ok := p.methods[name]; ok {
			fun(path, partyHandler(p.ctx, name))
			continue
		}
		switch name {
		case "Any":
			p.party.Any(path, partyHandler(p.ctx, name))
		case "Use":
			p.party.Use( partyHandler(p.ctx, name))
		case "Done":
			p.party.Done( partyHandler(p.ctx, name))
		}
	}
	return p
}

func (p *partyStruct) AllowMethods(methods ...string) *partyStruct{
	p.party.AllowMethods( methods...)
	return p
}
