package query

import (
	"manabie/interview/entity"
	"strings"
)

func TaskByUserID(userId string) (result []entity.Task, err error) {
	err = entity.Db().Where("user_id = ?", strings.TrimSpace(userId)).Find(&result).Error
	return result, err
}
