package main

import (
	log "github.com/sirupsen/logrus"
)

func getTotalPatientsByCountryId(cid int) TotalGlobalPatients{
	var data TotalGlobalPatients
	DB.Select(`*`).Where("country_id=?", cid).Find(&data)
	log.Infoln(data)
	return data
}