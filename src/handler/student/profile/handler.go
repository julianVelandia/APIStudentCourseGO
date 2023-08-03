package profile

import (
	"context"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/profile/contract"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/platform/log"
)

type UseCase interface {
	Execute(ctx context.Context, email string) error
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type Handler struct {
	useCase UseCase
}

func (h Handler) Handler(ginCTX *gin.Context) {

	request := &contract.Request{}
	if errBinding := ginCTX.BindJSON(request); errBinding != nil {
		log.Handler{}.Err(ginCTX, "Binding", errBinding, http.StatusBadRequest)
		return
	}

	errorUseCase := h.useCase.Execute(ginCTX, request.Email)

	if errorUseCase != nil {
		log.Handler{}.Err(ginCTX, "UseCase", errorUseCase, http.StatusInternalServerError)
		return
	}
}
