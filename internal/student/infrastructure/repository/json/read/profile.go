package read

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/dto"
)

type Mapper interface {
	DTOProfileToDomain(email string, profile dto.Profile) domain.Profile
	DTOClassesToDomain(classes []dto.Class) []domain.Class
}

type ProfileRepositoryRead struct {
	mapper              Mapper
	filenameProfile     string
	filenameClassesDone string
}

func NewProfileRepositoryRead(mapper Mapper, filenameProfile string, filenameClassesDone string) *ProfileRepositoryRead {
	return &ProfileRepositoryRead{mapper: mapper, filenameProfile: filenameProfile, filenameClassesDone: filenameClassesDone}
}

func (r ProfileRepositoryRead) GetProfileByEmail(email string) (domain.Profile, error) {
	data, err := os.ReadFile(r.filenameProfile)
	if err != nil {
		return domain.Profile{}, err
	}

	profiles := make(map[string]dto.Profile)
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return domain.Profile{}, err
	}

	foundProfileDTO, found := profiles[email]
	if !found {
		return domain.Profile{}, fmt.Errorf("profile not found for email: %s", email)
	}

	foundProfile := r.mapper.DTOProfileToDomain(email, foundProfileDTO)

	return foundProfile, nil
}
