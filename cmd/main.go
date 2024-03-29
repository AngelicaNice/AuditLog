package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AngelicaNice/AuditLog/internal/config"
	"github.com/AngelicaNice/AuditLog/internal/repository"
	"github.com/AngelicaNice/AuditLog/internal/server"
	"github.com/AngelicaNice/AuditLog/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repository.NewAudit(db)
	auditServise := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditServise)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
