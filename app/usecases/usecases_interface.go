package usecases

import (
	"github.com/dalas98/mekari-test/app/models"
	"github.com/dalas98/mekari-test/app/repositories"
	"github.com/dalas98/mekari-test/helpers"
	"github.com/gin-gonic/gin"
)

type APIGOUsecase struct {
	DBRepository repositories.DBRepositoryInterface
}

type APIGOUsecaseInterface interface {
	HealthCheck() (*models.HealthCheck, error)

	CreateEmployee(ctx *gin.Context, employeeRequest *models.EmployeeRequest) (*helpers.ValidationErrors, error)
	GetEmployee(ctx *gin.Context) (*[]models.Employee, error)
	DetailEmployee(ctx *gin.Context, id int64) (*models.Employee, error)
	UpdateEmployee(ctx *gin.Context, id int64, employeeRequest *models.EmployeeRequest) (*helpers.ValidationErrors, error)
	DeleteEmployee(ctx *gin.Context, id int64) error
}

func NewAPIGOUsecase(repo repositories.DBRepositoryInterface) APIGOUsecaseInterface {
	return &APIGOUsecase{
		DBRepository: repo,
	}
}
