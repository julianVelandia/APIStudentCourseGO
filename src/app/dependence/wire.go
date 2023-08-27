package dependence

import (
	useCaseClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/usecase"
	useCaseStudent "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/application/usecase"
	handlerListClasses "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/list"
	handlerMapperListClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/list/mapper"
	handlerViewClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view"
	handlerViewProfile "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view"
)

type HandlerContainer struct {
	ViewProfileHandler handlerViewProfile.Handler
	ListClassesHandler handlerListClasses.Handler
	ViewClassHandler   handlerViewClass.Handler
}

func NewWire() HandlerContainer {

	return HandlerContainer{
		ListClassesHandler: newWireListClassesHandler(repositoryClassList),
		ViewClassHandler:   newWireViewClassHandler(repositoryClassView, repositoryClassList),
		ViewProfileHandler: newWireViewProfileHandler(),
	}
}

func newWireListClassesHandler(
	repositoryList repositoryClassSQL.SqlListRepository,
) handlerListClasses.Handler {

	useCaseList := useCaseListClass.NewUseCase(
		repositoryList,
	)
	return *handlerListClasses.NewHandler(
		useCaseList,
		handlerMapperListClass.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func newWireViewClassHandler(
	repositoryView repositoryClassSQL.SqlViewRepository,
	repositoryList repositoryClassSQL.SqlListRepository,
) handlerViewClass.Handler {

	useCaseView := useCaseClass.NewUseCase(
		repositoryView,
		repositoryList,
	)
	return *handlerViewClass.NewHandler(
		useCaseView,
		handlerMapperViewClass.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func newWireViewProfileHandler() handlerViewProfile.Handler {
	useCaseViewProfile := useCaseStudent.NewUseCase(
		repositoryWrite,
	)

	return handlerViewProfile.NewGetHandler(
		useCaseViewProfile,
	)
}
