package internals  

import (
	
	"fmt"
	"log"
	// "net/http"
	"os"
	_ "github.com/lib/pq"
    "database/sql"
	"github.com/gin-gonic/gin"
	
	
)

type Server struct {
	app *Application
	db  *sql.DB
}


type ServerOptions struct {
	Host       string
	Address    string
	AddressTLS string
	Mode       string
	UseSSL     bool
}

type Application struct {
	host        string
	address     string
	addressTLS  string
	mode        string
	engine      *gin.Engine
	useSSL      bool
}

func (opts *ServerOptions) IsDebugMode() bool {
	return opts.Mode == gin.DebugMode
}

func (opts *ServerOptions) IsSecure() bool {
	return opts.UseSSL
}



func NewApp(engine *gin.Engine, opts *ServerOptions) *Application {

	var app = &Application{
		host:        opts.Host,
		address:     opts.Address,
		addressTLS:  opts.AddressTLS,
		mode:        opts.Mode,
		engine:      engine,
		useSSL:      opts.UseSSL,

	}

	var err = app.engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	


	return app
}






func (app *Application) Run() {
	
	if app.useSSL {
		var sslCert = os.Getenv("SSLCERT")
		var sslKey = os.Getenv("SSLKEY")

		// HTTP
		go func() {
			fmt.Println("Server @  ", app.address)
			if err := app.engine.Run(app.address); err != nil {
				log.Fatalf("Error running server @ %v\n", app.address)
			}
		}()
		// HTTPS
		fmt.Println("Server @ TLS ", app.addressTLS)
		if err := app.engine.RunTLS(app.addressTLS, sslCert, sslKey); err != nil {
			log.Fatalf("Error running server @ %v\n", app.addressTLS)
		}
	} else {
		// HTTP
		fmt.Println("Server @  ", app.address)
		if err := app.engine.Run(app.address); err != nil {
			log.Fatalf("Error running server @ %v\n", app.address)
		}
	}
}