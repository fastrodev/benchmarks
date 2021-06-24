package main

import "github.com/fastrodev/fastrex"

func handler(req fastrex.Request, res fastrex.Response) {
	res.Send("hello")
}

func createApp() fastrex.App {
	app := fastrex.New()
	app.Get("/", handler)
	return app
}

func main() {
	createApp().Listen(9000)
}
