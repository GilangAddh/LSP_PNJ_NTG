package main

import (
	"fmt"
	"log"

	"LSP_PNJ_NTG/entity"
	"LSP_PNJ_NTG/handler"
	"LSP_PNJ_NTG/jenis_lsp"
	"LSP_PNJ_NTG/lsp"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// MYSQL
	// dsn := "root:@tcp(127.0.0.1:3306)/buku?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Postgre
	dsn := "host=localhost user=postgres password=superadmin dbname=LSP_PNJ_NTG port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.JenisLSP{})
	// db.AutoMigrate(&entity.LSP{})

	jenisLSPHandler := handler.NewJenisLSPHandler(jenis_lsp.NewService(jenis_lsp.NewRepository(db)))
	lspHanlder := handler.NewLSPHandler(lsp.NewService(lsp.NewRepository(db)))

	if err != nil {
		log.Fatal("Db connection Error")
	}

	fmt.Println("koneksi berhasil", db)

	router := gin.Default()

	v1 := router.Group("/v1")

	// Jenis LSP Route
	v1.POST("/jenis-lsp", jenisLSPHandler.CreateJenisLSP)
	v1.GET("/jenis-lsp", jenisLSPHandler.GetJenisLSPs)
	v1.GET("/jenis-lsp/:id", jenisLSPHandler.GetJenisLSP)
	v1.DELETE("/jenis-lsp/:id", jenisLSPHandler.DeleteJenisLSP)
	v1.PUT("/jenis-lsp/:id", jenisLSPHandler.UpdateJenisLSP)

	// LSP ROUTE
	v1.POST("/lsp", lspHanlder.CreateLSP)
	v1.GET("/lsp", lspHanlder.GetLSPs)
	v1.GET("/lsp/:id", lspHanlder.GetLSP)
	v1.DELETE("/lsp/:id", lspHanlder.DeleteLSP)
	v1.PUT("/lsp/:id", lspHanlder.UpdateLSP)

	router.Run(":8081")
}
