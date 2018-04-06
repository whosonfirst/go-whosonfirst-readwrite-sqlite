package database

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite-features/tables"
	"github.com/whosonfirst/go-whosonfirst-sqlite/database"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

func NewSQLiteDatabase(dsn string) (*database.SQLiteDatabase, sqlite.Table, error) {

	db, err := database.NewDBWithDriver("sqlite3", dsn)

	if err != nil {
		return nil, nil, err
	}

	tbl, err := tables.NewGeoJSONTable()

	if err != nil {
		return nil, nil, err
	}

	ok, err := utils.HasTable(db, tbl.Name())

	if err != nil {
		return nil, nil, err
	}

	if !ok {
		return nil, nil, errors.New(fmt.Sprintf("Database is missing %s table", tbl.Name()))
	}

	return db, tbl, nil
}
