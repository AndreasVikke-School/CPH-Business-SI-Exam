package main

import (
	"net/http"
	"time"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	UserId   int64 `json:"userId,omitempty"`
	EntityId int64 `json:"entityId,omitempty"`
	UnixTime int64 `json:"unix,omitempty"`
}

// Create Log Entry
// @Schemes
// @Description Says Creates a log entry
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create_log_entry/ [post]
func CreateLogEntry(c *gin.Context) {
	var logEntry LogEntry
	err := c.BindJSON(&logEntry)
	eh.PanicOnError(err, "Couldn't bind JSON")
	logEntry.UnixTime = time.Now().UnixNano() / 1000000
	ProduceLogEntryToKafka(logEntry)
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
