package extras

import (
	"github.com/wawandco/fako"
	"github.com/machinebox/graphql"
	"../utils"
	"log"
	"context"
)
type SimplePlace struct {
	Latitude float64 `fako:"latitude_direction"`
	Longitude float64 `fako:"longitude_direction"`
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
	req.Var("latitude", place.Latitude)
	req.Var("longitude", place.Longitude)

	// run it and capture the response
	ctx := context.Background()
	if err := client.Run(ctx, req, nil); err != nil {
		log.Fatal(err)
	}
}