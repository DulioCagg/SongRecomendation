package articles

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var ArticleList = []Article{
	{
		ID:      1,
		Title:   "First",
		Content: "Article 1 content",
	},
	{
		ID:      2,
		Title:   "Second",
		Content: "Article 2 content",
	},
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticle(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}