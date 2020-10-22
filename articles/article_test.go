package articles

import (
	"github.com/matryer/is"
	"testing"
)

func Test_getAllArticles(t *testing.T) {
	is := is.New(t)

	alist := GetAllArticles()

	// Checks if the len of the returned list is the same as teh variable
	is.Equal(len(alist), len(ArticleList))

	// Checks that members are identical
	for i, a := range alist {
		is.Equal(a.ID, ArticleList[i].ID)
		is.Equal(a.Title, ArticleList[i].Title)
		is.Equal(a.Content, ArticleList[i].Content)
	}
}
