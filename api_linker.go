package main

import (
	"./api"
	"github.com/kataras/iris/core/router"
)

func LinkMiddleAPI(party router.Party) error {
	if err := api.LinkUserSchema(party); err != nil {
		return err
	}

	api.LinkStorageAPI(party)
	return nil
}
