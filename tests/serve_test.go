package tests

import (
	_"os"
	_ "fmt"
	_ "net/http"
	_ "net/http/httptest"
	 "testing"
	"github.com/gin-gonic/gin"
)

var testRouter *gin.Engine

// EFFECT: Initializes shared testing variables.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	//testRouter := *gin.Default()


}
func TestInitTestDB(t *testing.T) {
	err := InitTestDB()
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
}


