package main

import (
	"mrt-schedules/modules/stations"

	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	stations.Initiate(api)

	router.Run(":8080")
}
