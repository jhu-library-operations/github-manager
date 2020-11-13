package api

import (
	"time"
	"context"
	"fmt"
)

type Organization struct {
	AvatarUrl string
	CreatedAt time.Time
	Description string
	Email string
	Id string
	Login string
	Name string
}

func getOrganization(c *Client, login string) {
	type ResponseData struct {
		Organization struct {
			AvatarUrl string
			CreatedAt string
			Description string
			Email string
			ID string
			Login string
			Name string
		}`graphql:"organization(login: $login)"`
	}

	variables := map[string]interface{}{
		"login": login,
	}

	gql := graphQLClient(c.http, "https://api.github.com/graphql")

	var org Organization
	for {
		var query ResponseData
		err := gql.QueryNamed(context.Background(), "Organization", &query, variables)
		if err != nil {
			panic(err)
		}

		org = query.Organization
	}

	fmt.Println(org)
			
}
