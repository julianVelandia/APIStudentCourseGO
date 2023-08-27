package mapper

import (
	"fmt"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/application/query"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/internal/class/domain"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/platform/constant"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/platform/mysql"
	contract2 "github.com/julianVelandia/EDteam/SOLIDyHexagonal/ProyectoCurso/src/handler/course/list/contract"
	"strconv"
	"strings"
)

type HandlerMapper struct{}

func (hm HandlerMapper) RequestToQuery(request contract2.URLParams) (query.List, error) {
	intPage, errPage := strconv.Atoi(request.Page)
	if errPage != nil {
		return query.List{}, errPage
	}

	sort := request.Sort
	if len(request.Sort) < 1 {
		sort = mysql.Random.Value()
	}

	return *query.NewListDefault(
		intPage,
		request.Tags,
		sort,
		constant.DefaultBatchCourse.Value(),
		1,
		"",
	), nil
}
func (hm HandlerMapper) EntityToResponse(entities []domain.Course) contract2.Response {

	courses := make([]contract2.Course, len(entities))
	for i := range courses {
		tags := make([]string, 0)
		for j := range entities[i].Tags {
			tags = append(tags, entities[i].Tags[j].Tag)
		}
		var content string

		if len(strings.Split(entities[i].Content[1], " ")) > 15 {
			content = fmt.Sprintf("%v...", strings.Join(strings.Split(entities[i].Content[1], " ")[:15], " "))
		} else {
			content = fmt.Sprintf("%v...", entities[i].Content[1])
		}

		courses[i] = contract2.Course{
			CourseID: entities[i].CourseID,
			Title:    entities[i].Title,
			Content:  content,
			//Views:         entities[i].Views,
			//Likes:         entities[i].Likes,
			//Dislikes:      entities[i].Dislikes,
			//NumberAnswers: len(entities[i].Answers),
			Tags: tags,
		}
	}
	return contract2.Response{
		Courses: courses,
	}
}
