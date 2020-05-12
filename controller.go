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
func UpdateTotalThailandCovidLogic(c *gin.Context) error{
	url := "https://covid19.th-stat.com/api/open/today"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	var getData RequestTotalThailandPatients       //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	country, _ := getCountryByCode("TH") //219
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
		return err
	}
	return nil
}

func UpdateThailandPatientInfoLogic(c *gin.Context) error{
	url := "https://covid19.th-stat.com/api/open/cases"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		return err
	}
	var getData RequestThailandPatientInfo      //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		return err
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
	return nil
}

func UpdateTotalGlobalCovidLogic(c *gin.Context) error{
	url := "https://api.thevirustracker.com/free-api?countryTotals=ALL"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		return err
	}
	var getData RequestTotalGlobalPatients      //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		return err
	}
	countryPartialInfo := getData.Countryitems[0]
	for i :=1 ; i <= 182 ; i++ {
		TotalCases := int64(field(countryPartialInfo, fmt.Sprintf("countryPartialInfo.Num%d.TotalCases",i)).Interface().(int))
		TotalActiveCases := int64(field(countryPartialInfo,  fmt.Sprintf("countryPartialInfo.Num%d.TotalActiveCases",i)).Interface().(int))
		TotalRecovered := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalRecovered",i)).Interface().(int))
		TotalDeaths := int64(field(countryPartialInfo,  fmt.Sprintf("countryPartialInfo.Num%d.TotalDeaths",i)).Interface().(int))
		TotalCasesIncreases := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalNewCasesToday",i)).Interface().(int))
		TotalDeathsIncreases := int64(field(countryPartialInfo, fmt.Sprintf( "countryPartialInfo.Num%d.TotalNewDeathsToday",i)).Interface().(int))
		country, _ := getCountryByCode(field(countryPartialInfo,fmt.Sprintf( "countryPartialInfo.Num%d.Code",i)).Interface().(string))


		dateTime := time.Now()
		y, m, d  := dateTime.Date()
		hh := dateTime.Hour()
		mm := dateTime.Minute()
		date := fmt.Sprintf("%d/%d/%d %d:%d", y, m, d, hh, mm)

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
	return nil
}

func UpdateTotalThailandPatientsProvinceLogic(c *gin.Context) error{
	url := "https://covid19.th-stat.com/api/open/cases/sum"
	dataAsByte, err := httpRequest(url)
	if err != nil {
		return err
	}
	var getData RequestTotalThailandPatientsProvince       //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		return err
	}
	ProvinceInfo , _:= getAllProvince()
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
	return nil
}

func GetTotalPatientsEndPointLogic(c *gin.Context) (RespondTotalGlobalPatients,error) {
	cid := 219 //TH
	data, err:= getTotalPatientsByCountryId(cid)
	if err != nil {
		return RespondTotalGlobalPatients{},err
	}
	return data, nil
}

func UpdateReportPatientsCovidLogic(c *gin.Context) error{
	url := "https://covid19.th-stat.com/api/open/timeline"
	dataAsByte, err := httpRequest(url) //byte array
	if err != nil {
		return err
	}

	var getData ReportPatients //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		return err
	}
	for i := 0; i < len(getData.Data); i++ {
		mapping := ReportPatientsInfo{
			UpdateDate:      getData.Data[i].Date,
			NewConfirmed:    getData.Data[i].NewConfirmed,
			NewRecovered:    getData.Data[i].NewRecovered,
			NewHospitalized: getData.Data[i].NewHospitalized,
			NewDeaths:       getData.Data[i].NewDeaths,
			Confirmed:       getData.Data[i].Confirmed,
			Recovered:       getData.Data[i].Recovered,
			Hospitalized:    getData.Data[i].Hospitalized,
			Deaths:          getData.Data[i].Deaths,
		}
		log.Infoln(mapping)
		//Update
		DB.Table("REPORT_PATIENTS_THAILAND").
			Where("update_date = ? ", getData.Data[i].Date).Update(&mapping)
	}
	checkdata := getReportThailand
	fmt.Print(checkdata)
	return nil
}



