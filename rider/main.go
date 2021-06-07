package main

import "github.com/fastrodev/rider"

func handler(req rider.Request, res rider.Response) {
	res.Send("hello")
}

func createApp() rider.Router {
	app := rider.New()
	app.Get("/", handler)
	return app
}

func main() {
	createApp().Listen(9000)
}
