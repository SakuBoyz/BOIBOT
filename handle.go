// Function

package main

import (
	"encoding/json"
	"fmt"
	structs "github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

//--------------------------------------[ handle ]-------------------------------------------------------------------
func UpdateTotalThailandCovid(c *gin.Context) {
	functionName := "UpdateTotalThailandCovid"
	url := "https://covid19.th-stat.com/api/open/today"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; QUERY fail " + err.Error(), functionName),
		})
		return
	}
	var getData RequestTotalThailandPatients       //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Unmarshal fail " + err.Error(), functionName),
		})
		return
	}
	country := getCountryByCode("TH") //219
	mapping := UpdateTotalGlobalPatients {
		TotalCases : getData.Confirmed,
		TotalActiveCases : getData.Hospitalized,
		TotalRecovered : getData.Recovered,
		TotalDeaths : getData.Deaths,
		TotalCasesIncreases : getData.NewConfirmed,
		TotalActiveCasesIncreases : getData.NewHospitalized,
		TotalRecoveredIncreases : getData.NewRecovered,
		TotalDeathsIncreases : getData.NewDeaths,
		UpdateDate : getData.UpdateDate,
	}
	// Update
	UpdateData := DB.Table("TOTAL_GLOBAL_PATIENTS").
		Where("ct_id = ? ", country.Id).
		Update(&mapping)
	if UpdateData == nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; UpdateData fail " , functionName),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func UpdateThailandPatientInfo(c *gin.Context) {
	functionName := "UpdateThailandPatientInfo"
	url := "https://covid19.th-stat.com/api/open/cases"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; QUERY fail" + err.Error(), functionName),
		})
	}
	var getData RequestThailandPatientInfo      //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Unmarshal fail" + err.Error(), functionName),
		})
	}

	for i :=0 ; i < len(getData.Data) ; i++ {
		mapping := ThailandPatientInfo{
			//Id:     		&getData.Data[i].No,
			ConfirmDate:    &getData.Data[i].ConfirmDate,
			Age:       		&getData.Data[i].Age,
			GenderTh:       &getData.Data[i].Gender,
			GenderEn:  		&getData.Data[i].GenderEn,
			NationalityTh: 	&getData.Data[i].Nation,
			NationalityEn: 	&getData.Data[i].NationEn,
			District: 		&getData.Data[i].District,
			ProvinceId:  	&getData.Data[i].ProvinceId,
		}
		log.Infoln(mapping)
		// CREATE
		DB.Create(&mapping)
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func UpdateTotalGlobalCovid(c *gin.Context) {
	functionName := "UpdateTotalGlobalCovid"
	url := "https://api.thevirustracker.com/free-api?countryTotals=ALL"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; QUERY fail" + err.Error(), functionName),
		})
	}
	var getData RequestTotalGlobalPatients      //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Unmarshal fail" + err.Error(), functionName),
		})
	}
	countryPartialInfo := getData.Countryitems[0]
	for i :=1 ; i <= 182 ; i++ {
		TotalCases := int64(field(countryPartialInfo, fmt.Sprintf("countryPartialInfo.Num%d.TotalCases",i)).Interface().(int))
		TotalActiveCases := int64(field(countryPartialInfo,  fmt.Sprintf("countryPartialInfo.Num%d.TotalActiveCases",i)).Interface().(int))
		TotalRecovered := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalRecovered",i)).Interface().(int))
		TotalDeaths := int64(field(countryPartialInfo,  fmt.Sprintf("countryPartialInfo.Num%d.TotalDeaths",i)).Interface().(int))
		TotalCasesIncreases := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalNewCasesToday",i)).Interface().(int))
		TotalDeathsIncreases := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalNewDeathsToday",i)).Interface().(int))
		country := getCountryByCode(field(countryPartialInfo,fmt.Sprintf( "countryPartialInfo.Num%d.Code",i)).Interface().(string))


		dateTime := time.Now()
			y, m, d  := dateTime.Date()
		hh := dateTime.Hour()
		mm := dateTime.Minute()
		date := fmt.Sprintf("%d/%d/%d %d:%d", d, m, y, hh, mm)

		mapping := UpdateTotalGlobalPatients{
			//CountryId:			  country.Id,  //for Create
			TotalCases:           &TotalCases,
			TotalActiveCases:     &TotalActiveCases,
			TotalRecovered:       &TotalRecovered,
			TotalDeaths:          &TotalDeaths,
			TotalCasesIncreases:  &TotalCasesIncreases,
			TotalDeathsIncreases: &TotalDeathsIncreases,
			UpdateDate: &date,
		}
		log.Infoln(mapping)
		// Update
		DB.Table("TOTAL_GLOBAL_PATIENTS").
			Where("ct_id = ? ", country.Id).
			Update(&mapping)

		// Create
		//DB.Table("TOTAL_GLOBAL_PATIENTS").Create(&mapping)
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("[Function]=%s; success" , functionName),
	})
}

func UpdateTotalThailandPatientsProvince(c *gin.Context) {
	functionName := "TotalThailandPatientsProvince"
	url := "https://covid19.th-stat.com/api/open/cases/sum"
	dataAsByte, err := httpRequest(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; QUERY fail", functionName),
		})
	}
	var getData RequestTotalThailandPatientsProvince       //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("[Function]=%s; Unmarshal fail", functionName),
		})
	}
	ProvinceInfo := getAllProvince()
	s := structs.New(getData.Province)
	for _, v := range ProvinceInfo {
		fmt.Printf("getData.Province.%s\n", *v.ProvinceEn)
		replaceSpace := strings.ReplaceAll(*v.ProvinceEn, " ", "")
		findField, ok := s.FieldOk(replaceSpace)
		if ok {
			ValueInMap, _ := strconv.Atoi(fmt.Sprintf("%+v", findField.Value()))
			TotalCase := int64(ValueInMap)
			mapping := TotalThailandPatientsProvince{
				TotalCase:  &TotalCase,
			}
			// Update
			DB.Table("TOTAL_THAILAND_PATIENTS_PROVINCE").
				Where("province_id=?", *v.Id).
				Update(&mapping)
		}
	}
	c.JSON(200, gin.H{
		"message": "Update Success",
	})
}


func GetTotalPatientsEndPoint(c *gin.Context)  {
	cid := 219
	data := getTotalPatientsByCountryId(cid)
	c.JSON(200, gin.H{
		"message": data,
	})
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}