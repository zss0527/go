package db

import (
	"fmt"
	"gin_gorm_demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局数据库实例
var dbInstance *gorm.DB

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	if dbInstance == nil {
		var err error
		// 使用SQLite数据库，连接到test.db文件
		dsn := "root:Root123.@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
		dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// dbInstance, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			fmt.Printf("连接数据库失败: %v\n", err)
			panic("连接数据库失败")
		}
		// 执行自动迁移
		autoMigrate(dbInstance)
	}
	return dbInstance
}

// autoMigrate 自动迁移数据库表结构
func autoMigrate(db *gorm.DB) {
	// 导入所有需要自动迁移的模型
	// 这里以User模型为例，实际应用中可能有多个模型
	// 每个模型代表数据库中的一张表
	// 例如，如果有Product模型，也需要在这里添加
	// db.AutoMigrate(&model.Product{})
	// 这里假设User模型在your_project/model/user.go中定义
	// 导入模型所在的包
	// import (
	//     "gingormdemo/model"
	// )
	db.AutoMigrate(&model.User{})
}
