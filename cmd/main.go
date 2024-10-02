package main

import (
	"log"

	"github.com/AthThobari/simple_api_go/internal/configs"
	"github.com/AthThobari/simple_api_go/internal/handlers/memberships"
	"github.com/AthThobari/simple_api_go/internal/handlers/posts"
	membershipRepo "github.com/AthThobari/simple_api_go/internal/repository/memberships"
	postRepo "github.com/AthThobari/simple_api_go/internal/repository/posts"
	membershipSvc "github.com/AthThobari/simple_api_go/internal/service/memberships"
	postSvc "github.com/AthThobari/simple_api_go/internal/service/posts"
	"github.com/AthThobari/simple_api_go/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()
	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080
}
