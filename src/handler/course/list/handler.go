package ask

import (
	"context"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/platform/log"
	contract2 "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/list/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Execute(ctx context.Context, qry query.List) ([]domain.Course, error)
}

type Mapper interface {
	EntityToResponse(entities []domain.Course) contract2.Response
	RequestToQuery(request contract2.URLParams) (query.List, error)
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

	requestParam := &contract2.URLParams{}

	if errValidator := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); errValidator != nil {
		log.Handler{}.Err(ginCTX, "Validator", errValidator, http.StatusBadRequest)
		return
	}
	qry, errMapper := h.mapper.RequestToQuery(*requestParam)
	if errMapper != nil {
		return errMapper
	}

	entities, errUseCase := h.useCase.Execute(ginCTX, qry)
	if errUseCase != nil {
		return errUseCase
	}

	response := h.mapper.EntityToResponse(entities)

	ginCTX.JSON(http.StatusOK, response)
}
