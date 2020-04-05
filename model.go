//Struct Model

package main

import "time"

//--------------------------------------[ struct ]-------------------------------------------------------------------

type FACULTY struct {
	Id   int `gorm: "column:id;"`
	Name string `gorm: "column:name;" json:"name"`
}

func (FACULTY)TableName()string{
	return "FACULTY"
}

type RequestTotalThailandPatients struct {
	Confirmed int64 `json: "Confirmed"`
	Hospitalized int64 `json: "Hospitalized"`
	Recovered int64 `json: "Recovered"`
	Deaths int64 `json: "Deaths"`
	NewConfirmed int64 `json: "NewConfirmed"`
	NewHospitalized int64 `json: "NewHospitalized"`
	NewRecovered int64 `json: "NewRecovered"`
	NewDeaths int64 `json: "NewDeaths"`
	UpdateDate time.Time `json: "UpdateDate"`
}

type TotalGlobalPatients struct {
	CountryId int `gorm: "column:country_id;"`
	TotalCases int64 `gorm: "column:total_cases;"`
	TotalActiveCases int64 `gorm: "column:total_active_cases;"`
	TotalRecovered int64 `gorm: "column:total_recovered;"`
	TotalDeaths int64 `gorm: "column:total_deaths;"`
	TotalCasesIncreases int64 `gorm: "column:total_cases_increases;"`
	TotalActiveCasesIncreases int64 `gorm: "column:total_active_cases_increases;"`
	TotalRecoveredIncreases int64 `gorm: "column:total_recovered_increases;"`
	TotalDeathsIncreases int64 `gorm: "column:total_deaths_increases;"`
	UpdateDate time.Time `gorm: "column:update_date;"`
}

func (TotalGlobalPatients)TableName()string{
	return "TOTAL_GLOBAL_PATIENTS"
}

type Result struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Message string `json:"message"`
}