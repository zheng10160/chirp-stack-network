package storage

import (
	"github.com/jon177/lky-network-server/internal/config"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)


func Setup(config config.Config) error {
	d, err := sqlx.Open("mysql", config.Storage.Mysql.DSN)
	if err != nil {
		errors.Wrap(err, "storage: Mysql connection error")
	}

	db = &DBLogger{d}



	return nil
}
