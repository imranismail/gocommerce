package main

import "github.com/imranismail/ecommerce/internal/endpoint"

func main() {
	e := endpoint.New()
	e.Serve()
}
