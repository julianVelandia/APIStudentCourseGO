package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/profile/contract"
)

type Mapper interface {
	DomainToResponse(profile domain.Profile) contract.Response
}

type UseCase interface {
	Execute(email string) (domain.Profile, error)
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

	domainProfile, errorUseCase := h.useCase.Execute(request.Email)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile)
	ginCTX.JSON(http.StatusOK, response)
}
