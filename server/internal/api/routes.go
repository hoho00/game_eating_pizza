package api

import (
	"net/http"

	"game_eating_pizza/internal/api/handlers"
	"game_eating_pizza/internal/api/middleware"
	"game_eating_pizza/internal/config"
	"game_eating_pizza/internal/repository"
	"game_eating_pizza/internal/services"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter는 Gin 라우터를 설정하고 반환합니다
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	// 환경에 따라 Gin 모드 설정
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// 미들웨어 설정
	router.Use(middleware.CORS(cfg))
	router.Use(middleware.ErrorHandler())
	router.Use(gin.Recovery())

	// Repository 초기화
	repos := repository.NewRepositories(db, cfg)

	// Service 초기화
	authService := services.NewAuthService(repos.Player)
	playerService := services.NewPlayerService(repos.Player, repos.Weapon)
	weaponService := services.NewWeaponService(repos.Weapon, repos.Player)
	dungeonService := services.NewDungeonService(repos.Dungeon)

	// Handler 초기화
	authHandler := handlers.NewAuthHandler(authService)
	playerHandler := handlers.NewPlayerHandler(playerService)
	weaponHandler := handlers.NewWeaponHandler(weaponService)
	dungeonHandler := handlers.NewDungeonHandler(dungeonService)

	// Swagger 문서
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 그룹
	v1 := router.Group("/api/v1")
	{
		// 인증 관련
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// 인증이 필요한 라우트
		authenticated := v1.Group("")
		authenticated.Use(middleware.AuthMiddleware(cfg))
		{
			// 플레이어 관련
			players := authenticated.Group("/players")
			{
				players.GET("", playerHandler.GetPlayers)
				players.GET("/me", playerHandler.GetMe)
				players.GET("/leaderboard", playerHandler.GetLeaderboard)
				players.GET("/:id", playerHandler.GetPlayer)
				players.PUT("/me", playerHandler.UpdateMe)
				players.DELETE("/:id", playerHandler.DeletePlayer)
			}

			// 무기 관련
			weapons := authenticated.Group("/weapons")
			{
				weapons.GET("", weaponHandler.GetWeapons)
				weapons.POST("", weaponHandler.CreateWeapon)
				weapons.GET("/:id", weaponHandler.GetWeapon)
				weapons.PUT("/:id", weaponHandler.UpdateWeapon)
				weapons.DELETE("/:id", weaponHandler.DeleteWeapon)
				weapons.PUT("/:id/upgrade", weaponHandler.UpgradeWeapon)
				weapons.PUT("/:id/equip", weaponHandler.EquipWeapon)
			}

			// 던전 관련
			dungeons := authenticated.Group("/dungeons")
			{
				dungeons.POST("", dungeonHandler.CreateDungeon)
				dungeons.GET("/all", dungeonHandler.GetAllDungeons)
				dungeons.GET("/active", dungeonHandler.GetActiveDungeons)
				dungeons.GET("/:id", dungeonHandler.GetDungeon)
				dungeons.PUT("/:id", dungeonHandler.UpdateDungeon)
				dungeons.DELETE("/:id", dungeonHandler.DeleteDungeon)
				dungeons.POST("/:id/enter", dungeonHandler.EnterDungeon)
				dungeons.POST("/:id/clear", dungeonHandler.ClearDungeon)
			}
		}
	}

	return router
}
