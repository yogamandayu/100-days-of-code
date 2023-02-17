package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)
	dsn := "root:c77b0dbf56@tcp(127.0.0.1:3306)/isolation_level_test?charset=utf8mb4&parseTime=True&loc=Local"
	conn1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	},
	)
	if err != nil {
		panic(err)
	}
	conn2, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	},
	)
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
	go func() {
		data := &Menu{
			ID: menu.ID,
		}
		err := conn1.First(&data).Error
		if err != nil {
			panic(err)
		}
		fmt.Println("Transaction 1 first query, menu price : ", data.Price)

		time.Sleep(500 * time.Millisecond)

		err = conn1.First(&data).Error
		if err != nil {
			panic(err)
		}
		fmt.Println("Transaction 1 second query, menu price : ", data.Price)
	}()

	go conn2.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Menu{}).Where("id = ?", menu.ID).Update("Price", 15000).Error
		if err != nil {
			panic(err)
		}
		fmt.Println("Transaction 2 updating row")
		time.Sleep(1 * time.Second)
		tx.Commit()
		fmt.Println("Transaction 2 rollback")
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
}

func nonRepeatableReads(db *gorm.DB) {

}

func phantomReads(db *gorm.DB) {

}
