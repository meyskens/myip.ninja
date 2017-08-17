package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	cfg := getConfig()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.CORS,
	}))

	e.GET("/", handleRequest)
	if cfg.TLS {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(cfg.Hostnames...)
		e.Logger.Fatal(e.StartAutoTLS(cfg.Bind))
	} else {
		e.Logger.Fatal(e.Start(cfg.Bind))
	}
}

func handleRequest(c echo.Context) error {
	return c.String(http.StatusOK, c.RealIP())
}
