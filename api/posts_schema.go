package api

import (
	"github.com/graphql-go/graphql"
	"../db"
	"../types"
)

var locationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Description: "Location type of goint project",
	Fields: graphql.Fields{
		"longitude": &graphql.Field{
			Type: graphql.Float,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
	},

})

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Description: "Post of goint project",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"by": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"type": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"image": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: locationType,
		},
		"stock": &graphql.Field{
			Type: graphql.Int,
		},
	},

})

var postQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PostQuery",
	Description: "Query for Goint posts",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type:        postType,
			Description: "Get single post",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery := params.Args["id"].(string)
				user, err := db.GetPostById(idQuery)
				return user, err
			},
		},
		"posts": &graphql.Field{
			Type:        graphql.NewList(postType),
			Description: "Get all posts",

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return db.GetAllPosts()
			},
		},
	},
})

var postMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PostMutation",
	Description: "Modifications for posts of Goint project",
	Fields: graphql.Fields{
		"create_post": &graphql.Field{
			Type:        postType,
			Description: "Create new post",
			Args: graphql.FieldConfigArgument{
				"by": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"image": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"address": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"longitude": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"latitude": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"stock": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				by := p.Args["by"].(string)
				type_ := p.Args["type"].(string)
				title := p.Args["title"].(string)
				image := p.Args["image"].(string)
				description := p.Args["description"].(string)
				address := p.Args["address"].(string)
				longitude := p.Args["longitude"].(float64)
				latitude := p.Args["latitude"].(float64)
				stock := p.Args["stock"].(int)

				var post = types.Post{
					By: by,
					Type: type_,
					Title: title,
					Image: image,
					Description: description,
					Address: address,
					Location: types.Location{
						Longitude: longitude,
						Latitude: latitude,
					},
					Stock: stock,
				}

				_, err := db.PutPostInDb(&post)
				return post, err
			},
		},
	},
})

var postSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    postQuery,
	Mutation: postMutation,
})

