package db

import (
	"log"

	"github.com/AnanyaDevGo/microServices-grpc-project_order_svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func Init(url string) Handler {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Order{})

    return Handler{db}
}
