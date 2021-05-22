package main

import (
	"github.com/fastrodev/router"
	"github.com/fastrodev/router/web"
)

func main() {
	port := 9000
	app := router.New()
	app.Get("/", func(req web.Request, res web.Response) {
		res.Send("hello")
	})
	app.Listen(port, func() {
		print("Web app listening at http://localhost:", port)
	})
}
