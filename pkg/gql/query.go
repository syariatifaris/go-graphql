package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/syariatifaris/graphql/pkg/model"
)

//NewRootQuery create new querie
func NewRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"movies": &graphql.Field{
				Type: graphql.NewList(newMovieObj("Movies")),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return model.Movies, nil
				},
			},
			"movie": &graphql.Field{
				Type: newMovieObj("Movie"),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(int)
					for i, m := range model.Movies {
						if m.ID == int64(id) {
							return model.Movies[i], nil
						}
					}
					return nil, nil
				},
			},
			"artist": &graphql.Field{
				Type: newArtistObj("Artist"),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return model.Artists[0], nil
				},
			},
		},
	})
}
