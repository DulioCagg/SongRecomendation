package resput

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render HTML or JSON based on the 'Accept' header of the request.
// IF the header isn't specified, HTML will be render if the template exists.
func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
