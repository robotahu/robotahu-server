package gql

import "github.com/graphql-go/graphql"

var Image = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Image",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
			"caption": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
