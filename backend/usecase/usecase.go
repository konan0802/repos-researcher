package usecase

import (
	"repos-researcher/domain/repository"
)

type Usecase interface {
}

type usecase struct {
	grp repository.GitHubRepository
	drp repository.DynamoDBRepository
}

func NewUsecase(grp repository.GitHubRepository, drp repository.DynamoDBRepository) Usecase {
	return &usecase{
		grp: grp,
		drp: drp,
	}
}
