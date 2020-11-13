package api

import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"

	"io/ioutil"

	"strings"
	"github.com/shurcooL/graphql"
)

type ClientOption  func(http.RoundTripper) http.RoundTripper

func NewHTTPClient(opts ...ClientOption) *http.Client {
	tr := http.DefaultTransport
	for _, opt := range opts {
		tr = opt(tr)
	}

	return &http.Client{Transport: tr}
}

func NewClientFromHTTP(httpClient *http.Client) *Client {
	client := &Client{http: httpClient }
	client.host = "https://api.github.com/graphql"
	return  client
}

func NewClient(opts ... ClientOption) *Client {
	client := &Client{http: NewHTTPClient(opts...)}
	return client
}

// func c *Client REST (method string, url string, 

type Client struct {
	http *http.Client
	host string
}

type graphQLResponse struct {
	Data interface{}
	Errors []GraphQLError
}

type GraphQLError struct {
	Type string
	Path []string
	Message string
}

type GraphQLErrorResponse struct {
	Errors []GraphQLError
}

func (gr GraphQLErrorResponse) Error() string {
	errorMessages := make([]string, 0, len(gr.Errors))
	for _,e := range gr.Errors {
		errorMessages = append(errorMessages, e.Message)
	}
	return fmt.Sprintf("GraphQL error: %s", strings.Join(errorMessages, "\n"))
}

func (c Client) GraphQL (host string, query string, variables map[string]interface{}, data interface{}) error {
	reqBody, err := json.Marshal(map[string]interface{}{"query": query, "variables": variables})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", host, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	
	defer resp.Body.Close()

	return handleResponse(resp, data)
}

func graphQLClient(h *http.Client, host string) *graphql.Client {
	return graphql.NewClient(host, h)
}

func handleResponse(resp *http.Response, data interface{}) error {
	success := resp.StatusCode >= 200 && resp.StatusCode < 300

	if !success {
		return HandleHTTPError(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	gr := &graphQLResponse{Data: data}

	err = json.Unmarshal(body, &gr)
	if err != nil {
		return err
	}

	if len(gr.Errors) > 0 {
		return &GraphQLErrorResponse{Errors: gr.Errors}
	}
	return nil
}

func HandleHTTPError(resp *http.Response) error {
	return nil
}
