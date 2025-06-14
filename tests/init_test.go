package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"restful-movie-api/controllers"
	"restful-movie-api/repositories"
	"restful-movie-api/services"

	"github.com/gin-gonic/gin"
)

var (
	testRouter *gin.Engine
	testDB     *sql.DB
)

// EFFECT: Initializes shared testing variables.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	testRouter = gin.Default()

	var err error

	testDB, err = InitTestDB()
	if err != nil {
		log.Printf("Connection error to testDB: %v", err)
		log.Printf("Cannot run tests")
		os.Exit(1)
	}

	// Initialize Layers
	testRepo := repositories.UserRepository{DB: testDB}
	testSC := services.UserServices{Repo: &testRepo}
	testUC := controllers.UserController{Service: &testSC}

	// Initialize TestRoutes
	testRouter.GET("/users", testUC.GetAllUsers)
	testRouter.POST("/user", testUC.UploadNewUser)

	code := m.Run() // run all tests

	// Close
	testDB.Close()
	os.Exit(code)
}

// EFFECT: Helper function to clean/reset the testDB
//
//	Used before every test to ensure precision
func CleanTestDB() {
	_, err := testDB.Exec(`
        TRUNCATE TABLE reviews, movies, users RESTART IDENTITY CASCADE;
  `)

	if err != nil {
		panic(err)
	}
}
