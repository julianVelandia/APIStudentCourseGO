package dependence

import (
	useCaseListCourse "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/application/usecase"
	useCaseViewCourse "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/application/usecase/view"
	repositoryCourseSQL "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/infrastructure/repository/mysql/course"
	repositoryCourseMapperSQL "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/course/infrastructure/repository/mysql/mapper"
	useCaseNewsLetterEmail "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/application/usecase"
	repositoryStudentSQL "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/mysql/student/newsletter"
	platformParams "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/platform/params"
	configuration "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/api/app/config"
	handlerListCourses "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/list"
	handlerMapperListCourse "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/list/mapper"
	handlerViewCourse "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/view"
	handlerMapperViewCourse "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/view/mapper"
	handlerSaveNewsLetter "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/profile"
)

type HandlerContainer struct {
	ListCoursesHandler handlerListCourses.Handler
	ViewCourseHandler  handlerViewCourse.Handler
	ViewProfileHandler handlerViewProfile.Handler
}

func NewWire() HandlerContainer {

	environmentConfig := configuration.GetConfig()
	sqlClient := environmentConfig.ConfigMySQLClient().NewMysqlClient(environmentConfig.Scope().Value())
	repositoryMapperCourse := repositoryCourseMapperSQL.Mapper{}

	repositoryCourseList := repositoryCourseSQL.NewSqlListRepository(sqlClient, repositoryMapperCourse)
	repositoryCourseView := repositoryCourseSQL.NewSqlViewRepository(sqlClient, repositoryMapperCourse)

	repositoryStudentWrite := repositoryStudentSQL.NewSqlWriteRepository(sqlClient)

	return HandlerContainer{
		ListCoursesHandler:    newWireListCoursesHandler(repositoryCourseList),
		ViewCourseHandler:     newWireViewCourseHandler(repositoryCourseView, repositoryCourseList),
		SaveNewsLetterHandler: newWireSaveNewsLetterHandler(*repositoryStudentWrite),
	}
}

func newWireListCoursesHandler(
	repositoryList repositoryCourseSQL.SqlListRepository,
) handlerListCourses.Handler {

	useCaseList := useCaseListCourse.NewUseCase(
		repositoryList,
	)
	return *handlerListCourses.NewHandler(
		useCaseList,
		handlerMapperListCourse.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func newWireViewCourseHandler(
	repositoryView repositoryCourseSQL.SqlViewRepository,
	repositoryList repositoryCourseSQL.SqlListRepository,
) handlerViewCourse.Handler {

	useCaseView := useCaseViewCourse.NewUseCase(
		repositoryView,
		repositoryList,
	)
	return *handlerViewCourse.NewHandler(
		useCaseView,
		handlerMapperViewCourse.HandlerMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func newWireSaveNewsLetterHandler(
	repositoryWrite repositoryStudentSQL.SqlWriteRepository,
) handlerSaveNewsLetter.Handler {
	useCaseSaveEmailNewsLetter := useCaseNewsLetterEmail.NewUseCase(
		repositoryWrite,
	)

	return *handlerSaveNewsLetter.NewHandler(
		useCaseSaveEmailNewsLetter,
	)
}
