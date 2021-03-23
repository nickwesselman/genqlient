package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"github.com/Khan/genql/graphql"
)

type SimpleQueryResponse struct {
	User SimpleQueryUser `json:"user"`
}

type SimpleQueryUser struct {
	Id string `json:"id"`
}

func SimpleQuery(
	client *graphql.Client,
) (*SimpleQueryResponse, error) {
	var retval SimpleQueryResponse
	err := client.MakeRequest(
		nil,
		`
query SimpleQuery {
	user {
		id
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}
