package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
  	r.Use(gin.Logger())
	InitDB()

	//
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT", "POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

	//--------------------------------------[ router ]-------------------------------------------------------------------
	r.GET("/", hello)
	r.POST("/test", exampleFunc)
	r.POST("/testJSON", exampleJSON)
	r.POST("/callback", callbackHandler)
	r.GET("/updateTotalThailandCovid", UpdateTotalThailandCovid)
	r.GET("/getTotalGlobalPatients", getTotalPatientsEndPoint)
	r.GET("/getAllFaculty", getAllFaculty)
	r.POST("/createFaculty", createFaculty)
	r.POST("/updateFacultyById", updateFacultyById)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
