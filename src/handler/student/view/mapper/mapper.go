package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view/contract"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(profile domain.Profile, classesDone []domain.Class) contract.Response {
	classesResponse := make([]contract.Class, len(classesDone))
	for i := range classesDone {
		classesResponse[i] = contract.Class{
			ClassID: classesDone[i].ClassID(),
			Title:   classesDone[i].Title(),
		}
	}

	return contract.Response{
		Email:       profile.Email(),
		Name:        profile.Name(),
		ClassesDone: classesResponse,
	}
}
