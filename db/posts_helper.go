package db

import (
	"../types"
	"../utils"
	"github.com/asdine/storm/q"
	"time"
)

func PutPostInDb(post *types.Post) (string, error) {
	var err error
	if post.Id == "" {
		post.Id = utils.GetNewUUID()
	}
	if post.Image == "" {
		post.Image = "placeholder.jpeg"
	}
	post.CreatedAt = time.Now()
	err = MasterDB["posts"].Save(post)
	return post.Id, err
}

func GetPostById(id string) (*types.Post, error) {
	post := new(types.Post)
	err := MasterDB["posts"].One("Id", id, post)
	return post, err
}

func GetAllPosts() ([]types.Post, error) {
	var posts []types.Post
	err := MasterDB["posts"].All(&posts)
	return posts, err
}

func GetUsersByType(typePost string) ([]types.Post, error) {
	var posts []types.Post
	err := MasterDB["post"].Find("Type", typePost, &posts)
	return posts, err
}

func DeletePostById(id string) (*types.Post, error) {
	post, err := GetPostById(id)
	if err != nil {
		return post, err
	}
	err = MasterDB["posts"].DeleteStruct(post)
	return post, err

}

func DeleteAllPosts() error {
	query := MasterDB["posts"].Select(q.True())
	err := query.Delete(new(types.Post))
	return err
}

func UpdatePostById(modPost *types.Post) error {
	err := MasterDB["posts"].Update(modPost)
	return err
}
