package echoroutes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/libs/appctx"
)

func SetupRoutes(ctx appctx.AppContext, r *echo.Echo) {
	r.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
}
