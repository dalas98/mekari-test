package db

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalas98/mekari-test/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestStoreEmployee(t *testing.T) {
	// Create a new mock database using sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	timeMock := time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC)

	// Create a GORM DB instance using the mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating GORM DB: %v", err)
	}

	// Create a new instance of your DBRepository and set the GORM DB
	a := &DBRepository{db: gormDB}

	// Create a sample employee
	employee := &models.Employee{
		// Initialize with sample data
		FirstName: "Yusuf",
		LastName:  "Farhan Hasbullah",
		Email:     "dalas98@gmail.com",
		HireDate:  timeMock,
		CreatedAt: &timeMock,
		UpdatedAt: &timeMock,
	}

	// Configure the expectations for the GORM DB using sqlmock
	mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO employees").WillReturnResult(sqlmock.NewResult(1, 1)) // Adjust the query as needed
	mock.ExpectExec("INSERT INTO employees (.+) VALUES (.+)").
		WithArgs(employee.FirstName, employee.LastName, employee.Email, employee.HireDate, employee.CreatedAt, employee.UpdatedAt). // Provide actual field values here
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a mock Gin context
	ctx := &gin.Context{}

	// Call the function being tested
	err = a.StoreEmployee(ctx, employee)

	// Check for any errors (replace with your actual error handling)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %v", err)
	}
}
