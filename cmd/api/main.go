package main

import (
	"bunny-fairy-api/internal/config"
	"bunny-fairy-api/internal/models/home"
	"bunny-fairy-api/internal/models/rbac"
	"bunny-fairy-api/internal/router"
	"bunny-fairy-api/pkg/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	if err := db.AutoMigrate(
		&rbac.User{},
		&rbac.Role{},
		&rbac.Permission{},
		&rbac.UserRole{},
		&rbac.RolePermission{},
		&home.Category{},
		&home.Goods{},
		&home.Banner{},
	); err != nil {
		log.Fatal("Failed to migrate:", err)
	}
	//
	//// 测试用户
	//db.FirstOrCreate(&model.User{}, model.User{
	//	Username: "alice",
	//	Password: "123456",
	//	Email:    "alice@example.com",
	//})

	// 依赖注入
	//userRepo := repository.NewUserRepository(db)
	//userService := service.NewUserService(userRepo)
	//userHandler := handler.NewUserHandler(userService)
	//orderHandler := handler.NewOrderHandler()

	// 启动服务（一行注册所有 handler）
	// 设置路由
	r := router.SetupRouter(db, "my-secret-key-123")

	//handler.RegisterAll(r, userHandler, orderHandler) // 后续加新 handler 直接在这里加

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

/*
package main

import (
    "log"

    "your-project/models"
    "your-project/router"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // 数据库连接
    dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("数据库连接失败:", err)
    }

    // 自动迁移表结构
    err = db.AutoMigrate(
        &models.User{},
        &models.Role{},
        &models.Permission{},
        &models.UserRole{},
        &models.RolePermission{},
    )
    if err != nil {
        log.Fatal("表迁移失败:", err)
    }

    // 设置路由
    r := router.SetupRouter(db, "your-secret-key")

    // 启动服务
    err = r.Run(":8080")
    if err != nil {
        log.Fatal("服务启动失败:", err)
    }
}
*/
