package server

import (
	"net/http"

	"github.com/ahmedsameha1/ccjsonparser/internal/app"
	"github.com/gin-gonic/gin"
)

const CONTENT_TYPE_HEADER = "Content-Type"
const APPLICATION_JSON = "application/json"

type Server struct {
	*gin.Engine
}

func New() *Server {
	engine := gin.Default()
	engine.POST("/", Handler)
	return &Server{Engine: engine}
}

func Handler(context *gin.Context) {
	contentTypeHeader := context.GetHeader(CONTENT_TYPE_HEADER)
	if contentTypeHeader != APPLICATION_JSON {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrNoApplicationJsonHeader.Error()})
		return
	}
	var json string
	err := context.ShouldBindJSON(&json)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrWhileReadingBody.Error()})
		return
	}
	if len(json) > 1024 {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrBodyIsTooLong.Error()})
		return
	}
	err = app.Validate(json)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"result": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"result": "This is a valid JSON"})
}
