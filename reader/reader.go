package reader

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-readwrite-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-readwrite/bytes"
	wof_reader "github.com/whosonfirst/go-whosonfirst-readwrite/reader"
	wof_sqlite "github.com/whosonfirst/go-whosonfirst-sqlite"
	wof_database "github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"io"
)

type SQLiteReader struct {
	wof_reader.Reader
	database *wof_database.SQLiteDatabase
	table    wof_sqlite.Table
}

func NewSQLiteReader(dsn string, args ...interface{}) (wof_reader.Reader, error) {

	db, tbl, err := database.NewSQLiteDatabase(dsn)

	if err != nil {
		return nil, err
	}

	r := SQLiteReader{
		database: db,
		table:    tbl,
	}

	return &r, nil
}

func (r *SQLiteReader) Read(path string) (io.ReadCloser, error) {

	id, err := uri.IdFromPath(path)

	if err != nil {
		return nil, err
	}

	conn, err := r.database.Conn()

	if err != nil {
		return nil, err
	}

	q := fmt.Sprintf("SELECT body FROM %s WHERE id=?", r.table.Name())
	row := conn.QueryRow(q, id)

	var body string
	err = row.Scan(&body)

	if err != nil {
		return nil, err
	}

	return bytes.ReadCloserFromBytes([]byte(body))
}
