package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := getConfig()

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", handleRequest)
	if cfg.TLS {
		e.AutoTLSManager.Cache = autocert.DirCache(cfg.CertCache)
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(cfg.Hostnames...)
		e.Logger.Fatal(e.StartAutoTLS(cfg.Bind))
	} else {
		e.Logger.Fatal(e.Start(cfg.Bind))
	}
}

//IP is just a tring but called IP for XML conversion
type IP string

func handleRequest(c echo.Context) error {
	if c.QueryParam("format") == "json" {
		return c.JSON(http.StatusOK, map[string]string{"ip": c.RealIP()})
	}

	if c.QueryParam("format") == "jsonp" && c.QueryParam("callback") != "" {
		return c.JSONP(http.StatusOK, c.QueryParam("callback"), map[string]string{"ip": c.RealIP()})
	}

	if c.QueryParam("format") == "xml" {
		return c.XML(http.StatusOK, IP(c.RealIP()))
	}

	return c.String(http.StatusOK, c.RealIP())
}
