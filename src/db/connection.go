package db

import (
	"github.com/gocarina/gocsv"
	"os"
	"test_task_makves/src/model"
)

// todo: заменить путь на относительный, докрутить кроссплатформу

func ReadLocalDB() ([]model.DB_entity, error) {

	entityArray := []model.DB_entity{}

	path := ""
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	file, err := os.Open(path)
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
