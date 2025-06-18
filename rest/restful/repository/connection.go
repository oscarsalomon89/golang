package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = 3306
	DB_NAME = "meli"
	DB_USER = "root"
	DB_PASS = "root"
)

var db *sqlx.DB //nolint:gochecknoglobals

func GetConnectionDB() (*sqlx.DB, error) {
	var err error

	if db == nil {
		//"root:root@tcp(127.0.0.1:3306)/meli?charset=utf8&parseTime=True"
		db, err = sqlx.Open("mysql", dbConnectionURL())
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
			return nil, fmt.Errorf("### DB ERROR: %w", err)
		}
	}

	return db, nil
}

func dbConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}

// var db *gorm.DB

// func GetConnectionDB() (*gorm.DB, error) {
//     var err error

//     if db == nil {
//         dsn := dbConnectionURL() // Usa tu función existente para armar el DSN
//         db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//         if err != nil {
//             fmt.Printf("########## DB ERROR: %s #############\n", err.Error())
//             return nil, fmt.Errorf("### DB ERROR: %w", err)
//         }
//     }

//     return db, nil
// }
