package api

import (
	"github.com/minio/minio-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"../utils"
	"strings"
	"io/ioutil"
	"os"
)

var GointStorage *minio.Client

func init() {

}

func InitializedGointStorage(config *utils.GointConfig) error{

	endpoint := config.Storage.Endpoint
	accessKeyID := config.Storage.AccessKey
	secretAccessKey := config.Storage.SecretKey
	useSSL := false

	var err error
	GointStorage, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return err
	}
	return nil
}

func LinkStorageAPI(party router.Party) {
	location := "us-east-1"

	party.Get("/i/{id:string}", func(c iris.Context) {
		id := c.Params().Get("id")
		if !strings.Contains(id, ".jpg") {
			id += ".jpg"
		}

		existBucket, err := GointStorage.BucketExists("images")
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}

		if !existBucket {
			err := GointStorage.MakeBucket("images", location)
			if err != nil {
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(iris.Map{
					"error_at": err.Error(),
				})
				return
			}
		}

		if !strings.HasPrefix(id, "goint-") {
			id = "goint-" + id
		}

		image, err := GointStorage.GetObject("images", id, minio.GetObjectOptions{})
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}
		data, err := ioutil.ReadAll(image)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}
		err = ioutil.WriteFile("/tmp/image.jpg", data, os.ModePerm)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}
		c.SendFile("/tmp/image.jpg", "image.jpg")

	})

	party.Post("/i/{id:string}", func(c iris.Context) {
		newId := c.Params().Get("id")

		if newId == "" {
			newId = c.FormValue("name")
		}
		if newId == "" {
			newId = c.FormValue("id")
		}
		if !strings.Contains(newId, ".jpg") {
			newId += ".jpg"
		}
		file, header, err := c.FormFile("file")
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}


		err = ioutil.WriteFile("/tmp/image_w.jpg", data, os.ModeAppend|os.ModeDir)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}

		existBucket, err := GointStorage.BucketExists("images")
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}

		if !existBucket {
			err := GointStorage.MakeBucket("images", location)
			if err != nil {
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(iris.Map{
					"error_at": err.Error(),
				})
				return
			}
		}

		if !strings.HasPrefix(newId, "goint-") {
			newId = "goint-" + newId
		}

		_, err = GointStorage.PutObject("images", newId, file, header.Size, minio.PutObjectOptions{ContentType: "image/jpg"})
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error_at": err.Error(),
			})
			return
		}
	})
}