package gql

import (
	"robotahu-server/postgres"

	"github.com/graphql-go/graphql"
)

type Root struct {
	Query    *graphql.Object
	Mutation *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"images": &graphql.Field{
						Type:    graphql.NewList(Image),
						Resolve: resolver.ImagesResolver,
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createImage": &graphql.Field{
						Type:    Image,
						Resolve: resolver.CreateImageResolver,
					},
				},
			},
		),
	}

	return &root
}
