package database

import (
	"fmt"
	"log"
	"time"

	"agriculturalAssistancePlatform/config"
	"agriculturalAssistancePlatform/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error
	cfg := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = DB.AutoMigrate(
		&models.User{},
		&models.Merchant{},
		&models.Admin{},
		&models.ProductCategory{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
		&models.Favorite{},
		&models.Comment{},
		&models.Address{},
		&models.News{},
		&models.Forum{},
		&models.ForumReply{},
		&models.MessageBoard{},
		&models.Banner{},
		&models.Notice{},
		&models.RechargeRecord{},
		&models.PaymentRecord{},
	)

	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	InitAdmin()
	InitProductCategories()

	log.Println("数据库连接成功，迁移完成")
}

func InitAdmin() {
	var count int64
	DB.Model(&models.Admin{}).Count(&count)

	if count == 0 {
		admin := models.Admin{
			Username: "admin",
			Password: HashPassword("123456"),
			Nickname: "超级管理员",
			Role:     "super_admin",
			Status:   1,
		}
		DB.Create(&admin)
		log.Println("初始化管理员账号: admin / 123456")
	}
}

func InitProductCategories() {
	var count int64
	DB.Model(&models.ProductCategory{}).Count(&count)

	if count == 0 {
		categories := []models.ProductCategory{
			{Name: "新鲜蔬菜", Description: "当季新鲜蔬菜", SortOrder: 1},
			{Name: "时令水果", Description: "新鲜采摘的水果", SortOrder: 2},
			{Name: "特色农产品", Description: "地方特色农产品", SortOrder: 3},
			{Name: "粮油副食", Description: "精选粮油和副食", SortOrder: 4},
			{Name: "禽肉蛋类", Description: "新鲜禽肉和蛋类", SortOrder: 5},
			{Name: "水产海鲜", Description: "新鲜水产海鲜", SortOrder: 6},
		}

		DB.Create(&categories)
		log.Println("初始化商品分类成功")
	}
}

func HashPassword(password string) string {
	return "hashed_" + password
}
