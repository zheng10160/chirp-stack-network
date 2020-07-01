package storage

import (
	"golang.org/x/net/context"
	"github.com/jmoiron/sqlx"
)

type Device struct {
	ID 				int64 				`db:"id"`
	Reserved 		string 				`db:"reserved"`
	KeyMain			string				`db:"key_main"`
	KeyExt			string				`db:"key_ext"`
	BatchId 		int64 				`db:"batch_id"`
	BatchSerial     string 				`db:"batch_serial"`
	Status          int 				`db:"status"`
	CreateTs		int64			`db:"create_ts"`
	ActivateTs		int64			`db:"activate_ts"`
}


type DeviceListItem struct {
	Device
}

func GetDevice(ctx context.Context, db sqlx.Queryer, ID int64) (Device, error) {
	var d Device
	err := sqlx.Get(db, &d, "select * from device where id = ?", ID)

	if err != nil {
		return d, handleMysqlError(Select, err, "select error")
	}

	return d, nil
}