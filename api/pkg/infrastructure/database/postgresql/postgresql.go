package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/km1110/task-copilot-server/pkg/config"
	_ "github.com/lib/pq"
)

func NewPostgresConnector() (*sql.DB, error) {
	dbConf := config.NewConfig().DBConfig()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConf.USER,
		dbConf.PASSWORD,
		dbConf.HOST,
		dbConf.PORT,
		dbConf.DATABASE,
	)

	db, err := sql.Open(dbConf.DRIVER, dsn)

	if err != nil {
		tryConnCnt := 1
		for tryConnCnt <= 10 && err != nil {
			timer := time.NewTimer(2 * time.Second)
			db, err = sql.Open("postgres", dsn)
			<-timer.C

			log.Println("NewPostgresConnector: tryConnDB =", tryConnCnt)
			tryConnCnt += 1
		}
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connection has been established!")

	return db, nil
}
