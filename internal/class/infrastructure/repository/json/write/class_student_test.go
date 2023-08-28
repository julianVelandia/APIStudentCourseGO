package write_test

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/write"
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/dto"
	"github.com/stretchr/testify/assert"
)

const (
	email         = "test2@email.com"
	title         = "title Class 1"
	classID       = "classid1"
	name          = "Julián"
	filenameClass = "../../../../../../dbtest/StudentsClassesDone.json"
)

func TestUpdateClassesByEmailWhenSuccessShouldReturnProfile(t *testing.T) {
	mockMapper := new(MapperMock)
	updateCmd := *command.NewUpdate(
		email,
		classID,
		title,
	)
	classDTO := dto.ClassStudent{
		Title: title,
	}

	mockMapper.On("CommandToDTOClass", updateCmd).Return(classDTO)
	repository := write.NewClassRepositoryWrite(mockMapper, filenameClass)
	err := repository.UpdateClassesByEmail(updateCmd)

	assert.NoError(t, err)
	mockMapper.AssertExpectations(t)
}

func TestUpdateClassesByEmailWhenReadFailShouldReturnError(t *testing.T) {
	mockMapper := new(MapperMock)

	repository := write.NewClassRepositoryWrite(nil, "")
	err := repository.UpdateClassesByEmail(command.Update{})

	assert.Error(t, err)
	mockMapper.AssertExpectations(t)
}
