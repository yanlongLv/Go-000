package db

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// Connect mysql
type Connect struct {
	db  *sql.DB
	ctx context.Context
}

//NewConnect connect mysql
func NewConnect() (connect *Connect, err error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		return connect, errors.Wrap(err, "connect mysql error")
	}
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	return &Connect{db: db, ctx: ctx}, nil
}
