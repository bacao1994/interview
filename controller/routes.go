package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"manabie/interview/controller/healthy"
	"manabie/interview/controller/user"
	"manabie/interview/global"
	_ "manabie/interview/swagger"
	"net/http"
)

func Initialize() {
	var port = global.Config.ServerPort

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	healthy.RegisterRoutes(router)
	user.RegisterRoutes(router)
	// enable swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	fmt.Println("Interview backend API listening on port: " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
