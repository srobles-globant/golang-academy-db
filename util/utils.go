// Package util exports functions that are not part of the app but are helpful
package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/srobles-globant/golang-academy-db/model"
)

// GetArticles returns the array of articles from an external API
func GetArticles() ([]model.Article, error) {
	resp, err := http.Get("http://challenge.getsandbox.com/articles")
	if err != nil {
		return nil, err
	}
	articlesBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var articles []model.Article
	err = json.Unmarshal(articlesBytes, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticle returns one article from an external API
func GetArticle(articleID int) (*model.Article, error) {
	resp, err := http.Get(fmt.Sprintf("http://challenge.getsandbox.com/articles/%d", articleID))
	if err != nil {
		return nil, err
	}
	articleBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var article *model.Article
	err = json.Unmarshal(articleBytes, article)
	if err != nil {
		return nil, err
	}
	return article, nil
}
