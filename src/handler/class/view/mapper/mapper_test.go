package mapper

import (
	"testing"

	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/stretchr/testify/assert"
)

const (
	classID      = "123"
	title        = "Test Class"
	creationDate = "2023-08-27"
	readTime     = 1.5
)

func TestMapper_DomainToResponse(t *testing.T) {
	mapper := Mapper{}
	content := []string{"Content line 1", "Content line 2"}
	domainClass := domain.NewClass(classID, title, creationDate, content, readTime)

	response := mapper.DomainToResponse(*domainClass)

	assert.Equal(t, classID, response.ClassID)
	assert.Equal(t, title, response.Title)
	assert.Equal(t, creationDate, response.CreationDate)
	assert.Equal(t, content, response.Content)
	assert.Equal(t, readTime, response.ReadTime)
}