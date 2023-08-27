package list

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/list/contract"
)

type UseCase interface {
	Execute(ctx context.Context, qry query.List) ([]domain.Class, error)
}

type Mapper interface {
	DomainToResponse(entities []domain.Class) contract.Response
	RequestToQuery(request contract.URLParams) (query.List, error)
}

type ValidationParams interface {
	BindParamsAndValidation(obj interface{}, params gin.Params) error
}

func NewHandler(useCase UseCase, mapper Mapper, validationParams ValidationParams) *Handler {
	return &Handler{
		useCase:          useCase,
		mapper:           mapper,
		validationParams: validationParams,
	}
}

type Handler struct {
	useCase          UseCase
	mapper           Mapper
	validationParams ValidationParams
}

func (h Handler) Handler(ginCTX *gin.Context) {

	requestParam := &contract.URLParams{}

	if errValidator := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); errValidator != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	qry, errMapper := h.mapper.RequestToQuery(*requestParam)
	if errMapper != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	entities, errUseCase := h.useCase.Execute(ginCTX, qry)
	if errUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, nil)
		return
	}

	response := h.mapper.DomainToResponse(entities)
	ginCTX.JSON(http.StatusOK, response)
}
