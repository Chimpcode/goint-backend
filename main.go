package main

import (
	"./db"
	"./utils"
	"./api"
	// "./extras"
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


	gointConfig, err := utils.GetConfiguration("config.json")

	if err != nil {
		panic(err.Error())
	}

	err = db.InitDB(gointConfig)
	if err != nil {
		panic(err)
	}

	app := iris.New()

	app.Use(crs)

	apiRoute := app.Party("/api/v1")

	if err := api.LinkUserSchema(apiRoute); err != nil {
		panic(err)
	}
	if err := api.LinkStorageAPI(apiRoute, gointConfig); err != nil {
		panic(err)
	}

	err = api.LinkAuthApi(apiRoute)
	if err != nil {
		panic(err)
	}

	app.Logger().SetLevel("debug")

	// extras.FillFake(gointConfig)
	err = app.Run(iris.Addr(":9300"))
	if err != nil {
		panic(err)
	}
}
