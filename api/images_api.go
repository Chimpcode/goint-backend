package api

import (
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris"
	"../graphql"
	"../galaxy"
	"../utils"
	"strings"
)


func LinkStorageAPI(party router.Party, config *utils.GointConfig) error {
	err :=galaxy.InitGointStorage(config)
	if err != nil {
		return err
	}

	party.Get("/i/{star: string}/{imageName: string}", func(c iris.Context) {
		star := c.Params().Get("star")
		exist, err := graphql.CheckIfCompanyExistByRuc(star)

		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		if !exist {
			c.StatusCode(iris.StatusBadRequest)
			c.JSON(iris.Map{
				"error": "Company not exists",
			})
			return
		}

		imageName := c.Params().Get("imageName")
		if !strings.HasSuffix(imageName, ".png") {
			c.StatusCode(iris.StatusBadRequest)
			c.JSON(iris.Map{
				"error": "Invalid image name",
			})
			return
		}

		pathOfImage, err := galaxy.DownloadPhoto(imageName, star)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}


		err = c.ServeFile(pathOfImage, false)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}
		c.StatusCode(iris.StatusOK)
		return


	})

	party.Post("/i/{star: string}", func(c iris.Context) {
		star := c.Params().Get("star")
		if star != "milkyway" {
			exist, err := graphql.CheckIfCompanyExistByRuc(star)

			if err != nil {
				c.StatusCode(iris.StatusInternalServerError)
				c.JSON(iris.Map{
					"error": err.Error(),
				})
				return
			}

			if !exist {
				c.StatusCode(iris.StatusBadRequest)
				c.JSON(iris.Map{
					"error": "Company not exists",
				})
				return
			}
		} else if star == "milkyway" {
			
		}


		file, info, err := c.FormFile("image")
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		contentType := info.Header["Content-Type"][0]
		if contentType != "image/png" {
			c.StatusCode(iris.StatusBadRequest)
			c.JSON(iris.Map{
				"file_type": contentType,
				"error": "File type not support",
			})
			return
		}

		filename, err := galaxy.UploadPhoto(file, info.Filename, star)
		if err != nil {
			c.StatusCode(iris.StatusInternalServerError)
			c.JSON(iris.Map{
				"error": err.Error(),
			})
			return
		}

		c.StatusCode(iris.StatusOK)
		c.JSON(iris.Map{
			"filename": filename,
		})

	})

	return nil


}