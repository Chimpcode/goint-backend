package types

import "time"

type Company struct {
	Id string `json:"id" storm:"id"`
	CreatedAt time.Time `json:"created_at" storm:"index"`

	Name string `json:"name"`
	Categories []string `json:"categories"`

	PostsCount int `json:"posts_count"`

	ActivePosts []string `json:"active_posts"`
	Subscribes []string `json:"registered"`

}
