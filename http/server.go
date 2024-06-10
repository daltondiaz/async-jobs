package http

import (
	"daltondiaz/async-jobs/db"
	"daltondiaz/async-jobs/models"
	"daltondiaz/async-jobs/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.POST("/job/new", newJob)
	router.GET("/job/stop/:id", stopJob)
	router.GET("/job/status/:id", status)
	router.GET("/job/enabled/:id/:enabled", enabled)
	router.GET("/health", health)
	router.Run("localhost:8080")
}

// Stop the cron job of on job
func enabled(c *gin.Context) {
	paramId := c.Params.ByName("id")
	paramEnabled := c.Params.ByName("enabled")
	id, _ := strconv.Atoi(paramId)
	enabled, _ := strconv.ParseBool(paramEnabled)
	status, err := pkg.EnabledJob(int64(id), enabled)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Not exists or is unabled",
			"id":      paramId,
			"enabled": paramEnabled,
		})
		return
	}
	result := gin.H{
		"status":  status,
		"id":      paramId,
		"enabled": paramEnabled,
	}
	c.IndentedJSON(http.StatusCreated, result)
}

// Get the Status of the cron job, if is running or not
func status(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	status, err := pkg.StatusJob(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Not not exists or is unabled",
			"id":      paramId,
		})
		return
	}
	result := gin.H{
		"status": status,
		"id":     paramId,
	}
	c.IndentedJSON(http.StatusCreated, result)
}

// Stop the cron job of on job
func stopJob(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	job, err := db.LoadJob(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Not not exists or is unabled",
			"id":      paramId,
		})
	}
	pkg.StopJob(job)
	c.IndentedJSON(http.StatusCreated, job)
}

// Create new Job and add to execute in Cron
func newJob(c *gin.Context) {
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		return
	}
	job, err := db.InsertJob(job)

	// no try/catchs this is awesome for devs
	// who works with Java/Php and other languages who use that
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Job not created"})
	}
	pkg.AddJobNextExecution(job)
	c.IndentedJSON(http.StatusCreated, job)
}

// Only to check if server is running
func health(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, gin.H{"message": "Server is running"})
}
