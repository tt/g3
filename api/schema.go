package api

import (
	"os"

	"github.com/graphql-go/graphql"
	"github.com/tt/g3/bank"
)

var Schema *graphql.Schema

func init() {
	var err error
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}

	Schema = &schema
}

var accountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"account": &graphql.Field{
			Type: accountType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(string)

				return accountTable.Find(id)
			},
		},
		"accounts": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(accountType)),
			Resolve: func(graphql.ResolveParams) (interface{}, error) {
				return accountTable, nil
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"openAccount": &graphql.Field{
			Type: accountType,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				client, err := bank.Dial(os.Getenv("BANK_ADDR"))
				if err != nil {
					return nil, err
				}

				defer client.Conn.Close()

				id, err := client.OpenAccount()
				if err != nil {
					return nil, err
				}

				return Account{ID: id}, nil
			},
		},
	},
})
