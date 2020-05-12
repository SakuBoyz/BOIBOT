// Function

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//--------------------------------------[ handle ]-------------------------------------------------------------------
func UpdateData(c *gin.Context)  {
	functionName := "UpdateData"
	err := UpdateTotalGlobalCovidLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=UpdateTotalGlobalCovidLogic; Update  Fail [ERROR]:%s", err),
		})
	}
	err = UpdateTotalThailandCovidLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=UpdateTotalGlobalCovidLogic; Update Fail [ERROR]:%s", err),
		})
	}
	//err = UpdateThailandPatientInfoLogic(c) TODO:new insert
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"message": fmt.Sprintf("[Function]=UpdateTotalGlobalCovidLogic; Update Fail [ERROR]:%s", err),
	//	})
	//}
	err = UpdateTotalThailandPatientsProvinceLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=UpdateTotalGlobalCovidLogic; Update Fail [ERROR]:%s", err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; Update success" , functionName),
	})
}

func UpdateTotalThailandCovid(c *gin.Context) {
	functionName := "UpdateTotalThailandCovid"
	err := UpdateTotalThailandCovidLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; Update success" , functionName),
	})
}

func UpdateThailandPatientInfo(c *gin.Context) {
	functionName := "UpdateThailandPatientInfo"
	err := UpdateThailandPatientInfoLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; Update success" , functionName),
	})
}

func UpdateTotalGlobalCovid(c *gin.Context) {
	functionName := "UpdateTotalGlobalCovid"
	err := UpdateTotalGlobalCovidLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func UpdateTotalThailandPatientsProvince(c *gin.Context) {
	functionName := "TotalThailandPatientsProvince"
	err := UpdateTotalThailandPatientsProvinceLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func GetTotalPatientsEndPoint(c *gin.Context)  {
	functionName := "GetTotalPatientsEndPoint"
	data, err := GetTotalPatientsEndPointLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": data,
	})
}

func UpdateReportPatientsCovid(c *gin.Context){
	functionName := "UpdateReportPatientsCovid"
	err := UpdateReportPatientsCovidLogic(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Update Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func GetGlobalTop3(c *gin.Context) {
	functionName := "GetGlobalTop3"
	data, err := GetTotalTop3()
	log.Infoln(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; GET Fail [ERROR]:%s" , functionName, err),
		})
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
