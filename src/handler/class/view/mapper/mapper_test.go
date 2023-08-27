package mapper

import (
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/class/view/contract"
	"github.com/stretchr/testify/assert"
)

func TestDomainToResponse(t *testing.T) {
	domaines := []domain.Class{
		{ClassID: "id1", Title: "Clase 1"},
		{ClassID: "id2", Title: "Clase 2"},
	}
	mapper := Mapper{}
	expectedResponse := contract.Response{
		Classes: []contract.Class{
			{ClassID: "id1", Title: "Clase 1"},
			{ClassID: "id2", Title: "Clase 2"},
		},
	}

	response := mapper.DomainToResponse(domaines)

	assert.Equal(t, expectedResponse, response)
}
