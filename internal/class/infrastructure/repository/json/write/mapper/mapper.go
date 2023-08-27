package mapper

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/dto"
)

type Mapper struct{}

func (m Mapper) CommandToDTOClass(cmd command.Update) dto.ClassStudent {
	return dto.ClassStudent{
		ClassID: cmd.ClassID(),
		Title:   cmd.Title(),
	}
}
