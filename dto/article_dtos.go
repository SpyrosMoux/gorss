package dto

import "time"

type ArticleDto struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Link    string    `json:"link"`
	FeedId  string    `json:"feedId"`
	Date    time.Time `json:"date"`
}
