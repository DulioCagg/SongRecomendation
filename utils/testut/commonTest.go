package testut

import (
	"github.com/DulioCagg/webMicroservices/articles"
	"github.com/gin-gonic/gin"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []articles.Article

// Main function is used for setup before executing other test functions
func TestMain(m *testing.M) {
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Runs the test cases
	os.Exit(m.Run())
}

// GetRouter is a helper function to create a router during testing
func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../../templates/*")
	}
	return r
}

// HTTPResponse is a helper function to process a request and test its response
func HTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	is := is.New(t)
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request
	r.ServeHTTP(w, req)

	is.Equal(f(w), true)
}

// SaveList is used to store the main list into a temp one to avoid unwanted changes
func SaveLists() {
	tmpArticleList = articles.ArticleList
}

// RestoreFunction is used to restore the main list from the temp one
func restoreLists() {
	articles.ArticleList = tmpArticleList
}