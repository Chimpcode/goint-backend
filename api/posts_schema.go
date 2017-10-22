package api

import (
	"github.com/graphql-go/graphql"
	"../types"
	"../db"
	"strings"
)

var locationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Description: "Location type of goint project",
	Fields: graphql.Fields{
		"longitude": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.String,
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
					Type: graphql.String,
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"image": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"longitude": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"latitude": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"stock": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
		},
	},
})