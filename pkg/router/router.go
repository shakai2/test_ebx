package router

import (
	"TakeHomeApi/pkg/routes"
	"TakeHomeApi/pkg/stores"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, store *stores.AccountStore) {
	router.GET("/test", routes.Test())
	router.GET("/balance", routes.Balance(store))
	router.POST("/reset", routes.Reset(store))
	router.POST("/event", routes.Event(store))

	router.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	})
}
