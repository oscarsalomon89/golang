package repository

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// const (
// 	DB_HOST = "127.0.0.1"
// 	DB_PORT = 3306
// 	DB_NAME = "test-db"
// 	DB_USER = "root"
// 	DB_PASS = "secret"
// )

// var db *gorm.DB //nolint:gochecknoglobals

// func GetConnectionDB() (*gorm.DB, error) {
// 	var err error

// 	if db == nil {
// 		//"root:secret@tcp(127.0.0.1:3306)/test-db?charset=utf8&parseTime=True"
// 		db, err = gorm.Open(mysql.Open(dbConnectionURL()), &gorm.Config{})
// 		if err != nil {
// 			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
// 			return nil, fmt.Errorf("### DB ERROR: %w", err)
// 		}
// 	}

// 	return db, nil
// }

// func dbConnectionURL() string {
// 	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
// }
