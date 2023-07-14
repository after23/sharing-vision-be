package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/after23/sharing-vision-be/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	err = util.LoadEnv("../..")
	if err != nil {
		log.Panicf("failed to read config from env: %v", err)
	}
	testDB, err = sql.Open(util.GetConfig().DBDriver, util.GetConfig().DBSource)
	defer testDB.Close();
	if err != nil {
		log.Panicf("cannot connect to database: %v", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}