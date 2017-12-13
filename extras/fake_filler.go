package extras

import (
	"github.com/wawandco/fako"
	"github.com/machinebox/graphql"
	"../utils"
	"log"
	"context"
	"strconv"
)
type SimplePlace struct {
	Latitude string `fako:"latitude_direction"`
	Longitude string `fako:"longitude_direction"`
}

func FillFake(config *utils.GointConfig) {
	client := graphql.NewClient(config.GraphQLServer)

	// make a request
	//req := graphql.NewRequest(`
	//mutation ($title: String!, $address: String!, $description: String!, ) {
     //   createPost (id:$key) {
     //       field1
     //       field2
     //       field3
     //   }
	//}
	//`)
	req := graphql.NewRequest(`
    mutation ($latitude: Float!, $longitude: Float!) {
        createPlace (latitude: $latitude, longitude: $longitude) {
            id
			latitude
			longitude
        }
    }
	`)

	var place SimplePlace
	fako.Fill(&place)

	lat, _ := strconv.ParseFloat(place.Latitude, 64)
	lon, _ :=strconv.ParseFloat(place.Longitude, 64)

	req.Var("latitude", lat)
	req.Var("longitude", lon)

	// run it and capture the response
	ctx := context.Background()
	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
	}
}