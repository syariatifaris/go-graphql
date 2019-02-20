package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/syariatifaris/graphql/pkg/model"
)

//NewRootMutation creates new gql root mutation
func NewRootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create_artist": &graphql.Field{
				Type: newArtistObj("createArtist"),
				Args: graphql.FieldConfigArgument{
					"artist": &graphql.ArgumentConfig{
						Type: newArtistInputWithNameObj("createArtistObject"),
					},
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var artist model.Artist
					am := params.Args["artist"]
					err := mapstructure.Decode(am, &artist)
					if err != nil {
						return nil, err
					}
					model.Artists = append(model.Artists, artist)
					return artist, nil
				},
			},
		},
	})
}
