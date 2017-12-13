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
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"},
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

	app.Use(crs)

	api := app.Party("/api/v1")

	err = LinkMiddleAPI(api)

	if err != nil {
		panic(err)
	}



	err = app.Run(iris.Addr(":9300"))
	if err != nil {
		panic(err)
	}
}
