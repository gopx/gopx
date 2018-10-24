package main

import (
	golog "log"
	"net"
	"net/http"
	"os"
	"time"

	"gopx.io/gopx/internal/app/gopx/constants"
	"gopx.io/gopx/internal/app/gopx/helper"
	"gopx.io/gopx/internal/app/gopx/route"
	"gopx.io/gopx/pkg/log"
)

var serverLogger = golog.New(os.Stdout, "", golog.Ldate|golog.Ltime|golog.Lshortfile)

func main() {
	startServer()
}

func startServer() {
	host := helper.MustGetEnv(constants.EnvNameHost)
	port := helper.MustGetEnv(constants.EnvNamePort)
	addr := net.JoinHostPort(host, port)
	r := route.Router()
	server := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       constants.ServerReadTimeout * time.Second,
		ReadHeaderTimeout: constants.ServerReadTimeout * time.Second,
		WriteTimeout:      constants.ServerWriteTimeout * time.Second,
		IdleTimeout:       constants.ServerIdleTimeout * time.Second,
		ErrorLog:          serverLogger,
	}

	log.Info("GoPx service is running on: %s", addr)
	err := server.ListenAndServe()
	log.Fatal("Error: %s", err) // err is always non-nill
}
