package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/km1110/task-copilot-server/pkg/adapter/repository"
)

type mysqlHandler struct {
	db *sql.DB
}

func NewMySQLHandler(c *config) (*mysqlHandler, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?%s",
		c.user,
		c.password,
		c.host,
		c.database,
		c.option,
	)

	db, err := sql.Open(c.driver, dsn)
	if err != nil {
		tryConnCnt := 1
		for tryConnCnt <= 3 && err != nil {
			timer := time.NewTimer(2 * time.Second)
			db, err = sql.Open(c.driver, dsn)
			<-timer.C

			log.Println("NewMySQLHandler: tryConnDB = ", tryConnCnt)
			tryConnCnt += 1
		}

		if err != nil {
			return &mysqlHandler{}, err
		}
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &mysqlHandler{db: db}, nil
}

func (m mysqlHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return mysqlTx{}, err
	}

	return newMySQLTx(tx), nil
}

func (m mysqlHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m mysqlHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newMySQLRows(rows)

	return row, nil
}

func (m mysqlHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := m.db.QueryRowContext(ctx, query, args...)

	return newMySQLRow(row)
}

type mysqlRow struct {
	row *sql.Row
}

func newMySQLRow(row *sql.Row) mysqlRow {
	return mysqlRow{row: row}
}

func (pr mysqlRow) Scan(dest ...interface{}) error {
	if err := pr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}

type mysqlRows struct {
	rows *sql.Rows
}

func newMySQLRows(rows *sql.Rows) mysqlRows {
	return mysqlRows{rows: rows}
}

func (pr mysqlRows) Scan(dest ...interface{}) error {
	if err := pr.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (pr mysqlRows) Next() bool {
	return pr.rows.Next()
}

func (pr mysqlRows) Err() error {
	return pr.rows.Err()
}

func (pr mysqlRows) Close() error {
	return pr.rows.Close()
}

type mysqlTx struct {
	tx *sql.Tx
}

func newMySQLTx(tx *sql.Tx) mysqlTx {
	return mysqlTx{tx: tx}
}

func (m mysqlTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m mysqlTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := m.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newMySQLRows(rows)

	return row, nil
}

func (m mysqlTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := m.tx.QueryRowContext(ctx, query, args...)

	return newMySQLRow(row)
}

func (m mysqlTx) Commit() error {
	return m.tx.Commit()
}

func (m mysqlTx) Rollback() error {
	return m.tx.Rollback()
}
