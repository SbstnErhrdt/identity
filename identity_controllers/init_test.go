package identity_controllers

import (
	"fmt"
	"github.com/SbstnErhrdt/env"
	"github.com/SbstnErhrdt/identity/identity_install"
	"github.com/SbstnErhrdt/identity/identity_models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log/slog"
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
			slog.With("count", i).Warn("failed to connect to database, retrying in 5 seconds")
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
		slog.Error("database connection failed")
		panic("database connection failed")
	}
}

func connectToDb() (err error) {
	// gorm connect to postgres using env variables
	// create connection string with variables
	host := os.Getenv("SQL_HOST")
	user := os.Getenv("SQL_USER")
	pw := os.Getenv("SQL_PASSWORD")
	port := os.Getenv("SQL_PORT")
	dbName := os.Getenv("SQL_DATABASE")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=UTC", host, port, user, dbName, pw)
	fmt.Println("Connection string:", dsn)
	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return err
	}
	DbConnection = db
	return
}

func install() {
	err := identity_install.Install(DbConnection)
	if err != nil {
		slog.With("err", err).Error("failed to install database")
	}
}

func EmptyTable(s schema.Tabler) (err error) {
	return DbConnection.Exec(fmt.Sprintf("DELETE FROM %s", s.TableName())).Error
}

func EmptyIdentityTable() {
	err := EmptyTable(&identity_models.Identity{})
	if err != nil {
		slog.With("err", err).Error("failed to empty identity table")
	}
}

func EmptyRegistrationConfirmationTable() {
	err := EmptyTable(&identity_models.IdentityRegistrationConfirmation{})
	if err != nil {
		slog.With("err", err).Error("failed to empty identity table")
	}
}
