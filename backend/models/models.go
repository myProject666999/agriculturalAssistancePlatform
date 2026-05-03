package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password     string         `json:"-" gorm:"size:255;not null"`
	Nickname     string         `json:"nickname" gorm:"size:50"`
	Email        string         `json:"email" gorm:"size:100"`
	Phone        string         `json:"phone" gorm:"size:20"`
	Avatar       string         `json:"avatar" gorm:"size:255"`
	Balance      float64        `json:"balance" gorm:"default:0"`
	Status       int            `json:"status" gorm:"default:1"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Merchant struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password     string         `json:"-" gorm:"size:255;not null"`
	StoreName    string         `json:"store_name" gorm:"size:100"`
	ContactName    string         `json:"contact_name" gorm:"size:50"`
	ContactPhone string         `json:"contact_phone" gorm:"size:20"`
	Address      string         `json:"address" gorm:"size:255"`
	Description  string         `json:"description" gorm:"type:text"`
	Status       int            `json:"status" gorm:"default:1"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Admin struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Password     string         `json:"-" gorm:"size:255;not null"`
	Nickname     string         `json:"nickname" gorm:"size:50"`
	Role         string         `json:"role" gorm:"size:20;default:'admin'"`
	Status       int            `json:"status" gorm:"default:1"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ProductCategory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:50;not null"`
	Description string         `json:"description" gorm:"size:255"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Product struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:100;not null"`
	CategoryID     uint           `json:"category_id"`
	MerchantID     uint           `json:"merchant_id"`
	Price          float64        `json:"price" gorm:"not null"`
	OriginalPrice  float64        `json:"original_price"`
	Stock          int            `json:"stock" gorm:"default:0"`
	Sales          int            `json:"sales" gorm:"default:0"`
	Image          string         `json:"image" gorm:"size:255"`
	Images         string         `json:"images" gorm:"type:text"`
	Description    string         `json:"description" gorm:"type:text"`
	Status         int            `json:"status" gorm:"default:1"`
	IsRecommend int            `json:"is_recommend" gorm:"default:0"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

type Cart struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	OrderNo       string         `json:"order_no" gorm:"uniqueIndex;size:50;not null"`
	UserID        uint           `json:"user_id"`
	TotalAmount   float64        `json:"total_amount" gorm:"not null"`
	Status        int            `json:"status" gorm:"default:0"`
	PaymentMethod  int            `json:"payment_method"`
	PaymentTime   *time.Time     `json:"payment_time"`
	AddressID     uint           `json:"address_id"`
	AddressInfo   string         `json:"address_info" gorm:"size:500"`
	Remark        string         `json:"remark" gorm:"size:255"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	ProductID uint           `json:"product_id"`
	OrderID   uint           `json:"order_id"`
	Content   string         `json:"content" gorm:"type:text"`
	Rating    int            `json:"rating" gorm:"default:5"`
	Images    string         `json:"images" gorm:"type:text"`
	Status    int            `json:"status" gorm:"default:1"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Address struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	Name        string         `json:"name" gorm:"size:50;not null"`
	Phone       string         `json:"phone" gorm:"size:20;not null"`
	Province    string         `json:"province" gorm:"size:50"`
	City        string         `json:"city" gorm:"size:50"`
	District    string         `json:"district" gorm:"size:50"`
	Detail      string         `json:"detail" gorm:"size:255;not null"`
	IsDefault   int            `json:"is_default" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type News struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Author      string         `json:"author" gorm:"size:50"`
	Image       string         `json:"image" gorm:"size:255"`
	Views       int            `json:"views" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Forum struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Images      string         `json:"images" gorm:"type:text"`
	Views       int            `json:"views" gorm:"default:0"`
	Likes       int            `json:"likes" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type ForumReply struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ForumID   uint      `json:"forum_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageBoard struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	Content     string         `json:"content" gorm:"type:text;not null"`
	Reply       string         `json:"reply" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Banner struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:100"`
	Image       string         `json:"image" gorm:"size:255;not null"`
	Link        string         `json:"link" gorm:"size:255"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Notice struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type RechargeRecord struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Amount      float64   `json:"amount" gorm:"not null"`
	OrderNo     string    `json:"order_no" gorm:"uniqueIndex;size:50;not null"`
	Status      int       `json:"status" gorm:"default:0"`
	PaymentTime *time.Time `json:"payment_time"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaymentRecord struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	OrderID     uint      `json:"order_id"`
	Amount      float64   `json:"amount" gorm:"not null"`
	PaymentMethod int     `json:"payment_method"`
	Status      int       `json:"status" gorm:"default:0"`
	PaymentTime *time.Time `json:"payment_time"`
	CreatedAt   time.Time `json:"created_at"`
}
