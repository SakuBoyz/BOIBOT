// Function

package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

//--------------------------------------[ handle ]-------------------------------------------------------------------
func UpdateTotalThailandCovid(c *gin.Context) {
	functionName := "UpdateTotalThailandCovid"
	url := "https://covid19.th-stat.com/api/open/today"
	dataAsByte, err := httpRequest(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; QUERY fail", functionName),
		})
	}
	var getData RequestTotalThailandPatients //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Unmarshal fail", functionName),
		})
	}
	mapping := TotalGlobalPatients{
		CountryId:                 156,
		TotalCases:                getData.Confirmed,
		TotalActiveCases:          getData.Hospitalized,
		TotalRecovered:            getData.Recovered,
		TotalDeaths:               getData.Deaths,
		TotalCasesIncreases:       getData.NewConfirmed,
		TotalActiveCasesIncreases: getData.NewHospitalized,
		TotalRecoveredIncreases:   getData.NewRecovered,
		TotalDeathsIncreases:      getData.NewDeaths,
		UpdateDate:                getData.UpdateDate,
	}
	// Create
	DB.Update(&mapping)
	c.JSON(200, gin.H{
		"message": "UPDATE Success",
	})
}

func getTotalPatientsEndPoint(c *gin.Context) {
	cid := 219
	data := getTotalPatientsByCountryId(cid)
	c.JSON(200, gin.H{
		"message": data,
	})
}

func exampleFunc(c *gin.Context) {
	var result Result                //สร้างตัวแปร result ด้วย struct Result
	result.ID = c.Query("id")        //อ่านจาก parameter params
	result.Name = c.PostForm("name") // อ่านจาก body form-data
	result.Message = c.PostForm("message")

	c.JSON(200, gin.H{
		"message": result,
	})
}

func exampleJSON(c *gin.Context) {
	var input Result          //สร้างตัวแปร input ด้วย struct Result
	err := c.BindJSON(&input) //อ่าน parameter JSON แล้วจัดรูปแบบตาม input
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"id":      input.ID,
		"name":    input.Name,
		"message": input.Message,
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func createFaculty(c *gin.Context) {

	var input FACULTY

	err := c.ShouldBindJSON(&input) //อ่าน parameter JSON แล้วจัดรูปแบบตาม input
	if err != nil {
		fmt.Println(err)
	}

	// Create
	fmt.Println(input.Name)
	DB.Create(&input)
	c.JSON(201, gin.H{
		"message": "Create Success",
	})
}

func getAllFaculty(c *gin.Context) {
	// Read
	var data []FACULTY
	DB.Select(`*`).Find(&data)
	fmt.Println(data)

	c.JSON(200, gin.H{
		"message": data,
	})
}

func updateFacultyById(c *gin.Context) {
	// Update
	var input FACULTY

	err := c.ShouldBindJSON(&input) //อ่าน parameter JSON แล้วจัดรูปแบบตาม input
	if err != nil {
		fmt.Println(err)
	}

	DB.Find(&input).Update("Name", &input)
	c.JSON(200, gin.H{
		"message": "Update Success",
	})

}
