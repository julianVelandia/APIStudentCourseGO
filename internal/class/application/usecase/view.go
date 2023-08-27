package usecase

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
)

type RepositoryViewClass interface {
	GetClassByClassID(classID string) (domain.Class, error)
}

type RepositoryUpdateClassesDone interface {
	UpdateClassesByEmail(cmd command.Update) error
}

type ViewUseCase struct {
	repositoryViewClass         RepositoryViewClass
	repositoryUpdateClassesDone RepositoryUpdateClassesDone
}

func NewViewUseCase(repositoryViewClass RepositoryViewClass, repositoryUpdateClassesDone RepositoryUpdateClassesDone) *ViewUseCase {
	return &ViewUseCase{repositoryViewClass: repositoryViewClass, repositoryUpdateClassesDone: repositoryUpdateClassesDone}
}

func (uc ViewUseCase) Execute(qry query.View) (domain.Class, error) {
	domainClass, err := uc.repositoryViewClass.GetClassByClassID(qry.ClassID())
	if err != nil {
		return domain.Class{}, err
	}

	cmd := command.NewUpdate(
		qry.Email(),
		qry.ClassID(),
		qry.Title(),
	)
	err = uc.repositoryUpdateClassesDone.UpdateClassesByEmail(*cmd)
	if err != nil {
		return domain.Class{}, err
	}

	return domainClass, nil
}
