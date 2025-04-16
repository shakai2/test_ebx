package main

import (
	"TakeHomeApi/pkg/router"
	"TakeHomeApi/pkg/stores"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	accountStore := stores.NewAccountStore()

	apiEngine := gin.New()
	apiEngine.HandleMethodNotAllowed = true
	apiEngine.Use(gin.Logger())
	apiEngine.Use(gin.Recovery())

	router.Router(apiEngine, accountStore)

	apiEngine.Run()
}
