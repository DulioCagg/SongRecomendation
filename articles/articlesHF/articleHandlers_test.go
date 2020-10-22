package articlesHF

import (
	"github.com/DulioCagg/webMicroservices/utils/testut"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// This test checks that the status response is 200 and the title of the page
// is Dulio App
func TestIndexPage(t *testing.T) {
	r := testut.GetRouter(true)
	r.GET("/", GetAll)

	// Creates a new request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testut.HTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Checks if the status is ok
		statusOK := w.Code == http.StatusOK

		// Test that the title is correct
		p, err := ioutil.ReadAll(w.Body)
		title := strings.Index(string(p), "<title>Dulio App</title>")
		pageOK := title > 0 && err == nil

		return statusOK && pageOK
	})
}

