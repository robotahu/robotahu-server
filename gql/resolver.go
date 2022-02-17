package gql

import (
	"robotahu-server/postgres"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) ImagesResolver(p graphql.ResolveParams) (interface{}, error) {
	image, err := r.db.GetImages()
	if err != nil {
		return nil, err
	}

	return image, nil
}
