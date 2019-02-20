package gql

import (
	"github.com/graphql-go/graphql"
)

func newArtistObj(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"id":            &graphql.Field{Type: graphql.Int},
			"name":          &graphql.Field{Type: graphql.String},
			"birth_of_date": &graphql.Field{Type: graphql.String},
		},
	})
}

func newArtistInputWithNameObj(name string) *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name,
		Fields: graphql.InputObjectConfigFieldMap{
			"id":            &graphql.InputObjectFieldConfig{Type: graphql.Int},
			"name":          &graphql.InputObjectFieldConfig{Type: graphql.String},
			"birth_of_date": &graphql.InputObjectFieldConfig{Type: graphql.String},
		},
	})
}

func newMovieObj(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
			"year": &graphql.Field{Type: graphql.Int},
		},
	})
}
