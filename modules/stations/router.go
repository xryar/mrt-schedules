package stations

import (
	"mrt-schedules/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(ctx *gin.Context) {
		//code service
		GetAllStation(ctx, stationService)
	})

	station.GET("/:id", func(ctx *gin.Context) {
		CheckSchedulesByStation(ctx, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		//handle error
		c.JSON(http.StatusBadRequest,
			response.ApiResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}

	//balikin response
	c.JSON(
		http.StatusOK,
		response.ApiResponse{
			Success: true,
			Message: "Successfully get all station",
			Data:    datas,
		},
	)
}

func CheckSchedulesByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScheduleByStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ApiResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.ApiResponse{
			Success: true,
			Message: "Successfully get schedules by station",
			Data:    datas,
		},
	)
}
