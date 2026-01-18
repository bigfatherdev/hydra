package main

import (
	"log"
	"os"
	 "context"
	 "time"

	"github.com/gin-gonic/gin"

	"geoip-api/internal/api"
	"geoip-api/internal/auth"
	"geoip-api/internal/cache"
	"geoip-api/internal/geoip"
	"geoip-api/internal/middleware"
	"geoip-api/internal/storage"
	"geoip-api/internal/workers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// ============================
	// INFRA
	// ============================
	redis := cache.NewRedis()
	db := storage.NewMariaDB()

	mmdb := geoip.NewMMDB(
		"./geoip-api/cmd/GeoLite2-Country.mmdb",
	)

	manticore := geoip.NewManticore()

	handler := api.NewHandler(redis, db, mmdb, manticore)

	// ============================
	// ROUTER (CRIAR UMA VEZ)
	// ============================
	r := gin.New()
	r.Use(gin.Recovery())

	// 🔥 CORS GLOBAL
	r.Use(middleware.CORSHydra())

	// ============================
	// ROTAS
	// ============================
	v1 := r.Group("/v1")

	// ---- GEOIP (API KEY) ----
	geoipGroup := v1.Group("/geoip")
	geoipGroup.Use(auth.RequireAPIKey(db.DB()))
	geoipGroup.GET("", handler.GeoIP)
	
	// ---- MASSCAN ----
masscanGroup := v1.Group("/masscan")
masscanGroup.Use(auth.RequireAPIKey(db.DB()))
masscanGroup.POST("", handler.Masscan)
	
	
	maria := storage.NewMariaFromDB(db.DB())

r.GET("/status/ws", api.StatusWebSocket(maria))
r.GET("/status", api.StatusPage)


	// ---- KEYS (COOKIE / SSO) ----
	keys := v1.Group("/keys")
	keys.Use(auth.RequireUser())

	keys.GET("", handler.ListAPIKeys)
	keys.POST("", handler.CreateAPIKey)
	keys.POST("/:id/regenerate", handler.RegenerateAPIKey)
	keys.DELETE("/:id", handler.DeleteAPIKey)

	// ============================
	// IMPORTAÇÃO MANUAL (CLI)
	// ============================
	if len(os.Args) > 1 {
		folder := os.Args[1]
		log.Println("[main] importing IPs from:", folder)

		if err := workers.ImportIPsFromFolderBatch(
			db.DB(),
			folder,
			1000,
		); err != nil {
			log.Fatal("import failed:", err)
		}
	}
	
	worker := workers.NewIPEnrichmentWorker(
			db.DB(),
			"ef385f6322a319406f9f7e3a2b7576e29b2e1e29e0e161634f1356fce704c053", // API KEY interna
		)

	// ============================
	// WORKERS
	// ============================
	workDir := "/www/wwwroot/blackcerb.xyz/hydra.services/geoip-api/ip_lists"
	

	// Cria contexto cancelável
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // garante que o worker será parado ao sair do main

	// Executa worker em background
	go workers.StopForumSpamWorker(ctx, db.DB(), workDir)
go workers.URLHausWorker(ctx, db.DB())
go workers.URLHausOnWorker(ctx, db.DB())
go worker.RunMasscanWorker(ctx, 1*time.Minute) // roda a cada 1 min
go worker.RunSTFSWorker(ctx, 1*time.Minute) // roda a cada 1 min
go workers.FuzzerWSWorker(ctx, db.DB())


	

	// ============================
	// START
	// ============================
	log.Println("🚀 GeoIP API running on :8102")
	if err := r.Run(":8102"); err != nil {
		log.Fatal(err)
	}
}

