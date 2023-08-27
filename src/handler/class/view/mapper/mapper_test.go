package mapper

import (
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/student/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/student/profile/contract"
	"github.com/stretchr/testify/assert"
)

const (
	testEmail = "test@example.com"
	testName  = "My Name"
)

func TestDomainToResponse(t *testing.T) {
	domainProfile := domain.NewProfile(testEmail, testName)
	mapper := Mapper{}
	expectedResponse := contract.Response{
		Email: testEmail,
		Name:  testName,
	}

	response := mapper.DomainToResponse(*domainProfile)

	assert.Equal(t, expectedResponse, response)
}
