package main

import (
	log "github.com/sirupsen/logrus"
)

func getTotalPatientsByCountryId(cid int) RespondTotalGlobalPatients{
	var data RespondTotalGlobalPatients
	DB.Select(`*`).Where("country_id=?", cid).Find(&data)
	log.Infoln(data)
	return data
}

func getCountryByCode(c string) Country{
	var data Country
	DB.Where("code=?",c).Find(&data)
	log.Infoln(data)
	return data
}

func getProvinceByProvinceEn(prove string) Province{
	var data Province
	DB.Where("province_en=?",prove).Find(&data)
	log.Infoln(data)
	return data
}

func getAllProvince() []Province{
	var data []Province
	DB.Find(&data)
	log.Infoln(data)
	return data
}
