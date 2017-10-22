package types

import "time"

type User struct {
	Id string `json:"id" storm:"id"`

	Group string `json:"group"`
	CreatedAt time.Time `json:"created_at" storm:"index"`


	FullName string `json:"full_name" fako:"full_name"`
	Age int `json:"age"`
	Gender string `json:"gender" fako:"gender"`

	LoginType string `json:"login_type"`

	Email string `json:"email" fako:"email_address"`

	Username string `json:"username" fako:"username"`
	Password string `json:"password" fako:"simple_password"`

	FacebookAccount string `json:"facebook_account"`

	LastLocation string `json:"last_location"`

	FollowPosts []string `json:"follow_posts"`

}