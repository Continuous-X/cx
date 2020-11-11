package github

import (
	"cx/pkg/output"
	"cx/pkg/metrics"
	"fmt"
	"github.com/google/go-github/v32/github"
)

type GHRepository struct {
	Organisation string
	RepositoryName string
	GhToken string
}

func (ghRepo GHRepository) IsCompliant() bool {
	ghRepo.getRepositoryList()
	ghRepo.writeMetric()
	return false
}

func (ghRepo GHRepository) writeMetric()  {
	metrics.MetricGHRepositoryProtection{
		CliCommand:          "irgendwas",
		GhProtectionActive:  false,
		GhPullrequestActive: false,
		GhStatusCheckActive: false,
	}.WriteMetric()

}

func (ghRepo GHRepository) getRepositoryList() ([]*github.Repository, error) {

	client, ctx := GHBase{ghToken: ghRepo.GhToken}.getCient()
	repositoryList, _, listError := client.Repositories.List(ctx,"", nil)
	if listError != nil {
		return repositoryList, listError
	}
	for index, repository := range repositoryList {
		output.PrintLogfile(fmt.Sprintf("%d: %s", index, repository.GetFullName()))
	}
	return repositoryList, nil
}