package api

import (
	"github.com/gin-gonic/gin"
	v1 "test_task_makves/src/api/v1"
)

const ActualAPI string = "/api/v1/"

func API(router *gin.Engine) {

	router.GET(ActualAPI+"/get-items", v1.GetItemsByIds)
}
