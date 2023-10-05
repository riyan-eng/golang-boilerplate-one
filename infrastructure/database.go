package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/riyan-eng/golang-boilerplate-one/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SqlDB *sql.DB
var GormDB *gorm.DB

func ConnectSqlDB() {
	dsn := fmt.Sprintf(`
		host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v
	`, env.POSTGRE_HOST, env.POSTGRE_USERNAME, env.POSTGRE_PASSWORD, env.POSTGRE_DATABASE, env.POSTGRE_PORT, env.POSTGRE_TIMEZONE)
	var err error
	SqlDB, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("sql database: can't connect to database")
		os.Exit(1)
	}
	SqlDB.SetMaxIdleConns(env.POSTGRE_CONN_MAX_IDLE)
	SqlDB.SetMaxOpenConns(env.POSTGRE_CONN_MAX_OPEN)
	SqlDB.SetConnMaxLifetime(time.Minute * env.POSTGRE_CONN_MAX_LIFETIME)
	if err := SqlDB.Ping(); err != nil {
		fmt.Printf("sql database: can't ping to database - %v \n", err)
		os.Exit(1)
	}

	fmt.Println("sql database: connection opened to database")
}

func ConnectGormDB() {
	var err error
	GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: SqlDB,
	}))
	if err != nil {
		fmt.Println("gorm database: can't connect to database")
		os.Exit(1)
	}
	fmt.Println("gorm database: connection opened to database")
}
