package view

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view/contract"
)

type Mapper interface {
	DomainToResponse(class domain.Class) contract.Response
}

type UseCase interface {
	Execute(qry query.View) (domain.Class, error)
}

type Handler struct {
	mapper  Mapper
	useCase UseCase
}

func NewHandler(mapper Mapper, useCase UseCase) *Handler {
	return &Handler{mapper: mapper, useCase: useCase}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	request := &contract.Request{}
	if errBinding := ginCTX.BindJSON(request); errBinding != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	qry := query.NewView(
		request.Email,
		request.ClassID,
		request.Title,
	)
	domainProfile, errorUseCase := h.useCase.Execute(*qry)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile)
	ginCTX.JSON(http.StatusOK, response)
}
