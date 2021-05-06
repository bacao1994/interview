package entity_tests

import (
	"github.com/go-playground/assert/v2"
	"log"
	"manabie/interview/entity"
	"manabie/interview/util"
	"testing"
)

func TestSaveUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}
	newUser := entity.User{
		ID:       "firstUser",
		Password: util.HashPassword("firstUser", "password"),
		MaxTodo:  5,
	}
	err = newUser.Create()
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, "firstUser")
	assert.Equal(t, newUser.MaxTodo,5)
}
