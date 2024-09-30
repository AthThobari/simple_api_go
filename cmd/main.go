package main

import (
	"fmt"
	"log"

	"github.com/AthThobari/simple_api_go/internal/configs"
	"github.com/AthThobari/simple_api_go/internal/handlers/memberships"
	 membershipRepo "github.com/AthThobari/simple_api_go/internal/repository/memberships"
	 membershipSvc "github.com/AthThobari/simple_api_go/internal/service/memberships"
	"github.com/AthThobari/simple_api_go/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"},),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	fmt.Println("config", cfg)

	db,err:=internalsql.Connect(cfg.Database.DataSourceName)
	
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080
}

