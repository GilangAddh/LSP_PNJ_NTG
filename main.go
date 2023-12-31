package main

import (
	"fmt"
	"log"

	"LSP_PNJ_NTG/entity"

	"LSP_PNJ_NTG/account"
	"LSP_PNJ_NTG/handler"
	"LSP_PNJ_NTG/jenis_lsp"
	"LSP_PNJ_NTG/jenis_sk"
	"LSP_PNJ_NTG/lsp"
	"LSP_PNJ_NTG/lsp_sk"
	"LSP_PNJ_NTG/sk"

	"github.com/gin-gonic/gin"

	// "gorm.io/driver/mysql"

	"gorm.io/driver/postgres"
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
	db.AutoMigrate(&entity.LSP{})
	db.AutoMigrate(&entity.JenisSK{})
	db.AutoMigrate(&entity.SK{})
	db.AutoMigrate(&entity.LSP_SK{})
	db.AutoMigrate(&entity.Accounts{})
	db.AutoMigrate(&entity.Skema_Sertifikasi{})
	db.AutoMigrate(&entity.Persyaratan_Spesifikasi{})

	jenisLSPHandler := handler.NewJenisLSPHandler(jenis_lsp.NewService(jenis_lsp.NewRepository(db)))
	lspHandler := handler.NewLSPHandler(lsp.NewService(lsp.NewRepository(db)))
	accountHandler := handler.NewAccountHandler(account.NewService(account.NewRepository(db)))
	jenisSKHandler := handler.NewJenisSKHandler(jenis_sk.NewService(jenis_sk.NewRepository(db)))
	skHandler := handler.NewSKHandler(sk.NewService(sk.NewRepository(db)))
	lspskHandler := handler.NewLSPSKHandler(lsp_sk.NewService(lsp_sk.NewRepository(db)))

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
	v1.POST("/lsp", lspHandler.CreateLSP)
	v1.GET("/lsp", lspHandler.GetLSPs)
	v1.GET("/lsp/:id", lspHandler.GetLSP)
	v1.DELETE("/lsp/:id", lspHandler.DeleteLSP)
	v1.PUT("/lsp/:id", lspHandler.UpdateLSP)

	// Account Route
	v1.POST("/auth", accountHandler.Authentification)
	v1.POST("/register", accountHandler.Registration)

	// Jenis SK Route
	v1.POST("/jenis-sk", jenisSKHandler.CreateJenisSK)
	v1.GET("/jenis-sk", jenisSKHandler.GetJenisSKs)
	v1.GET("/jenis-sk/:id", jenisSKHandler.GetJenisSK)
	v1.DELETE("/jenis-sk/:id", jenisSKHandler.DeleteJenisSK)
	v1.PUT("/jenis-sk/:id", jenisSKHandler.UpdatejenisSK)

	// SK ROUTE
	v1.POST("/sk", skHandler.CreateSK)
	v1.GET("/sk", skHandler.GetSKs)
	v1.GET("/sk/:id", skHandler.GetSKs)
	v1.DELETE("/sk/:id", skHandler.DeleteSK)
	v1.PUT("/sk/:id", skHandler.UpdateSK)

	// LSP SK ROUTE
	v1.POST("/lsp-sk", lspskHandler.CreateLSPSK)
	v1.GET("/lsp-sk", lspskHandler.GetLSPSKs)
	v1.GET("/lsp-sk/:id", lspskHandler.GetLSPSK)
	v1.DELETE("/lsp-sk/:id", lspskHandler.DeleteLSPSK)
	v1.PUT("/lsp-sk/:id", lspskHandler.UpdateLSPSK)

	router.Run(":8081")
}
