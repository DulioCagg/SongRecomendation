package app

import (
	"github.com/DulioCagg/webMicroservices/articles/articlesHF"
)

// Routes sets up all the routes with their respective controllers
func initRoutes() {
	// Handles home page, returns all articles
	router.GET("/", articlesHF.GetAll)

	// Returns the view of a single article
	router.GET("/article/view/:id", articlesHF.GetArticle)
}
