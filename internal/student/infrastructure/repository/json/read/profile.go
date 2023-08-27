package read

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/infrastructure/repository/json/dto"
)

const (
	filenameProfile = "dbtest/StudentsProfile.json"
)

type Mapper interface {
	DTOProfileToDomain(email string, profile dto.Profile) domain.Profile
	DTOClassesToDomain(classes []dto.Class) []domain.Class
}

type ProfileRepositoryRead struct {
	mapper Mapper
}

func NewProfileRepositoryRead(mapper Mapper) *ProfileRepositoryRead {
	return &ProfileRepositoryRead{mapper: mapper}
}

func (r ProfileRepositoryRead) GetProfileByEmail(email string) (domain.Profile, error) {
	data, err := os.ReadFile(filenameProfile)
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
