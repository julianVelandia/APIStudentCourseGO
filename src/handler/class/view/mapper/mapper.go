package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view/contract"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(classes domain.Class) contract.Response {

	return contract.Response{
		ClassID:      classes.ClassID(),
		Title:        classes.Title(),
		CreationDate: classes.CreationDate(),
		Content:      classes.Content(),
		ReadTime:     classes.ReadTime(),
	}
}
