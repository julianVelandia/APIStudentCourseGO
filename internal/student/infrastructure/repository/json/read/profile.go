package read

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/dto"
)

const (
	filenameProfile = "jsonStudentsProfile.json"
)

type Mapper interface {
	DTOProfileToDomain(profile dto.Profile) domain.Profile
	DTOClassesToDomain(classes []dto.Class) []domain.Class
}

type ProfileRepositoryRead struct {
	mapper Mapper
}

func (r ProfileRepositoryRead) GetProfileByEmail(emailToFind string) (domain.Profile, error) {
	data, err := os.ReadFile(filenameProfile)
	if err != nil {
		return domain.Profile{}, err
	}

	profiles := make(map[string]dto.Profile)
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		return domain.Profile{}, err
	}

	foundProfileDTO, found := profiles[emailToFind]
	if !found {
		return domain.Profile{}, fmt.Errorf("profile not found for email: %s", emailToFind)
	}

	foundProfile := r.mapper.DTOProfileToDomain(foundProfileDTO)

	return foundProfile, nil
}
