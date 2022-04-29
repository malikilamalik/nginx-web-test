package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type (
	response struct {
		Message        string         `json:"message"`
		RequestPayload RequestPayload `json:"requestPayload"`
		TimeStamp      string         `json:"timestamp"`
	}
	RequestPayload struct {
		Method   string              `json:"method"`
		Hostname string              `json:"hostname"`
		Header   map[string][]string `json:"headers"`
	}
)

var jsonOptions = context.JSON{
	Indent: "    ",
}
var port string = ":3202"
var appName string = "APP 3"

func main() {
	app := iris.New()
	app.Handle("GET", "/", commonHandler)
	app.Handle("GET", "/roundRobin", commonHandler)
	app.Handle("GET", "/ipHash", commonHandler)
	app.Handle("GET", "/leastConnection", commonHandler)
	app.Listen(port)
}

func commonHandler(ctx iris.Context) {
	t := time.Now()
	timestamp := fmt.Sprint(t.Format("02/01/2006"))

	resp := response{
		Message: fmt.Sprintf("Hello world from %s", appName),
		RequestPayload: RequestPayload{
			Method:   ctx.Request().Method,
			Hostname: ctx.Request().Host,
			Header:   ctx.Request().Header,
		},
		TimeStamp: timestamp,
	}

	ctx.JSON(resp, jsonOptions)
}
