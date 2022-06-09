package config

import (
	"fmt"

	goconfig "github.com/iglin/go-config"
)

type Configurations struct {
	Database DBConfiguration
	Server   ServerConfigurations
}

type DBConfiguration struct {
	DBHost  string
	DBUserR string
	DBPassR string
	DBUserW string
	DBPassW string
	Schema  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

func BuildDataSource() (string, string) {

	config := goconfig.NewConfig("../../config/development.yaml", goconfig.Yaml)

	// reading strings
	port := config.GetString("server.port", ":8080")
	host := config.GetString("database.dbhost")
	dbUserR := config.GetString("database.dbuserr")
	dbPassR := config.GetString("database.dbpassr")
	schema := config.GetString("database.schema")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbUserR, dbPassR, host, schema)

	return dataSource, port
}
