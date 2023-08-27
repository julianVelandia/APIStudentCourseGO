package write

import (
	"encoding/json"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/command"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/infrastructure/repository/json/dto"
)

const (
	filenameClassesDone = "jsonStudentsClassesDone.json"
)

type Mapper interface {
	CommandToDTOClass(cmd command.Update) dto.Class
}

type ClassRepositoryWrite struct {
	mapper Mapper
}

func NewClassRepositoryWrite(mapper Mapper) *ClassRepositoryWrite {
	return &ClassRepositoryWrite{mapper: mapper}
}

func (r ClassRepositoryWrite) UpdateClassesByEmail(cmd command.Update) error {
	data, err := os.ReadFile(filenameClassesDone)
	if err != nil {
		return err
	}

	classesDoneByUser := make(map[string][]dto.Class)
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

	err = os.WriteFile(filenameClassesDone, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
