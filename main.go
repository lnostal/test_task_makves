package main

import (
	"github.com/gin-gonic/gin"
	"test_task_makves/src/api"
)

func main() {

	var router = gin.Default()

	api.API(router)
	//db.ReadLocalDB()

	router.Run()

}
