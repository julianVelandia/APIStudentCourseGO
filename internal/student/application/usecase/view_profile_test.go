package usecase_test

import (
	"errors"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/application/usecase"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/stretchr/testify/assert"
)

const (
	email      = "test@email.com"
	classID1   = "class_1"
	classID2   = "class_2"
	title      = "title"
	name       = "julián"
	className1 = "name 1"
	className2 = "name 2"
)

func TestViewExecuteWhenRepositoryResponseOkShouldReturnOK(t *testing.T) {

	classesDomain := []domain.Class{
		*domain.NewClass(classID1, className1),
		*domain.NewClass(classID2, className2),
	}

	profileDomain := *domain.NewProfile(
		email,
		name,
	)

	profileDomain
	classDomain
	cmd := *command.NewUpdate(
		email,
		classID,
		title,
	)
	repositoryViewProfile := new(RepositoryViewProfileMock)

	repositoryViewProfile.On("GetProfileByEmail", email).Return(profileDomain, nil).Once()
	repositoryViewProfile.On("GetClassesDoneByEmail", email).Return(classesDomain, nil).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewProfile,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.NoError(t, errResult)
	assert.Equal(t, result, classDomain)
	repositoryViewProfile.AssertExpectations(t)
}

func TestViewExecuteWhenRepositoryViewFailShouldReturnError(t *testing.T) {
	qry := *query.NewView(
		email,
		classID,
		title,
	)
	repositoryViewMock := new(RepositoryViewClassMock)

	repositoryViewMock.On("GetClassByClassID", classID).Return(
		domain.Class{},
		errors.New(""),
	).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewMock,
		nil,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.Error(t, errResult)
	assert.Equal(t, result, domain.Class{})
	repositoryViewMock.AssertExpectations(t)
}
