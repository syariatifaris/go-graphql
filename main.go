package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/syariatifaris/graphql/pkg/gql"
)

var schema graphql.Schema

func main() {
	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    gql.NewRootQuery(),
		Mutation: gql.NewRootMutation(),
	})
	if err != nil {
		log.Fatalln("unable to create gql schema", err.Error())
		os.Exit(1)
	}
	schema = s
	http.HandleFunc("/graphql", resolve)
	log.Println("starting graphql..")
	err = http.ListenAndServe("0.0.0.0:12345", nil)
	if err != nil {
		log.Fatalln("err handling: ", err.Error())
	}
}

func readPostQuery(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	return data["query"].(string), nil
}

func resolve(w http.ResponseWriter, r *http.Request) {
	query, err := readPostQuery(r)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	json.NewEncoder(w).Encode(result)
}
