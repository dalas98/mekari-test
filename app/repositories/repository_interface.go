package repositories

import (
	"github.com/dalas98/mekari-test/app/models"
	"github.com/gin-gonic/gin"
)

//go:generate moq -out repository_interface_moq_test.go . DBRepositoryInterface

type DBRepositoryInterface interface {
	HealthCheck() (*models.HealthCheck, error)

	StoreEmployee(ctx *gin.Context, employee *models.Employee) error
	GetEmployee(ctx *gin.Context) (*[]models.Employee, error)
	DetailEmployee(ctx *gin.Context, id int64) (*models.Employee, error)
	UpdateEmployee(ctx *gin.Context, id int64, employee *models.Employee) error
	DeleteEmployee(ctx *gin.Context, id int64) error
}
