package identity_controllers

import (
	"fmt"
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_install"
	"github.com/SbstnErhrdt/identity/identity_models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/mail"
	"os"
	"time"
)

var testUserEmail string = ""
var testSenderEmail string = ""
var testAdminEmail string = ""

// testService is the test service
var testService *ControllerService

// DbConnection is the database connection
var DbConnection *gorm.DB

func init() {

	env.LoadEnvFiles("../test/.env")
	env.CheckRequiredEnvironmentVariables(
		"TEST_SENDER_EMAIL",
		"TEST_USER_EMAIL",
		"TEST_ADMIN_EMAIL",
	)

	testSenderEmail = os.Getenv("TEST_SENDER_EMAIL")
	testUserEmail = os.Getenv("TEST_USER_EMAIL")
	testAdminEmail = os.Getenv("TEST_ADMIN_EMAIL")

	testService = NewService("Identity-Test", mail.Address{
		Name:    "Test-Sender",
		Address: testSenderEmail,
	})
	testService.SetAdminEmail(testAdminEmail)

	// connect to database
	ConnectToDbAndRetry()
	testService.SetSQLClient(DbConnection)
	install()
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

func install() {
	err := identity_install.Install(DbConnection)
	if err != nil {
		log.WithError(err).Fatal("failed to install database")
	}
}

func EmptyTable(s schema.Tabler) (err error) {
	return DbConnection.Exec(fmt.Sprintf("DELETE FROM %s", s.TableName())).Error
}

func EmptyIdentityTable() {
	err := EmptyTable(&identity_models.Identity{})
	if err != nil {
		log.WithError(err).Error("failed to empty identity table")
	}
}

func EmptyRegistrationConfirmationTable() {
	err := EmptyTable(&identity_models.IdentityRegistrationConfirmation{})
	if err != nil {
		log.WithError(err).Error("failed to empty identity table")
	}
}
