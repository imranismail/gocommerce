package main

import "github.com/imranismail/ecommerce/internal/endpoint"

func main() {
	app := endpoint.New()
	app.Serve()
}
