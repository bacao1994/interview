package controller_tests

import (
	"github.com/google/uuid"
	"log"
	"manabie/interview/entity"
	"manabie/interview/global"
	pkg_rd "manabie/interview/pkg/rd"
	"manabie/interview/util"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	global.FetchTestEnvironmentVariables()
	entity.InitializeDb()

	pkg_rd.InitializeRdV8()

	//os.Exit(m.Run())
	log.Printf("Before calling m.Run() !!!")
	ret := m.Run()
	log.Printf("After calling m.Run() !!!")
	//os.Exit(m.Run())
	os.Exit(ret)

}

func refreshUserTable() error {
	err := entity.Db().Migrator().DropTable(entity.User{})
	if err != nil {
		return err
	}
	err = entity.Db().AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (entity.User, error) {

	_ = refreshUserTable()

	user := entity.User{
		ID:       "firstUser",
		Password: util.HashPassword("firstUser", "password"),
		MaxTodo:  5,
	}

	err := user.Create()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {
	users := []entity.User{
		{
			ID:       "firstUser",
			Password: util.HashPassword("firstUser", "password"),
			MaxTodo:  3,
		},
		{
			ID:       "secondUser",
			Password: util.HashPassword("secondUser", "password"),
			MaxTodo:  5,
		},
	}

	for i, _ := range users {
		err := users[i].Create()
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndTaskTable() error {

	err := entity.Db().Migrator().DropTable(entity.User{}, entity.Task{})
	if err != nil {
		return err
	}
	err = entity.Db().AutoMigrate(entity.User{}, entity.Task{})
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOneTask() (entity.User, entity.Task, error) {
	err := refreshUserAndTaskTable()
	if err != nil {
		return entity.User{}, entity.Task{}, err
	}
	user := entity.User{
		ID:       "firstUser",
		Password: util.HashPassword("firstUser", "password"),
		MaxTodo:  3,
	}
	err = user.Create()
	if err != nil {
		return entity.User{}, entity.Task{}, err
	}
	task := entity.Task{
		ID:          uuid.New().String(),
		Content:     "content",
		UserID:      user.ID,
		CreatedDate: time.Now().Format(util.DefaultTimeFormat),
	}
	err = task.Create()
	if err != nil {
		return entity.User{}, entity.Task{}, err
	}
	return user, task, nil
}
