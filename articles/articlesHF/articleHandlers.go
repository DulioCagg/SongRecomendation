package articlesHF

import (
	"github.com/DulioCagg/webMicroservices/articles"
	"github.com/DulioCagg/webMicroservices/utils/resput"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAll responds with the index html
func GetAll(c *gin.Context) {
	// Fetches all articles
	arts := articles.GetAllArticles()

	// Call the render function with the name of the template to render
	resput.Render(c, gin.H{
		"title":   "Dulio App",
		"payload": arts,
	}, "index.html")

}

func GetArticle(c *gin.Context) {
	// Extracts the article ID from the request path.
	// Tries to transform the id into a number, if invalid, return
	// with no found
	artID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	// Searches for the article by ID
	article, err := articles.GetArticle(artID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.HTML(
		// Status code
		http.StatusOK,
		// HTML file
		"article.html",
		// Pass data to template
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
