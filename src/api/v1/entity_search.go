package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"test_task_makves/src/db"
	"test_task_makves/src/model"
)

func convertQuery(s string) []int {
	var ids []int

	values := strings.Split(s, ",")
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		} else {
			ids = append(ids, num)
		}

	}

	return ids
}

func findEntities(ids []int) []model.DB_entity {
	var entities []model.DB_entity

	data, err := db.ReadLocalDB()

	if err != nil {
		return entities
	}

	for _, id := range ids {
		for _, entity := range data {
			if id == entity.Id {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

func GetItemsByIds(c *gin.Context) {
	query := c.Request.URL.Query()
	valuesStr := query.Get("id")

	ids := convertQuery(valuesStr)

	entities := findEntities(ids)

	if entities == nil {
		c.IndentedJSON(http.StatusNotFound, entities)
		return
	}

	c.IndentedJSON(http.StatusOK, entities)
}
