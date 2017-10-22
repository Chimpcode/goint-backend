package db

import (
	"../utils"
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"time"
	"reflect"
)

var MasterDB map[string]*storm.DB

const TAG = "goint"

func InitDB(config *utils.GointConfig) error {
	dbT := reflect.ValueOf(config.Db)

	for i:=0;i<dbT.NumField();i++ {
		field := dbT.Type().Field(i)
		tag := field.Tag.Get(TAG)
		if tag == "-" || tag == "" {
			panic("Error at parse Goint config. File: init_db.go, line 20")
		}
		pathOfField := dbT.Field(i).String()
		tdb, err := storm.Open(pathOfField, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
		if err != nil {
			return err
		}
		MasterDB[tag] = tdb
	}

	return nil
}

func GetDBStatus() map[string]string {
	return make(map[string]string)
}