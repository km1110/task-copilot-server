package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/km1110/task-copilot-server/pkg/config"
	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.POSTGRES_USER,
		config.Env.POSTGRES_PASSWORD,
		config.Env.POSTGRES_HOST,
		"5432",
		config.Env.POSTGRES_DB,
	)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		tryConnCnt := 1
		for tryConnCnt <= 10 && err != nil {
			timer := time.NewTimer(2 * time.Second)
			db, err = sql.Open("postgres", dsn)
			<-timer.C

			log.Println("NewDB: tryConnDB =", tryConnCnt)
			tryConnCnt += 1
		}
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connection has been established!")

	return db, nil
}
