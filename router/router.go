package router

import (
	"net/http"

	"github.com/youshintop/apiserver/handler/user"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/youshintop/apiserver/handler/sd"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) {

	g.Use(gin.Recovery())

	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) { c.String(http.StatusNotFound, "The incorrect API route.") })

	pprof.Register(g)

	u := g.Group("/v1/user")
	{
		u.POST("/", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET(":id", user.Get)
		u.GET("/", user.List)
	}
	svcd := g.Group("/sd")

	svcd.GET("/health", sd.HealthCheck)
	svcd.GET("/dis", sd.DiskCheck)
}
