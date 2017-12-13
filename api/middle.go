package api

import (
	"github.com/graphql-go/handler"
	"github.com/kataras/iris/core/router"
	//"github.com/kataras/iris/middleware/basicauth"
	//"time"
	"github.com/kataras/iris"
)

const ACTIVATEGRAPHIQL = true

func LinkUserSchema(party router.Party) error {

	/*
	authConfig… := basicauth.Config{
		Users:   map[string]string{"bregymr": "malpartida1", "admin": "admin"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	party.Use(authentication)
	*/

	userHandler := handler.New(&handler.Config{
		Schema:   &userSchema,
		Pretty:   true,
		GraphiQL: ACTIVATEGRAPHIQL,
	})

	postHandler := handler.New(&handler.Config{
		Schema:   &postSchema,
		Pretty:   true,
		GraphiQL: ACTIVATEGRAPHIQL,
	})

	party.Any("/u", func(c iris.Context) {
		userHandler.ServeHTTP(c.ResponseWriter(), c.Request())
	})
	party.Any("/p", func(c iris.Context) {
		postHandler.ServeHTTP(c.ResponseWriter(), c.Request())
	})
	return nil
}

