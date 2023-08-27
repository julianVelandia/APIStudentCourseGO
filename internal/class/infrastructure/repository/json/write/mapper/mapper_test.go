package mapper_test

import (
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/write/mapper"
	"github.com/stretchr/testify/assert"
)

const (
	email   = "test@email.com"
	classID = "123"
	title   = "Updated Class"
)

func TestMapper_CommandToDTOClass(t *testing.T) {
	mapperWrite := mapper.Mapper{}
	updateCmd := command.NewUpdate(
		email,
		classID,
		title,
	)

	dtoClass := mapperWrite.CommandToDTOClass(*updateCmd)

	assert.Equal(t, classID, dtoClass.ClassID)
	assert.Equal(t, title, dtoClass.Title)
}
