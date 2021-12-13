package main

import (
	"net/http"

	eh "github.com/andreasvikke-school/CPH-Bussiness-SI-Exam/applications/services/api/errorhandler"
	"github.com/gin-gonic/gin"
)

type BorrowerEntry struct {
	BorrowerId int64 `json:"borrowerId,omitempty"`
	EntityId   int64 `json:"entityId,omitempty"`
}

func CreateBorrower(c *gin.Context) {
	var borrowerEntry BorrowerEntry
	err := c.BindJSON(&borrowerEntry)
	eh.PanicOnError(err, "Couldn't bind JSON")
	ProduceBorrowerEntryToRabbitmq(borrowerEntry)
	c.IndentedJSON(http.StatusOK, gin.H{
		"queued": "success",
	})
}
