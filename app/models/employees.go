package models

import "time"

type Employee struct {
	ID        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	HireDate  time.Time  `json:"hire_date"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type EmployeeRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	HireDate  string `json:"hire_date"`
}

type EmployeeRules struct {
	FirstName string `valid:"required~parameter is empty"`
	LastName  string `valid:"required~parameter is empty"`
	Email     string `valid:"required~parameter is empty,email~Please provide a valid email address"`
	HireDate  string `valid:"required~parameter is empty"`
}
