package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	//"log"
	//"time"
	//"path/filepath"
	//"os"
	)

func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	fmt.Println(client.Projects)
	fmt.Println("organization printing....")
	return orgs, err
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)
	//bring the username from input stream
	organizations, err := fetchOrganizations(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}

	fmt.Println("this is a line-break")
}
