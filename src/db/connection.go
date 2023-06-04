package db

import (
	"github.com/gocarina/gocsv"
	"os"
	"test_task_makves/src/model"
)

// todo: заменить путь на относительный, докрутить кроссплатформу

func ReadLocalDB() ([]model.DB_entity, error) {

	entityArray := []model.DB_entity{}

	file, err := os.Open("/Users/nosta/Documents/projects/test_task_makves/src/db/table.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = gocsv.Unmarshal(file, &entityArray)
	if err != nil {
		return nil, err
	}

	return entityArray, nil
}
