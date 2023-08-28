package read_test

import (
	"errors"
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/dto"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/read"
	"github.com/stretchr/testify/assert"
)

const (
	classID1            = "classid1"
	className1          = "clase 1"
	filenameClassesDone = "../../../../../../dbtest/StudentsClassesDone.json"
)

func TestGetClassesDoneByEmailWhenSuccessShouldReturnProfile(t *testing.T) {
	mockMapper := new(MapperMock)
	expectedClassesDTO := []dto.Class{
		{ClassID: classID1, Title: className1},
	}
	expectedClasses := []domain.Class{
		*domain.NewClass(classID1, className1),
	}
	mockMapper.On("DTOClassesToDomain", expectedClassesDTO).Return(expectedClasses, nil)
	repository := read.NewProfileRepositoryRead(mockMapper, filenameProfile, filenameClassesDone)
	result, err := repository.GetClassesDoneByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, expectedClasses, result)
	mockMapper.AssertExpectations(t)
}

func TestGetClassesDoneByEmailWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := read.NewProfileRepositoryRead(nil, "", "")
	result, err := repository.GetClassesDoneByEmail(email)

	assert.Error(t, err)
	assert.Equal(t, result, []domain.Class{})
	mockMapper.AssertExpectations(t)
}

func TestGetClassesDoneByEmailWhenMapperNotFoundShouldReturnEmpty(t *testing.T) {
	mockMapper := new(MapperMock)
	expectedClassesDTO := []dto.Class{
		{ClassID: classID1, Title: className1},
	}

	mockMapper.On("DTOClassesToDomain", expectedClassesDTO).Return([]domain.Class{}, errors.New(""))
	repository := read.NewProfileRepositoryRead(mockMapper, filenameProfile, filenameClassesDone)
	result, err := repository.GetClassesDoneByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, result, []domain.Class{})
	mockMapper.AssertExpectations(t)
}
