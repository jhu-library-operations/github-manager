package main

import (
	"fmt"
	"context"
	"os"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type EnterpriseResponse struct {
	Name string `json:"name"`
	Id string `json:"id"`
}

func fetchOrganization(client *githubv4.Client, login string) (string, error) {
	var q struct {
		Organization struct {
			CreatedAt string
			Login string
		}`graphql:"organization(login: $login)"`
	}

	vars := map[string]interface{}{
		"login": githubv4.String(login),
	}

	err := client.Query(context.Background(), &q, vars)
	if err != nil {
		return "", err
	}


	return q.Organization.Login, nil
}

func fetchEnterpriseId(client *githubv4.Client, slug string) (string, error) {
	var q struct {
		Enterprise struct {
			Name string
			Id string
		}`graphql:"enterprise(slug: $slug)"`
	}

	vars := map[string]interface{}{
		"slug": githubv4.String(slug),
	}

	err := client.Query(context.Background(), &q, vars)
	if err != nil {
		return "", err
	}

	return q.Enterprise.Id, nil
}
func main() {
	var bearer = os.Getenv("GH_TOKEN")
	src := oauth2.StaticTokenSource( &oauth2.Token{AccessToken: bearer }, )
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	
	enterpriseId, err := fetchEnterpriseId(client, "johns-hopkins-university")

	fmt.Println(enterpriseId)
	
	login, err := fetchOrganization(client, "jhu-sheridan-libraries")
	fmt.Println(login)

	if err != nil { panic(err) }
}
