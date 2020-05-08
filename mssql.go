// Connect SQL

package main


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	viper.AddConfigPath("./configs")   // path to look for the config file in
	if err = viper.ReadInConfig();
	err != nil {
		log.Errorln("Fatal Error Config File: ",err)
		panic("Fatal Error Config File")
	}
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		viper.GetString("mssql.server"),
		viper.GetString("mssql.user"),
		viper.GetString("mssql.password"),
		viper.GetString("mssql.port"),
		viper.GetString("mssql.database"))
	log.Infoln(connectionString)
	db, err := gorm.Open(viper.GetString("mssql.databaseType"), connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	// Update - update product's price to 2000
	//db.Model(&data).Update("Price", 2000)

	// Delete - delete product
	//db.Delete(&data)

	DB = db
	return DB
}