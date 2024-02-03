package main

import (
	"database/sql"
	"example/gosterx/todo/handler"
	"example/gosterx/todo/repository"
	"example/gosterx/todo/router"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Error was raised during config reading: %s", err.Error())
	}
}

func main() {
	mysqlCfg := mysql.Config{
		User:      viper.GetString("database.user"),
		Passwd:    viper.GetString("database.password"),
		Net:       viper.GetString("database.net"),
		Addr:      fmt.Sprintf("%s:%d", viper.GetString("database.host"), viper.GetInt("database.port")),
		DBName:    viper.GetString("database.name"),
		ParseTime: viper.GetBool("database.parseTime"),
	}

	dbConn, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		log.Panicf("Error was raised during db connecting: %s", err.Error())
	}
	defer dbConn.Close()

	dbPingErr := dbConn.Ping()
	if dbPingErr != nil {
		log.Panicf("Error was raised during db ping: %s", err.Error())
	}

	r := router.New()

	v1 := r.Group("/api")

	// tr := repository.NewInMemoryRepository()
	dbTr := repository.NewDatabaseRepository(dbConn)

	h := handler.NewHandler(dbTr)

	h.Register(v1)

	r.Logger.Fatal(r.Start(viper.GetString("server.port")))
}
