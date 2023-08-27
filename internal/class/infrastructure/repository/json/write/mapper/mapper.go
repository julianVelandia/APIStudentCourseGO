package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/dto"
)

type Mapper struct{}

func (m Mapper) CommandToDTOClass(cmd command.Update) dto.Class {
	return dto.Class{
		ClassID: cmd.ClassID(),
		Title:   cmd.Title(),
	}
}
