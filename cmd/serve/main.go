package main

import "github.com/imranismail/ecommerce/app"

func main() {
	app := app.New()
	app.Serve()
}
