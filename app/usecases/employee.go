package usecases

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dalas98/mekari-test/app/models"
	"github.com/dalas98/mekari-test/helpers"
	"github.com/gin-gonic/gin"
)

func (a *APIGOUsecase) CreateEmployee(ctx *gin.Context, employeeRequest *models.EmployeeRequest) (*helpers.ValidationErrors, error) {
	rules := &models.EmployeeRules{
		FirstName: employeeRequest.FirstName,
		LastName:  employeeRequest.LastName,
		Email:     employeeRequest.Email,
		HireDate:  employeeRequest.HireDate,
	}

	validate, err := govalidator.ValidateStruct(rules)
	if !validate && err != nil {
		errors := helpers.ValidationError(rules, err)
		return errors, err
	}

	hireDate, _ := time.Parse(time.DateOnly, employeeRequest.HireDate)

	employee := &models.Employee{
		FirstName: employeeRequest.FirstName,
		LastName:  employeeRequest.LastName,
		Email:     employeeRequest.Email,
		HireDate:  hireDate,
	}

	err = a.DBRepository.StoreEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *APIGOUsecase) GetEmployee(ctx *gin.Context) (*[]models.Employee, error) {
	employees, err := a.DBRepository.GetEmployee(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (a *APIGOUsecase) DetailEmployee(ctx *gin.Context, id int64) (*models.Employee, error) {
	employee, err := a.DBRepository.DetailEmployee(ctx, id)
	if err != nil {
		return nil, err
	}

	if employee.ID == 0 {
		return nil, errors.New("employee not found")
	}

	return employee, nil
}

func (a *APIGOUsecase) UpdateEmployee(ctx *gin.Context, id int64, employeeRequest *models.EmployeeRequest) (*helpers.ValidationErrors, error) {
	rules := &models.EmployeeRules{
		FirstName: employeeRequest.FirstName,
		LastName:  employeeRequest.LastName,
		Email:     employeeRequest.Email,
		HireDate:  employeeRequest.HireDate,
	}

	validate, err := govalidator.ValidateStruct(rules)
	if !validate && err != nil {
		errors := helpers.ValidationError(rules, err)
		return errors, err
	}

	hireDate, _ := time.Parse(time.DateOnly, employeeRequest.HireDate)

	employee := &models.Employee{
		FirstName: employeeRequest.FirstName,
		LastName:  employeeRequest.LastName,
		Email:     employeeRequest.Email,
		HireDate:  hireDate,
	}

	err = a.DBRepository.UpdateEmployee(ctx, id, employee)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *APIGOUsecase) DeleteEmployee(ctx *gin.Context, id int64) error {

	if err := a.DBRepository.DeleteEmployee(ctx, id); err != nil {
		return err
	}

	return nil
}
