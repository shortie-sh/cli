package cmd

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// __createRedirectGenEndingInput is used internally by genqlient
type __createRedirectGenEndingInput struct {
	Url string `json:"url"`
}

// __createRedirectInput is used internally by genqlient
type __createRedirectInput struct {
	Url    string `json:"url"`
	Ending string `json:"ending"`
}

// __getRedirectInput is used internally by genqlient
type __getRedirectInput struct {
	Ending string `json:"ending"`
}

// createRedirectCreateRedirect includes the requested fields of the GraphQL type Redirect.
type createRedirectCreateRedirect struct {
	Url    string `json:"url"`
	Ending string `json:"ending"`
}

// createRedirectGenEndingCreateRedirect includes the requested fields of the GraphQL type Redirect.
type createRedirectGenEndingCreateRedirect struct {
	Url    string `json:"url"`
	Ending string `json:"ending"`
}

// createRedirectGenEndingResponse is returned by createRedirectGenEnding on success.
type createRedirectGenEndingResponse struct {
	CreateRedirect createRedirectGenEndingCreateRedirect `json:"createRedirect"`
}

// createRedirectResponse is returned by createRedirect on success.
type createRedirectResponse struct {
	CreateRedirect createRedirectCreateRedirect `json:"createRedirect"`
}

// getRedirectGetRedirect includes the requested fields of the GraphQL type Redirect.
type getRedirectGetRedirect struct {
	Url string `json:"url"`
}

// getRedirectResponse is returned by getRedirect on success.
type getRedirectResponse struct {
	GetRedirect getRedirectGetRedirect `json:"getRedirect"`
}

func getRedirect(
	ctx context.Context,
	client graphql.Client,
	ending string,
) (*getRedirectResponse, error) {
	__input := __getRedirectInput{
		Ending: ending,
	}
	var err error

	var retval getRedirectResponse
	err = client.MakeRequest(
		ctx,
		"getRedirect",
		`
query getRedirect ($ending: String!) {
	getRedirect(ending: $ending) {
		url
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func createRedirect(
	ctx context.Context,
	client graphql.Client,
	url string,
	ending string,
) (*createRedirectResponse, error) {
	__input := __createRedirectInput{
		Url:    url,
		Ending: ending,
	}
	var err error

	var retval createRedirectResponse
	err = client.MakeRequest(
		ctx,
		"createRedirect",
		`
mutation createRedirect ($url: String!, $ending: String) {
	createRedirect(url: $url, ending: $ending) {
		url
		ending
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}

func createRedirectGenEnding(
	ctx context.Context,
	client graphql.Client,
	url string,
) (*createRedirectGenEndingResponse, error) {
	__input := __createRedirectGenEndingInput{
		Url: url,
	}
	var err error

	var retval createRedirectGenEndingResponse
	err = client.MakeRequest(
		ctx,
		"createRedirectGenEnding",
		`
mutation createRedirectGenEnding ($url: String!) {
	createRedirect(url: $url) {
		url
		ending
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
