package main

import (
	log "github.com/sirupsen/logrus"
)

func getTotalPatientsByCountryId(cid int) (RespondTotalGlobalPatients, error){
	var data RespondTotalGlobalPatients
	DB.Select(`*`).Where("ct_id=?", cid).Find(&data)
	log.Infoln(data)
	return data, nil
}

func getCountryByCode(c string) (Country, error){
	var data Country
	DB.Where("code=?",c).Find(&data)
	log.Infoln(data)
	return data,nil
}

func getProvinceByProvinceEn(prove string) (Province, error){
	var data Province
	DB.Where("province_en=?",prove).Find(&data)
	log.Infoln(data)
	return data, nil
}

func getAllProvince() ([]Province, error){
	var data []Province
	DB.Find(&data)
	log.Infoln(data)
	return data, nil
}

func getReportThailand() ([]ReportPatientsInfo,error){
	var data []ReportPatientsInfo
	DB.Find(&data)
	log.Infoln(data)
	return data, nil
}


func GetTotalTop3() ([]RespondTotalTop3,error){
	var data []RespondTotalTop3

	DB.Raw(`(SELECT TOP 3 row_number() OVER (ORDER BY total_cases DESC) AS sequence, 
		c.ct_nameTH,
		c.ct_nameEN,
		tgp.total_cases,
		tgp.total_cases_increases,
		tgp.total_active_cases,
		tgp.total_recovered,
		tgp.total_deaths,
		tgp.update_date,
		c.emoji
    		FROM boibot.dbo.TOTAL_GLOBAL_PATIENTS tgp LEFT JOIN boibot.dbo.COUNTRY c ON tgp.ct_id = c.id)
		UNION
		(
    	SELECT sequence,
			ct_nameTH,
			ct_nameEN,
			total_cases,
			total_cases_increases,
			total_active_cases,
			total_recovered,
			total_deaths,
			update_date,
			emoji
    	FROM (
       		 SELECT row_number() OVER (ORDER BY total_cases DESC) AS sequence,
				c.id,c.ct_nameTH,
				c.ct_nameEN,
				tgp.total_cases,
				tgp.total_cases_increases,
				tgp.total_active_cases,
				tgp.total_recovered,
				tgp.total_deaths,
				tgp.update_date,
				c.emoji
			FROM boibot.dbo.TOTAL_GLOBAL_PATIENTS tgp LEFT JOIN boibot.dbo.COUNTRY c ON tgp.ct_id = c.id
		)AS global
   		 WHERE global.id = 219)`).Find(&data)
	return data, nil
}
