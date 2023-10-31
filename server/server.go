package server

import (
	"fmt"
	"log"
	"os"

	dbr "github.com/dalas98/mekari-test/app/repositories/db"
	"github.com/dalas98/mekari-test/app/usecases"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

type RestRunner interface {
	Run()
}

func (rs *Server) Run() {
	appPort := os.Getenv("PORT")

	log.Printf("%s", fmt.Sprintf("Server Running on port :%s", appPort))

	if err := rs.router.Run(":" + appPort); err != nil {
		log.Fatal(err)
	}

}

func NewServer() RestRunner {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db, err := dbr.NewPostgreSQLConn(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("error when starting db: ", err.Error())
	}

	dbRepo := dbr.NewDBRepository(db)

	us := usecases.NewAPIGOUsecase(dbRepo)

	APIGORoutes(r, us)

	return &Server{
		router: r,
	}
}
