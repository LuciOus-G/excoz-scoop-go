package main

import (
	"net/http"
	"strings"

	"text/template"

	utils "excoz.scoop/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit(Conf utils.Conf) {
	host := Conf.DBHost
	password := "admin"
	user := "postgres"
	dbName := "eshier_db"
	port := "5432"

	sentence := "host={{.hostt}} user={{.usert}} password={{.passwordt}} dbname={{.dbnamet}} port={{.portt}} sslmode=disable"
	t, b := new(template.Template), new(strings.Builder)
	template.Must(t.Parse(sentence)).Execute(b, map[string]interface{}{
		"hostt":     host,
		"usert":     user,
		"passwordt": password,
		"dbnamet":   dbName,
		"portt":     port})

	database, err = gorm.Open(postgres.Open(b.String()), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}

func main() {
	conf := utils.Config()
	e := echo.New()

	// start database
	DatabaseInit(conf)
	gorm := DB()
	_, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, conf.DBHost)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
