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
	email           = "test1@email.com"
	name            = "Julián"
	filenameProfile = "../../../../../../dbtest/StudentsProfile.json"
)

func TestGetProfileByEmailWhenSuccessShouldReturnProfile(t *testing.T) {
	mockMapper := new(MapperMock)
	expectedProfileDTO := dto.Profile{
		Name: name,
	}
	expectedProfile := *domain.NewProfile(email, name)
	mockMapper.On("DTOProfileToDomain", email, expectedProfileDTO).Return(expectedProfile, nil)
	repository := read.NewProfileRepositoryRead(mockMapper, filenameProfile, filenameClassesDone)
	result, err := repository.GetProfileByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, expectedProfile, result)
	mockMapper.AssertExpectations(t)
}

func TestGetProfileByEmailWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := read.NewProfileRepositoryRead(nil, "", "")
	result, err := repository.GetProfileByEmail(email)

	assert.Error(t, err)
	assert.Equal(t, result, domain.Profile{})
	mockMapper.AssertExpectations(t)
}

func TestGetProfileByEmailWhenMapperNotFoundShouldReturnEmpty(t *testing.T) {
	mockMapper := new(MapperMock)
	expectedProfileDTO := dto.Profile{
		Name: name,
	}
	mockMapper.On("DTOProfileToDomain", email, expectedProfileDTO).Return(domain.Profile{}, errors.New(""))
	repository := read.NewProfileRepositoryRead(mockMapper, filenameProfile, filenameClassesDone)
	result, err := repository.GetProfileByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, result, domain.Profile{})
	mockMapper.AssertExpectations(t)
}
