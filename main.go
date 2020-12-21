package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	app.HandleDir("/", "./views")

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
