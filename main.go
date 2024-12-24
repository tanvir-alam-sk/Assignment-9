package main

import (
	_ "example-beego/routers"

	"github.com/beego/beego/v2/server/web/context"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	// Add CORS support
	beego.InsertFilter("*", beego.BeforeRouter, corsMiddleware)

	beego.Run()
}

func corsMiddleware(ctx *context.Context) {
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*") // Replace '*' with specific origins if needed
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight OPTIONS request
	if ctx.Input.Method() == "OPTIONS" {
		ctx.ResponseWriter.WriteHeader(204) // No Content
	}
}