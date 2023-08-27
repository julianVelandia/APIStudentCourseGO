package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view/contract"
)

type Mapper interface {
	DomainToResponse(profile domain.Profile, classesDone []domain.Class) contract.Response
}

type UseCase interface {
	Execute(email string) (domain.Profile, []domain.Class, error)
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

	domainProfile, classesDone, errorUseCase := h.useCase.Execute(request.Email)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile, classesDone)
	ginCTX.JSON(http.StatusOK, response)
}
