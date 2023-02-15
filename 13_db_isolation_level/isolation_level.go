package main

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Menu struct {
	ID        string `gorm:"primaryKey;index;not null"`
	Name      string `gorm:"size:255"`
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func main() {
	conn1, conn2 := connect()
	conn1.AutoMigrate(&Menu{})

	dirtyReads(conn1, conn2)
	time.Sleep(3 * time.Second)
}

func connect() (*gorm.DB, *gorm.DB) {
	dsn := "host=127.0.0.1 user=devplay password=c77b0dbf56 dbname=isolation_level_test port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	conn1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	conn2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return conn1, conn2
}

// To do transaction:
// 1. Insert 1 row
// 2. Select row
// 3. Update row, but delay for commit
// 4. Select row again before committing update
// Expected result at second query:
// READ UNCOMMITTED : Have updated value.
// READ COMMITTED : Have init value.
// REPEATABLE READ : Have init value.
// SERIALIZABLE : Have init value.
func dirtyReads(conn1 *gorm.DB, conn2 *gorm.DB) {
	menu := &Menu{
		ID:    uuid.NewString(),
		Name:  "Chocolate",
		Price: 10000,
	}

	err := conn1.Create(&menu).Error
	if err != nil {
		panic(err)
	}

	// Transaction 1 & 3.
	go conn1.Transaction(func(tx *gorm.DB) error {
		data := &Menu{
			ID: menu.ID,
		}
		err := tx.First(&data).Error
		if err != nil {
			return err
		}
		fmt.Println("Transaction 1 first query, menu price : ", data.Price)

		time.Sleep(500 * time.Millisecond)

		err = tx.First(&data).Error
		if err != nil {
			return err
		}
		fmt.Println("Transaction 1 second query, menu price : ", data.Price)
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})

	go func() {

		tx := conn2.Begin()
		err := tx.Model(&Menu{}).Where("id = ?", menu.ID).Update("Price", 15000).Error
		if err != nil {
			panic(err)
		}
		fmt.Println("Transaction 2 updating row")
		time.Sleep(2 * time.Second)
		fmt.Println("Transaction 2 rollback")
		tx.Rollback()
	}()
}

func nonRepeatableReads(db *gorm.DB) {

}

func phantomReads(db *gorm.DB) {

}
