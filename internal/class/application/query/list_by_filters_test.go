package query_test

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Batch(t *testing.T) {
	list := query.NewListDefault(1, "tag1", "asc", 10)
	assert.Equal(t, 10, list.Batch(), "Batch should return the correct value")
}

func TestList_OffSet(t *testing.T) {
	list := query.NewListDefault(2, "tag2", "desc", 20)
	assert.Equal(t, 20, list.OffSet(), "OffSet should return the correct value")
}

func TestList_Tag(t *testing.T) {
	list := query.NewListDefault(3, "tag3", "asc", 5)
	assert.Equal(t, "tag3", list.Tag(), "Tag should return the correct value")
}

func TestList_Sort(t *testing.T) {
	list := query.NewListDefault(4, "tag4", "desc", 15)
	assert.Equal(t, "desc", list.Sort(), "Sort should return the correct value")
}
