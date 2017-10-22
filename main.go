package main

import (
	"./db"
	"./utils"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
)

func main() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})


	lcpConfig, err := utils.GetConfiguration("config.json")

	if err != nil {
		panic(err.Error())
	}

	err = db.InitDB(lcpConfig)

	if err != nil {
		panic(err)
	}

	/*
		err = db.FeedDbWithFakeUsers(10)

		if err != nil {
			panic(err)
		}
	*/

	app := iris.New()


	api := app.Party("/api/v1")

	api.Use(crs)


	err = LinkMiddleAPI(api)

	if err != nil {
		panic(err)
	}

	err = app.Run(iris.Addr(":9300"))
	if err != nil {
		panic(err)
	}
}
