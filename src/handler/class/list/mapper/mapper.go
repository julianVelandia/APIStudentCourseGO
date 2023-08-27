package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/list/contract"
)

type HandlerMapper struct{}

func (hm HandlerMapper) RequestToQuery(request contract.URLParams) (query.List, error) {

}

func (hm HandlerMapper) DomainToResponse(entities []domain.Class) contract.Response {

}
