package writer

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	wof_writer "github.com/whosonfirst/go-whosonfirst-readwrite/writer"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite-features/tables"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"io"
)

type SQLiteWriter struct {
	wof_writer.Writer
	database *database.SQLiteDatabase
	table    sqlite.Table
}

func NewSQLiteWriter(dsn string, args ...interface{}) (wof_writer.Writer, error) {

	db, err := database.NewDBWithDriver("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	tbl, err := tables.NewGeoJSONTableWithDatabase(db)

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

     	f, err := feature.LoadFeatureFromReader(fh)

	if err != nil {
		return err
	}

	return wr.table.IndexRecord(wr.database, f)
}

func (wr *SQLiteWriter) URI(path string) string {
     return fmt.Sprintf("sqlite://%s/%s#%s", wr.database.DSN(), wr.table.Name(), path)
}
