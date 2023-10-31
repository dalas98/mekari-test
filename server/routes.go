package server

import (
	"github.com/dalas98/mekari-test/app/handlers"
	"github.com/dalas98/mekari-test/app/usecases"
	"github.com/gin-gonic/gin"
)

var baseUrl = "/api"

func APIGORoutes(r *gin.Engine, uc usecases.APIGOUsecaseInterface) {
	handler := &handlers.APIGOHandler{
		APIGOUsecase: uc,
	}

	route := r.Group(baseUrl)
	route.GET("health-check", handler.Check)

	employees := route.Group("employees")
	employees.POST("/", handler.CreateEmployee)
	employees.GET("/", handler.GetEmployee)
	employees.GET("/:id", handler.DetailEmployee)
	employees.PUT("/:id", handler.UpdateEmployee)
	employees.DELETE("/:id", handler.DeleteEmployee)
}
