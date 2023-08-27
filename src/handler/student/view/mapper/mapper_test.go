package mapper

import (
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/view/contract"
	"github.com/stretchr/testify/assert"
)

const (
	testEmail = "test@example.com"
	testName  = "My Name"
)

func TestMapper_DomainToResponse(t *testing.T) {
	domainProfile := domain.NewProfile(testEmail, testName)
	domainDoneArray := []domain.Class{
		*domain.NewClass("id1", "Clase 1"),
		*domain.NewClass("id2", "Clase 2"),
	}
	responseClassesDone := []contract.Class{
		{ClassID: "id1", Title: "Clase 1"},
		{ClassID: "id2", Title: "Clase 2"},
	}
	mapper := Mapper{}
	expectedResponse := contract.Response{
		Email:       testEmail,
		Name:        testName,
		ClassesDone: responseClassesDone,
	}

	response := mapper.DomainToResponse(*domainProfile, domainDoneArray)

	assert.Equal(t, expectedResponse, response)
}
