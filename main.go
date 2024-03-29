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
		autoTLSManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache(cfg.CertCache),
			HostPolicy: autocert.HostWhitelist(cfg.Hostnames...),
		}

		httpServer := &http.Server{
			Handler: autoTLSManager.HTTPHandler(e),
			Addr:    ":80",
		}
		go httpServer.ListenAndServe()

		s := http.Server{
			Addr:      cfg.Bind,
			Handler:   e,
			TLSConfig: autoTLSManager.TLSConfig(),
		}
		if err := s.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
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
