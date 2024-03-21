package pg

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/AjxGnx/software-engineer-exercise/config"
	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq"
)

var (
	instance *sql.DB
	once     sync.Once
)

func ConnInstance() *sql.DB {
	once.Do(func() {
		instance = getConnection()
	})

	return instance
}

func getConnection() *sql.DB {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable",
		config.Environments().DBUser,
		config.Environments().DBPass,
		config.Environments().DBHost,
		config.Environments().DBPort,
		config.Environments().DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
