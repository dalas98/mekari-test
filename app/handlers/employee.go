package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dalas98/mekari-test/app/models"
	"github.com/dalas98/mekari-test/helpers"
	"github.com/gin-gonic/gin"
)

func (h *APIGOHandler) CreateEmployee(c *gin.Context) {
	var err error
	reqBody := models.EmployeeRequest{}

	if err = c.BindJSON(&reqBody); err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	if valid, err := h.APIGOUsecase.CreateEmployee(c, &reqBody); err != nil {
		if valid != nil {
			helpers.RespondErrorJSON(c, http.StatusUnprocessableEntity, errors.New(valid.Message))
			return
		}

		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusCreated, gin.H{"message": "Employee Created"}, nil)
}

func (h *APIGOHandler) GetEmployee(c *gin.Context) {
	employees, err := h.APIGOUsecase.GetEmployee(c)
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusCreated, employees, nil)
}

func (h *APIGOHandler) DetailEmployee(c *gin.Context) {
	employeeId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	employee, err := h.APIGOUsecase.DetailEmployee(c, employeeId)
	if err != nil {
		helpers.RespondErrorJSON(c, http.StatusNotFound, err)
		return
	}

	helpers.RespondJSON(c, http.StatusCreated, employee, nil)
}

func (h *APIGOHandler) UpdateEmployee(c *gin.Context) {
	var err error
	reqBody := models.EmployeeRequest{}
	employeeId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err = c.BindJSON(&reqBody); err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	if valid, err := h.APIGOUsecase.UpdateEmployee(c, employeeId, &reqBody); err != nil {
		if valid != nil {
			helpers.RespondErrorJSON(c, http.StatusUnprocessableEntity, errors.New(valid.Message))
			return
		}

		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, gin.H{"message": "Employee Updated"}, nil)
}

func (h *APIGOHandler) DeleteEmployee(c *gin.Context) {
	employeeId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := h.APIGOUsecase.DeleteEmployee(c, employeeId); err != nil {
		helpers.RespondErrorJSON(c, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(c, http.StatusOK, gin.H{"message": "Employee Deleted"}, nil)
}
