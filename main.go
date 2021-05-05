package main

import (
	"manabie/interview/controller"
	"manabie/interview/entity"
	"manabie/interview/global"
	pkg_logrus "manabie/interview/pkg/logger"
	pkg_rd "manabie/interview/pkg/rd"
)

func main() {
	global.FetchEnvironmentVariables()

	entity.InitializeDb()

	pkg_logrus.InitLogrus()

	pkg_rd.InitializeRdV8()

	controller.Initialize()
}
