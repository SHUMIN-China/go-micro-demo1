package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-micro-demo1/user-service/API/handler"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	g.GET("/addtime/:name", handler.GetTimeOrAddUser)
	return g
}
