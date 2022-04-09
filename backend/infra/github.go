package infra

import "repos-researcher/domain/repository"

type githubInfra struct{}

func NewGitHubInfra() repository.GitHubRepository {
	return &githubInfra{}
}
