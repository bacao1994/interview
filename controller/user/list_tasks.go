package user

import (
	"github.com/gin-gonic/gin"
	"manabie/interview/entity"
	"manabie/interview/query"
	"manabie/interview/util"
	"net/http"
)

func ListTasks(c *gin.Context) {
	id := util.TokenUserID(c)

	task, err := query.TaskByUserID(id)
	if err != nil {
		task = []entity.Task{}
	}
	c.JSON(http.StatusOK, task)
}