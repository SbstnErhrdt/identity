package identity_controllers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/mail"
	"os"
	"time"
)

const testUserEmail = "test@erhardt.net"

// testService is the test service
var testService *ControllerService

// DbConnection is the database connection
var DbConnection *gorm.DB

func init() {
	testEmailSender := os.Getenv("TEST_IDENTITY_EMAIL_SENDER")

	testService = NewService("Identity-Test", mail.Address{
		Name:    "Test-Sender",
		Address: testEmailSender,
	})

	// connect to database
	ConnectToDbAndRetry()
	testService.SetSQLClient(DbConnection)
}

const maxRetry = 5

func ConnectToDbAndRetry() {
	i := 1
	for {
		errConn := connectToDb()
		if errConn == nil {
			break
		} else {
			log.Printf("failed to connect to database for the %d time, retrying in 5 seconds", i)
			i++
			time.Sleep(time.Second * 5)
		}
		if i >= maxRetry {
			panic(errConn)
		}
	}
	var res int
	DbConnection.Raw("SELECT 1 + 1").Scan(&res)
	if res != 2 {
		log.Fatal("database connection failed")
	}
}

func connectToDb() (err error) {
	// gorm connect to postgres using env variables
	// create connection string with variables
	var SqlHost = os.Getenv("SQL_HOST")
	var SqlUser = os.Getenv("SQL_USER")
	var SqlPassword = os.Getenv("SQL_PASSWORD")
	var SqlPort = os.Getenv("SQL_PORT")
	var SqlDatabase = os.Getenv("SQL_DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin",
		SqlHost, SqlUser, SqlPassword, SqlDatabase, SqlPort)
	fmt.Println("Connection string:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DbConnection = db
	return
}
