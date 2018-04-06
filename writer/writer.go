package writer

import (
	"errors"
	wof_writer "github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	wof_sqlite "github.com/whosonfirst/go-whosonfirst-sqlite"
	wof_database "github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io"
)

type SQLiteWriter struct {
	wof_writer.Writer
	database *wof_database.SQLiteDatabase
	table    wof_sqlite.Table
}

func NewSQLiteWriter(dsn string, args ...interface{}) (wof_writer.Writer, error) {

	db, tbl, err := database.NewSQLiteDatabase(dsn)

	if err != nil {
		return nil, err
	}

	wr := SQLiteWriter{
		database: db,
		table:    tbl,
	}

	return &wr, nil
}

func (wr *SQLiteWriter) Write(path string, fh io.ReadCloser) error {

	id, err := uri.IdFromPath(path)

	if err != nil {
		return err
	}

	return errors.New("Please write me")
}
