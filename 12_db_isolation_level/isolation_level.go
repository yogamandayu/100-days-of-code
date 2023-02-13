package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = connect()
}

func connect() *gorm.DB {
	dsn := "host=127.0.0.1 user= password= dbname=isolation_level_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func dirtyReads(db *gorm.DB) {

}

func nonRepeatableReads(db *gorm.DB) {

}

func phantomReads(db *gorm.DB) {

}
