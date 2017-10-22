package api

import (
	"github.com/graphql-go/graphql"
	"../types"
	"../db"
	"strings"
)

// var userFields = graphql.BindFields(types.User{})

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Description: "User of goint project",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"group": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"full_name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"gender": &graphql.Field{
			Type: graphql.String,
		},
		"login_type": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"facebook_account": &graphql.Field{
			Type: graphql.String,
		},
		"last_location": &graphql.Field{
			Type: graphql.String,
		},
		"follow_posts": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},

	},

})

var userQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserQuery",
	Description: "Unidirectional getting user schema",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        userType,
			Description: "Get single user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				idQuery := params.Args["id"].(string)
				user, err := db.GetUserById(idQuery)
				return user, err
			},
		},
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "Get all users",

			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return db.GetAllUsers()
			},
		},
	},
})


var userMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "UserMutation",
	Description: "Modifications for users of Goint project",
	Fields: graphql.Fields{
		"create_user": &graphql.Field{
			Type:        userType,
			Description: "Create new user",
			Args: graphql.FieldConfigArgument{
				"full_name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},

				"age": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},

				"gender": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"login_type": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"group": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "client",
				},

				"email": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},

				"username": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"password": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},

				"facebook_account": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				fullName := p.Args["full_name"].(string)
				age := p.Args["age"].(int)
				gender := p.Args["gender"].(string)
				login_type := p.Args["login_type"].(string)
				group := p.Args["group"].(string)
				email := p.Args["email"].(string)
				username := p.Args["username"].(string)
				password := p.Args["password"].(string)
				fbAccount := p.Args["facebook_account"].(string)

				var user = types.User{
					Group:    group,
					FullName: fullName,
					Email:    email,
					Username: username,
					Password: password,
					Age: age,
					Gender: gender,
					LoginType: login_type,
					FacebookAccount: fbAccount,
				}

				_, err := db.PutUserInDB(&user)

				return user, err
			},
		},

		"update_user": &graphql.Field{
			Type:        userType,
			Description: "Update any field on user struct",
			Args: graphql.FieldConfigArgument{

				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},

				"full_name": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"age": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},

				"gender": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"login_type": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"group": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "client",
				},

				"email": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"username": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"password": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"facebook_account": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},

				"follow_posts": &graphql.ArgumentConfig{
					Type:         graphql.NewList(graphql.String),
					DefaultValue: []string{},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)

				fullName := p.Args["full_name"].(string)
				age := p.Args["age"].(int)
				gender := p.Args["gender"].(string)
				login_type := p.Args["login_type"].(string)
				group := p.Args["group"].(string)
				email := p.Args["email"].(string)
				username := p.Args["username"].(string)
				password := p.Args["password"].(string)
				fbAccount := p.Args["facebook_account"].(string)
				follow := p.Args["follow_posts"].([]string)

				user, err := db.GetUserById(id)

				if err != nil {
					return user, err
				}

				if !strings.EqualFold(fullName, "") {
					user.FullName = fullName
				}
				if age != 0 {
					user.Age = age
				}
				if !strings.EqualFold(gender, "") {
					user.Gender = gender
				}
				if !strings.EqualFold(login_type, "") {
					user.LoginType = login_type
				}
				if !strings.EqualFold(group, "") {
					user.Group = group
				}
				if !strings.EqualFold(email, "") {
					user.Email = email
				}
				if !strings.EqualFold(username, "") {
					user.Username = username
				}
				if !strings.EqualFold(password, "") {
					user.Password = password
				}
				if !strings.EqualFold(fbAccount, "") {
					user.FacebookAccount = fbAccount
				}
				if len(follow) != 0 {
					user.FollowPosts = follow
				}
				err = db.UpdateUserById(user)

				return user, err

			},
		},

		"delete_user": &graphql.Field{
			Type:        userType,
			Description: "Delete one user using the id key",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				return db.DeleteUserById(id)
			},
		},

		"delete_users": &graphql.Field{
			Type:        graphql.String,
			Description: "Delete all users, be careful",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				err := db.DeleteAllUsers()
				return "Deleted all users", err
			},
		},
	},
})

var userSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    userQuery,
	Mutation: userMutation,
})

