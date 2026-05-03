package handlers

import (
	"net/http"
	"strconv"
	"time"

	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/middleware"
	"agriculturalAssistancePlatform/models"

	"github.com/gin-gonic/gin"
)

func GetProductCategories(c *gin.Context) {
	var categories []models.ProductCategory
	database.DB.Where("status = 1").Order("sort_order ASC").Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": categories,
	})
}

func GetProductList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryID := c.Query("category_id")
	keyword := c.Query("keyword")
	isRecommend := c.Query("is_recommend")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Product{}).Where("status = 1")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	if isRecommend == "1" {
		query = query.Where("is_recommend = 1")
	}

	var total int64
	query.Count(&total)

	var products []models.Product
	query.Order("sales DESC, id DESC").Offset(offset).Limit(pageSize).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       products,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func GetProductDetail(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	if product.Status != 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品已下架",
		})
		return
	}

	var comments []models.Comment
	database.DB.Where("product_id = ? AND status = 1", productID).Order("created_at DESC").Limit(5).Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"product":  product,
			"comments": comments,
		},
	})
}

func GetProductComments(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var total int64
	database.DB.Model(&models.Comment{}).Where("product_id = ? AND status = 1", productID).Count(&total)

	var comments []models.Comment
	database.DB.Where("product_id = ? AND status = 1", productID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       comments,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateComment(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		ProductID uint   `json:"product_id" binding:"required"`
		OrderID   uint   `json:"order_id" binding:"required"`
		Content   string `json:"content" binding:"required"`
		Rating    int    `json:"rating"`
		Images    string `json:"images"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, req.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	if order.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "无权操作此订单",
		})
		return
	}

	if order.Status < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "订单未完成，无法评论",
		})
		return
	}

	var existingComment models.Comment
	database.DB.Where("user_id = ? AND product_id = ? AND order_id = ?", userID, req.ProductID, req.OrderID).First(&existingComment)
	if existingComment.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "已评论过此商品",
		})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		req.Rating = 5
	}

	comment := models.Comment{
		UserID:    userID,
		ProductID: req.ProductID,
		OrderID:   req.OrderID,
		Content:   req.Content,
		Rating:    req.Rating,
		Images:    req.Images,
		Status:    1,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "评论失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "评论成功",
	})
}

func GetBanners(c *gin.Context) {
	var banners []models.Banner
	database.DB.Where("status = 1").Order("sort_order ASC, id DESC").Find(&banners)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": banners,
	})
}

func GetRecommendProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if limit < 1 || limit > 50 {
		limit = 10
	}

	var products []models.Product
	database.DB.Where("status = 1 AND is_recommend = 1").Order("sales DESC, id DESC").Limit(limit).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": products,
	})
}

func GetHotProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if limit < 1 || limit > 50 {
		limit = 10
	}

	var products []models.Product
	database.DB.Where("status = 1").Order("sales DESC, id DESC").Limit(limit).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": products,
	})
}

func GetNewProducts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if limit < 1 || limit > 50 {
		limit = 10
	}

	var products []models.Product
	database.DB.Where("status = 1").Order("created_at DESC, id DESC").Limit(limit).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": products,
	})
}

func GetMerchantProducts(c *gin.Context) {
	merchantID := middleware.GetUserID(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Product{}).Where("merchant_id = ?", merchantID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var products []models.Product
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       products,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateMerchantProduct(c *gin.Context) {
	merchantID := middleware.GetUserID(c)

	var req struct {
		Name          string  `json:"name" binding:"required"`
		CategoryID    uint    `json:"category_id" binding:"required"`
		Price         float64 `json:"price" binding:"required"`
		OriginalPrice float64 `json:"original_price"`
		Stock         int     `json:"stock"`
		Image         string  `json:"image"`
		Images        string  `json:"images"`
		Description   string  `json:"description"`
		IsRecommend   int     `json:"is_recommend"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	product := models.Product{
		Name:          req.Name,
		CategoryID:    req.CategoryID,
		MerchantID:    merchantID,
		Price:         req.Price,
		OriginalPrice: req.OriginalPrice,
		Stock:         req.Stock,
		Image:         req.Image,
		Images:        req.Images,
		Description:   req.Description,
		Status:        1,
		IsRecommend:   req.IsRecommend,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": product,
	})
}

func UpdateMerchantProduct(c *gin.Context) {
	merchantID := middleware.GetUserID(c)
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var product models.Product
	if err := database.DB.Where("id = ? AND merchant_id = ?", productID, merchantID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	var req struct {
		Name          string  `json:"name"`
		CategoryID    uint    `json:"category_id"`
		Price         float64 `json:"price"`
		OriginalPrice float64 `json:"original_price"`
		Stock         int     `json:"stock"`
		Image         string  `json:"image"`
		Images        string  `json:"images"`
		Description   string  `json:"description"`
		Status        int     `json:"status"`
		IsRecommend   int     `json:"is_recommend"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.CategoryID != 0 {
		product.CategoryID = req.CategoryID
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.OriginalPrice != 0 {
		product.OriginalPrice = req.OriginalPrice
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}
	if req.Image != "" {
		product.Image = req.Image
	}
	if req.Images != "" {
		product.Images = req.Images
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Status != 0 || req.Status == 0 {
		product.Status = req.Status
	}
	product.IsRecommend = req.IsRecommend
	product.UpdatedAt = time.Now()

	database.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteMerchantProduct(c *gin.Context) {
	merchantID := middleware.GetUserID(c)
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Where("id = ? AND merchant_id = ?", productID, merchantID).Delete(&models.Product{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
