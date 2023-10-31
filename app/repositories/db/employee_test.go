package db

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dalas98/mekari-test/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDBRepository_GetEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	timeMock := time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating GORM DB: %v", err)
	}

	a := &DBRepository{db: gormDB}

	type fields struct {
		db *gorm.DB
	}

	type args struct {
		ctx *gin.Context
	}

	columns := []string{
		"id", "first_name", "last_name", "email", "hire_date", "created_at", "updated_at",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args)
	}{
		{
			name: "Negative | Record not found",
			fields: fields{
				db: a.db,
			},
			args: args{
				ctx: &gin.Context{},
			},
			wantErr: true,
			mock: func(a args) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "employees"`)).WillReturnError(errors.New("record not found"))
				mock.ExpectRollback()
			},
		},
		{
			name: "Positive | Record found",
			fields: fields{
				db: a.db,
			},
			args: args{
				ctx: &gin.Context{},
			},
			wantErr: false,
			mock: func(a args) {
				mock.ExpectBegin()
				rows := mock.NewRows(columns).AddRow(1, "Yusuf", "Farhan Hasbullah", "dalas98@gmail.com", timeMock, timeMock, timeMock)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "employees"`)).WillReturnRows(rows)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			_, err := a.GetEmployee(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.GetEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDBRepository_DetailEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	timeMock := time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating GORM DB: %v", err)
	}

	a := &DBRepository{db: gormDB}

	type fields struct {
		db *gorm.DB
	}

	type args struct {
		ctx *gin.Context
		id  int64
	}

	columns := []string{
		"id", "first_name", "last_name", "email", "hire_date", "created_at", "updated_at",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args)
	}{
		{
			name: "Negative | Record not found",
			args: args{
				ctx: &gin.Context{},
				id:  1,
			},
			fields: fields{
				db: a.db,
			},
			wantErr: true,
			mock: func(a args) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "employees" WHERE (id = $1)`)).WillReturnError(errors.New("record not found"))
				mock.ExpectRollback()
			},
		},
		{
			name: "Positive | Record found",
			fields: fields{
				db: a.db,
			},
			args: args{
				ctx: &gin.Context{},
			},
			wantErr: true,
			mock: func(a args) {
				mock.ExpectBegin()
				rows := mock.NewRows(columns).AddRow(1, "Yusuf", "Farhan Hasbullah", "dalas98@gmail.com", timeMock, timeMock, timeMock)
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "employees"`)).WillReturnRows(rows)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			_, err := a.DetailEmployee(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.DetailEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDBRepository_StoreEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	timeMock := time.Date(2023, 12, 05, 0, 0, 0, 0, time.UTC)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating GORM DB: %v", err)
	}

	a := &DBRepository{db: gormDB}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx      *gin.Context
		employee *models.Employee
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(args)
	}{
		{
			name: "Positive | Record found",
			fields: fields{
				db: a.db,
			},
			args: args{
				ctx: &gin.Context{},
				employee: &models.Employee{
					FirstName: "Yusuf",
					LastName:  "Farhan Hasbullah",
					Email:     "dalas98@gmail.com",
					HireDate:  timeMock,
				},
			},
			wantErr: true,
			mock: func(a args) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO employees (.+) VALUES (.+)").
					WithArgs(a).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.args)
			if err := a.StoreEmployee(tt.args.ctx, tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.StoreEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
