package github

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/v32/github"
)

type GHRepository struct {
	Organisation string
	RepositoryName string
	GHToken string
}

func (ghRepo GHRepository) getLatestRelease()  {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghRepo.GHToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	client.Repositories.List(ctx,)
}