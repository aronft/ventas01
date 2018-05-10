package configuration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//Configuration es la estructura para poblar la configuracion
type Configuration struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//GetConfiguration obtiene y carga la configuration en Configuration
func GetConfiguration() Configuration {
	var c Configuration

	file, err := os.Open("./configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

//GetConnection obtiene un conexion a la base de datos
func GetConnection() *sql.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
	// dsn := "postgres://postgres:aster123@127.0.0.1:5432/ventas01?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
