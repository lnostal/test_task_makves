package db

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"test_task_makves/src/dto"
)

// todo: заменить путь на относительный, докрутить кроссплатформу

func ReadLocalDB() []*model.DB_entity {

	entityArray := []*model.DB_entity{}

	file, err := os.Open("/Users/nosta/Documents/projects/test_task_makves/src/db/table_test.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = gocsv.Unmarshal(file, &entityArray)
	if err != nil {
		return nil
	}

	return entityArray
}
