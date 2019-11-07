package db

import (
	"encoding/json"

	"../../models"
)

// Article -
type Article models.Article

// Mock datas
var articles = []Article{
	Article{"1", "Hello", "Article Description", "Article Content"},
	Article{"2", "Hello 2", "Article Description", "Article Content"},
}

// AllArticle -
func AllArticle() []Article {
	return articles
}

// GetArticle -
func GetArticle(id string) Article {
	for _, article := range articles {
		if article.ID == id {
			return article
		}
	}

	return Article{}
}

// CreateArticle -
func CreateArticle(data []byte) Article {
	article := Article{}
	json.Unmarshal(data, &article)
	articles = append(articles, article)

	return article
}

// DeleteArticle -
func DeleteArticle(id string) Article {
	for index, article := range articles {
		if article.ID == id {
			articles = append(articles[:index], articles[index+1:]...)
			return article
		}
	}

	return Article{}
}
