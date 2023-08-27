package dependence

import (
	useCaseClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/usecase"
	repositoryViewClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/read"
	mapperRepositoryViewClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/read/mapper"
	repositoryUpdateClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/write"
	mapperRepositoryUpdateClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/write/mapper"
	useCaseStudent "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/application/usecase"
	repositoryViewProfile "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/read"
	mapperRepositoryViewProfile "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/read/mapper"
	handlerViewClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view"
	mapperViewClass "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view/mapper"
	handlerViewProfile "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view"
	mapperViewProfile "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view/mapper"
)

type HandlerContainer struct {
	ViewProfileHandler handlerViewProfile.Handler
	ViewClassHandler   handlerViewClass.Handler
}

func NewWire() HandlerContainer {
	repositoryClassRead := repositoryViewClass.NewClassRepositoryRead(
		mapperRepositoryViewClass.Mapper{},
	)
	repositoryClassUpdate := repositoryUpdateClass.NewClassRepositoryWrite(
		mapperRepositoryUpdateClass.Mapper{},
	)
	repositoryProfileRead := repositoryViewProfile.NewProfileRepositoryRead(
		mapperRepositoryViewProfile.Mapper{},
	)

	return HandlerContainer{
		ViewClassHandler: newWireViewClassHandler(
			*repositoryClassRead,
			*repositoryClassUpdate,
		),
		ViewProfileHandler: newWireViewProfileHandler(
			*repositoryProfileRead,
		),
	}
}

func newWireViewClassHandler(
	repositoryViewClass repositoryViewClass.ClassRepositoryRead,
	repositoryUpdateClass repositoryUpdateClass.ClassRepositoryWrite,
) handlerViewClass.Handler {

	useCaseView := useCaseClass.NewViewUseCase(
		repositoryViewClass,
		repositoryUpdateClass,
	)
	return *handlerViewClass.NewHandler(
		mapperViewClass.Mapper{},
		useCaseView,
	)
}

func newWireViewProfileHandler(
	repositoryViewProfile repositoryViewProfile.ProfileRepositoryRead,
) handlerViewProfile.Handler {
	useCaseViewProfile := useCaseStudent.NewViewUseCase(
		repositoryViewProfile,
	)

	return *handlerViewProfile.NewHandler(
		mapperViewProfile.Mapper{},
		useCaseViewProfile,
	)
}
