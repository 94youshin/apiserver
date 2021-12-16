package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler/sd"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) {

	g.Use(gin.Recovery())

	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) { c.String(http.StatusNotFound, "The incorrect API route.") })

	svcd := g.Group("/sd")

	svcd.GET("/health", sd.HealthCheck)
	svcd.GET("/dis", sd.DiskCheck)
}
