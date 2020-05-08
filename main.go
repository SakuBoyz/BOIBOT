package main
import (
       	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	InitDB()

	//
	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT", "POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

//--------------------------------------[ router ]-------------------------------------------------------------------
	//GET
	r.GET("/", hello)
	r.GET("/GetTotalGlobalPatients", GetTotalPatientsEndPoint)
	//POST
	r.POST("/callback", callbackHandler)
	r.POST("/UpdateTotalThailandCovid", UpdateTotalThailandCovid)
	r.POST("/UpdateTotalGlobalCovid", UpdateTotalGlobalCovid)
	r.POST("/UpdateThailandPatientInfo", UpdateThailandPatientInfo)
	r.POST("/UpdateTotalThailandPatientsProvince", UpdateTotalThailandPatientsProvince)
	r.Run(":1443") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}






