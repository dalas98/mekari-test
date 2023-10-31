package db

import (
	"errors"
	"fmt"

	"github.com/dalas98/mekari-test/app/models"
	"github.com/gin-gonic/gin"
)

func (a *DBRepository) StoreEmployee(ctx *gin.Context, employee *models.Employee) error {
	tx := a.db.WithContext(ctx).Begin()
	if err := tx.Debug().Create(&employee).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (a *DBRepository) GetEmployee(ctx *gin.Context) (*[]models.Employee, error) {
	var employees []models.Employee
	tx := a.db.WithContext(ctx).Begin()
	if err := tx.Find(&employees).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &employees, nil
}

func (a *DBRepository) DetailEmployee(ctx *gin.Context, id int64) (*models.Employee, error) {
	var employee models.Employee

	tx := a.db.WithContext(ctx).Begin()
	if err := tx.Where("id", id).Find(&employee).Error; err != nil {
		return nil, err
	}

	if employee.ID == 0 {
		return nil, errors.New("employee not found")
	}

	tx.Commit()

	return &employee, nil
}

func (a *DBRepository) UpdateEmployee(ctx *gin.Context, id int64, employee *models.Employee) error {
	var emp models.Employee

	fmt.Println("email", employee.Email)

	tx := a.db.WithContext(ctx).Begin()

	tx.Where("id", id).Find(&emp)
	if emp.ID == 0 {
		return errors.New("employee not found")
	}

	if err := tx.Where("id", id).Updates(employee).Error; err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func (a *DBRepository) DeleteEmployee(ctx *gin.Context, id int64) error {
	var employee models.Employee

	tx := a.db.WithContext(ctx).Begin()

	tx.Where("id", id).Find(&employee)
	if employee.ID == 0 {
		return errors.New("employee not found")
	}

	if err := tx.Where("id", id).Delete(employee).Error; err != nil {
		return err
	}
	tx.Commit()

	return nil
}
