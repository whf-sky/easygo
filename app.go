package easygo

import (
	"github.com/kataras/iris/v12"
	"os"
)

var (
	app *iris.Application
	EASYGO_ENV string
)
func Run(hostPort string, withOrWithout ...iris.Configurator)  {
	app.Listen(hostPort, withOrWithout...)
}

func env()  {
	EASYGO_ENV = os.Getenv("EASYGO_ENV")
	if EASYGO_ENV == "" {
		EASYGO_ENV = "production"
	}
}

func init() {
	env()
	app = iris.New()
	app.Use(loggeeHandler())
}
