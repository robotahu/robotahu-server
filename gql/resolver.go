package gql

import (
	"robotahu-server/postgres"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) ImagesResolver(p graphql.ResolveParams) (interface{}, error) {
	images, err := r.db.GetImages()
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (r *Resolver) CreateImageResolver(p graphql.ResolveParams) (interface{}, error) {
	url := p.Args["url"].(string)
	caption := p.Args["caption"].(string)
	image, err := r.db.CreateImage(url, caption)

	if err != nil {
		return nil, err
	}

	return image, nil
}
