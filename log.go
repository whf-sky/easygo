package easygo

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"time"
)

func loggeeHandler() context.Handler {
	return logger.New(logger.Config{
		Status:     true,
		IP:         true,
		Method:     true,
		Path:       true,
		Query:      false,
		Columns:    false,
		LogFunc:    logFunc,
		LogFuncCtx: nil,
		Skippers:   nil,
	})
}

func logFunc(endTime time.Time,
	latency time.Duration,
	status, ip, method, path string,
	message, headerMessage interface{}){
	fmt.Println("[iris] ",endTime, "|",status,"|",latency,"|",ip,"|",method,"|",path)
}