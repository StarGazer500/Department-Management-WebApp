package internals

// package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// db "github.com/intdxdt/dblite"
	"github.com/StarGazer500/Department-Management-WebApp/internals/database"
	"github.com/StarGazer500/Department-Management-WebApp/internals/middlewares"
	"github.com/StarGazer500/Department-Management-WebApp/internals/routes"

	"os"
)

func NewServer() *Server {

	var useSLL bool
	var mode = gin.DebugMode
	var production = os.Getenv("PRODUCTION")
	var port = os.Getenv("PORTDEV")
	var portTLS = os.Getenv("PORTDEVTLS")
	var host = os.Getenv("HOSTDEV")

	if production == "1" {
		useSLL = true
		port = os.Getenv("PORT")
		portTLS = os.Getenv("PORTTLS")
		host = os.Getenv("HOST")
		mode = gin.ReleaseMode
	}

	var address = fmt.Sprintf(":%v", port)
	var addressTLS = fmt.Sprintf(":%v", portTLS)

	var engine = gin.New()
	engine.Use(middlewares.CorsMiddleware())
	routes.RegisterAllRoutes(engine)

	var opts = &ServerOptions{
		Host:       host,
		Address:    address,
		AddressTLS: addressTLS,
		Mode:       mode,
		UseSSL:     useSLL,
	}
	fmt.Println("port", *opts)

	var app = NewApp(engine, opts)

	var server = &Server{app: app, db: database.Dbinstance}
	return server
}

func (server *Server) Run() {
	server.app.Run()
}
