package usecase

import (
	"context"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
)

type RepositoryList interface {
	GetByQuery(qry query.List) ([]domain.Class, error)
}

type UseCase struct {
	repositoryRead RepositoryList
}

func NewUseCase(repositoryList RepositoryList) UseCase {
	return UseCase{
		repositoryRead: repositoryList,
	}
}

func (uc UseCase) Execute(ctx context.Context, qry query.List) ([]domain.Class, error) {

	classesEntities, errRepository := uc.repositoryRead.GetByQuery(ctx, qry)

	if errRepository != nil {
		return make([]domain.Class, 0), errRepository
	}

	return classesEntities, nil
}
