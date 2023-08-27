package usecase_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/usecase"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/stretchr/testify/assert"
)

const (
	email        = "test@email.com"
	classID      = "class_1"
	title        = "title"
	creationDate = "26/08/2023"
	content      = "content"
	readTime     = 5.5
)

func TestViewExecuteWhenRepositoryResponseOkShouldReturnOK(t *testing.T) {
	classDomain := *domain.NewClass(
		classID,
		title,
		creationDate,
		[]string{content, content},
		readTime,
	)
	cmd := *command.NewUpdate(
		email,
		classID,
		title,
	)
	qry := *query.NewView(
		email,
		classID,
		title,
	)
	repositoryViewMock := new(RepositoryViewClassMock)
	repositoryUpdateMock := new(RepositoryUpdateClassesDoneMock)

	repositoryViewMock.On("GetClassByClassID", classID).Return(classDomain, nil).Once()
	repositoryUpdateMock.On("UpdateClassesByEmail", cmd).Return(nil).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewMock,
		repositoryUpdateMock,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.NoError(t, errResult)
	assert.Equal(t, result, classDomain)
	repositoryViewMock.AssertExpectations(t)
	repositoryUpdateMock.AssertExpectations(t)
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

func TestViewExecuteWhenRepositoryUpdateFailShouldReturnError(t *testing.T) {
	classDomain := *domain.NewClass(
		classID,
		title,
		creationDate,
		[]string{content, content},
		readTime,
	)
	cmd := *command.NewUpdate(
		email,
		classID,
		title,
	)
	qry := *query.NewView(
		email,
		classID,
		title,
	)
	repositoryViewMock := new(RepositoryViewClassMock)
	repositoryUpdateMock := new(RepositoryUpdateClassesDoneMock)

	repositoryViewMock.On("GetClassByClassID", classID).Return(classDomain, nil).Once()
	repositoryUpdateMock.On("UpdateClassesByEmail", cmd).Return(errors.New("")).Once()
	getUseCase := usecase.NewViewUseCase(
		repositoryViewMock,
		repositoryUpdateMock,
	)
	result, errResult := getUseCase.Execute(qry)

	assert.Error(t, errResult)
	assert.Equal(t, result, domain.Class{})
	repositoryViewMock.AssertExpectations(t)
	repositoryUpdateMock.AssertExpectations(t)
}
