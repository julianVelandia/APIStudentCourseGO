package dependence

import (
	useCaseListClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/usecase"
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

	environmentConfig := configuration.GetConfig()
	sqlClient := environmentConfig.ConfigMySQLClient().NewMysqlClient(environmentConfig.Scope().Value())
	repositoryMapperClass := repositoryClassMapperSQL.Mapper{}

	repositoryClassList := repositoryClassSQL.NewSqlListRepository(sqlClient, repositoryMapperClass)
	repositoryClassView := repositoryClassSQL.NewSqlViewRepository(sqlClient, repositoryMapperClass)

	repositoryStudentWrite := repositoryStudentSQL.NewSqlWriteRepository(sqlClient)

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

	useCaseView := useCaseViewClass.NewUseCase(
		repositoryView,
		repositoryList,
	)
	return *handlerViewClass.NewHandler(
		useCaseView,
		handlerMapperViewClass.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func newWireViewProfileHandler() handlerSaveNewsLetter.Handler {
	useCaseSaveEmailNewsLetter := useCaseNewsLetterEmail.NewUseCase(
		repositoryWrite,
	)

	return *handlerSaveNewsLetter.NewHandler(
		useCaseSaveEmailNewsLetter,
	)
}
