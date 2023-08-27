package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/profile/contract"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(profile domain.Profile) contract.Response {
	return contract.Response{
		Email: profile.Email(),
		Name:  profile.Name(),
	}
}
