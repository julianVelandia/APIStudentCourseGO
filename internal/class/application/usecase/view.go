package usecase

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
)

type RepositoryViewClass interface {
	GetClassByClassID(classID string) (domain.Class, error)
}

type RepositoryUpdateClassesDone interface {
	UpdateClassesByEmail(cmd command.View) error
}

type ViewUseCase struct {
	repositoryViewClass         RepositoryViewClass
	repositoryUpdateClassesDone RepositoryUpdateClassesDone
}

func (uc ViewUseCase) Execute(cmd command.View) (domain.Class, error) {

	domainClass, err := uc.repositoryViewClass.GetClassByClassID(cmd.ClassID())
	if err != nil {
		return domain.Class{}, err
	}

	err = uc.repositoryUpdateClassesDone.UpdateClassesByEmail(cmd)
	if err != nil {
		return domain.Class{}, err
	}

	return domainClass, nil
}
