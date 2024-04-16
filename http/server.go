package http

import (
	"daltondiaz/async-jobs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(){
    router := gin.Default()
    router.POST("/new/job", newJob)
    router.Run("localhost:8080")
}

func newJob(gin *gin.Context){
    var job models.Job
    if err := gin.BindJSON(&job); err != nil {
        return
    }
    gin.IndentedJSON(http.StatusCreated, job)
}
