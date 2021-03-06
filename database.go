package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	jtltojson "github.com/nicholasvuono/jtl-to-json"
)

type Database struct {
	DB *bolt.DB
}

func (d *Database) open() {
	var err error
	d.DB, err = bolt.Open("bolt.db", 0600, nil)
	checkErr(err)
}

func (d *Database) createBuckets() {
	err := d.DB.Update(
		func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("plu"))
			checkErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte("blu"))
			checkErr(err)
			return err
		})
	checkErr(err)
}

func (d *Database) write(b string, r *jtltojson.Result) {
	err := d.DB.Update(
		func(tx *bolt.Tx) error {
			err := tx.Bucket([]byte(b)).Put([]byte(r.TestName+"_"+r.DateTime), r.JSON())
			checkErr(err)
			return err
		})
	checkErr(err)
}

func (d *Database) readAll(b string) []byte {
	l := make([]jtltojson.Result, 0)
	err := d.DB.View(
		func(tx *bolt.Tx) error {
			c := tx.Bucket([]byte(b)).Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				r := jtltojson.Result{}
				err := json.Unmarshal(v, &r)
				checkErr(err)
				l = append(l, r)
			}
			return nil
		})
	checkErr(err)
	json, err := json.Marshal(l)
	checkErr(err)
	return json
}

func (d *Database) backup(w http.ResponseWriter) error {
	err := d.DB.View(
		func(tx *bolt.Tx) error {
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", `attachment; filename="my.db"`)
			w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
			_, err := tx.WriteTo(w)
			return err
		})
	return err
}

func newDB() *Database {
	db := Database{}
	db.open()
	time.Sleep(5 * time.Second)
	db.createBuckets()
	return &db
}

var db = newDB()
