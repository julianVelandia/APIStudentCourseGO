package write

import (
	"encoding/json"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/dto"
)

type Mapper interface {
	CommandToDTOClass(cmd command.Update) dto.ClassStudent
}

type ClassRepositoryWrite struct {
	mapper              Mapper
	filenameClassesDone string
}

func NewClassRepositoryWrite(mapper Mapper, filenameClassesDone string) *ClassRepositoryWrite {
	return &ClassRepositoryWrite{mapper: mapper, filenameClassesDone: filenameClassesDone}
}

func (r ClassRepositoryWrite) UpdateClassesByEmail(cmd command.Update) error {
	data, err := os.ReadFile(r.filenameClassesDone)
	if err != nil {
		return err
	}

	classesDoneByUser := make(map[string][]dto.ClassStudent)
	err = json.Unmarshal(data, &classesDoneByUser)
	if err != nil {
		return err
	}

	newClass := r.mapper.CommandToDTOClass(cmd)
	classesDoneByUser[cmd.Email()] = append(classesDoneByUser[cmd.Email()], newClass)

	updatedData, err := json.MarshalIndent(classesDoneByUser, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filenameClassesDone, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
