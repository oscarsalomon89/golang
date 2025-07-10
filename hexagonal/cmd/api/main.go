package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/osalomon/market-api/internal/adapters/http/itemhdl"
	"github.com/osalomon/market-api/internal/adapters/mysql/itemrepo"
	itemsrv "github.com/osalomon/market-api/internal/application/item"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	ginEngine := gin.Default()
	v1 := ginEngine.Group("/v1/market-api")

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	//mysql.Open(root:root@tcp(127.0.0.1:3306)/meli?parseTime=True)
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	itemRepo := itemrepo.NewItemRepository(db)
	itemUseCase := itemsrv.NewItemUseCase(itemRepo)
	itemHandler := itemhdl.NewHandler(itemUseCase)

	itemhdl.NewRouter(itemHandler).AddRoutesV1(v1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := newHTTPGinServer(ginEngine).ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error running server: %v", err)
		}
	}()

	log.Println("Server running on port 8080")
	wg.Wait()
}

func newHTTPGinServer(router *gin.Engine) *http.Server {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
}
