package main
import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}
//this code need API Key(given from github)
func main() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "<!-- Your API Keys -->"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	repo, _, err := client.Repositories.Get(context, "Golang-Coach", "Lessons")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName: *repo.FullName,
		Description: *repo.Description,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)
}