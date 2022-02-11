package storage

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

//Driver of storage
type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

//New create a connection with DB
func New(d Driver) {
	loadEnv()
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newPostgresDB singleton
func newPostgresDB() {

	dbUser := os.Getenv("PSQL_DB_USERNAME")
	dbPassword := os.Getenv("PSQL_DB_PASSWORD")
	dbDatabase := os.Getenv("PSQL_DB_DATABASE")
	url := "postgres://" + dbUser + ":" + dbPassword + "@localhost:5432/" + dbDatabase + "?sslmode=disable"
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

// newMySQLDB singleton
func newMySQLDB() {
	dbUser := os.Getenv("MYSQL_DB_USERNAME")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")
	dbDatabase := os.Getenv("MYSQL_DB_DATABASE")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")

	url := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?parseTime=true"
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to mysql")
	})
}

// Pool return a unique instance of DB
func DB() *gorm.DB {
	return db
}
