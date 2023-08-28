package read

import (
	"encoding/json"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/dto"
)

func (r ProfileRepositoryRead) GetClassesDoneByEmail(email string) ([]domain.Class, error) {
	data, err := os.ReadFile(r.filenameClassesDone)
	if err != nil {
		return []domain.Class{}, err
	}

	classesDoneByUser := make(map[string][]dto.Class)
	err = json.Unmarshal(data, &classesDoneByUser)
	if err != nil {
		return []domain.Class{}, err
	}

	classesDTO, found := classesDoneByUser[email]
	if !found {
		return []domain.Class{}, nil
	}

	classesDone := r.mapper.DTOClassesToDomain(classesDTO)

	return classesDone, nil
}
