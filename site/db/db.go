package db

// Imported Packages
import (
	"billing/site/configs"
	"database/sql"
	"fmt"

	// Postgres library
	_ "github.com/lib/pq"
)

// Making global variables in order to share
// With other functions
var (
	client *sql.DB
	err    error
)

// Init our db
func Init() (err error) {
	database := configs.Get().Database
	appConfigs := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Dbname)

	// Connecting to our Postgres server
	client, err = sql.Open("postgres", appConfigs)

	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}

	return
}

// Connect to Database
func Connect() *sql.DB {
	return client
}

// Exit to Database
func Exit() {
	client.Close()
}
