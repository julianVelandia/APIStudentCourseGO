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

type Handler interface {
	Handler(ginCtx *gin.Context)
}

type GetHandler struct {
	mapper  Mapper
	useCase UseCase
}

func NewGetHandler(mapper Mapper, useCase UseCase) *GetHandler {
	return &GetHandler{mapper: mapper, useCase: useCase}
}

func (h GetHandler) handler(ginCTX *gin.Context) {

	request := &contract.Request{}
	if errBinding := ginCTX.BindJSON(request); errBinding != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	qry := query.NewView(request.ClassID, request.Email)
	domainProfile, errorUseCase := h.useCase.Execute(*qry)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile)
	ginCTX.JSON(http.StatusOK, response)
}
