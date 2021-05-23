package main

import (
	"github.com/fastrodev/router"
	"github.com/fastrodev/router/http"
)

func main() {
	port := 9000
	app := router.New()
	app.Get("/", func(req http.Request, res http.Response) {
		res.Send("hello")
	})
	app.Listen(port, func() {
		print("Web app listening at http://localhost:", port)
	})
}
