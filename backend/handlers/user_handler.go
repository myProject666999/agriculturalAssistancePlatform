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

func GetUserProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
			"phone":    user.Phone,
			"email":    user.Email,
			"avatar":   user.Avatar,
			"balance":  user.Balance,
		},
	})
}

func UpdateUserProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	user.UpdatedAt = time.Now()
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetAddressList(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var addresses []models.Address
	database.DB.Where("user_id = ?", userID).Order("is_default DESC, id ASC").Find(&addresses)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": addresses,
	})
}

func GetAddressDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	addressID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var address models.Address
	if err := database.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "地址不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": address,
	})
}

func CreateAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Name     string `json:"name" binding:"required"`
		Phone    string `json:"phone" binding:"required"`
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
		Detail   string `json:"detail" binding:"required"`
		IsDefault int    `json:"is_default"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if req.IsDefault == 1 {
		database.DB.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_default", 0)
	}

	address := models.Address{
		UserID:    userID,
		Name:      req.Name,
		Phone:     req.Phone,
		Province:  req.Province,
		City:      req.City,
		District:  req.District,
		Detail:    req.Detail,
		IsDefault: req.IsDefault,
	}

	if err := database.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": address,
	})
}

func UpdateAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)
	addressID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
		Detail   string `json:"detail"`
		IsDefault int    `json:"is_default"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var address models.Address
	if err := database.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "地址不存在",
		})
		return
	}

	if req.IsDefault == 1 && address.IsDefault != 1 {
		database.DB.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_default", 0)
	}

	if req.Name != "" {
		address.Name = req.Name
	}
	if req.Phone != "" {
		address.Phone = req.Phone
	}
	if req.Province != "" {
		address.Province = req.Province
	}
	if req.City != "" {
		address.City = req.City
	}
	if req.District != "" {
		address.District = req.District
	}
	if req.Detail != "" {
		address.Detail = req.Detail
	}
	address.IsDefault = req.IsDefault
	address.UpdatedAt = time.Now()

	database.DB.Save(&address)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)
	addressID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Where("id = ? AND user_id = ?", addressID, userID).Delete(&models.Address{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "地址不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func SetDefaultAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)
	addressID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var address models.Address
	if err := database.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "地址不存在",
		})
		return
	}

	database.DB.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_default", 0)

	address.IsDefault = 1
	address.UpdatedAt = time.Now()
	database.DB.Save(&address)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置成功",
	})
}

func GetDefaultAddress(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var address models.Address
	if err := database.DB.Where("user_id = ? AND is_default = 1", userID).First(&address).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": address,
	})
}

func GetFavoriteList(c *gin.Context) {
	userID := middleware.GetUserID(c)
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
	database.DB.Model(&models.Favorite{}).Where("user_id = ?", userID).Count(&total)

	var favorites []models.Favorite
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&favorites)

	var productIDs []uint
	for _, f := range favorites {
		productIDs = append(productIDs, f.ProductID)
	}

	var products []models.Product
	if len(productIDs) > 0 {
		database.DB.Where("id IN ?", productIDs).Find(&products)
	}

	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	type FavoriteWithProduct struct {
		ID        uint            `json:"id"`
		ProductID uint            `json:"product_id"`
		CreatedAt time.Time       `json:"created_at"`
		Product   *models.Product `json:"product,omitempty"`
	}

	var result []FavoriteWithProduct
	for _, f := range favorites {
		product := productMap[f.ProductID]
		result = append(result, FavoriteWithProduct{
			ID:        f.ID,
			ProductID: f.ProductID,
			CreatedAt: f.CreatedAt,
			Product:   &product,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       result,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func AddFavorite(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, req.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商品不存在",
		})
		return
	}

	var existing models.Favorite
	database.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&existing)
	if existing.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已收藏",
		})
		return
	}

	favorite := models.Favorite{
		UserID:    userID,
		ProductID: req.ProductID,
	}

	if err := database.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "收藏失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "收藏成功",
	})
}

func RemoveFavorite(c *gin.Context) {
	userID := middleware.GetUserID(c)
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Favorite{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已取消收藏",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "取消收藏成功",
	})
}

func CheckFavorite(c *gin.Context) {
	userID := middleware.GetUserID(c)
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var favorite models.Favorite
	result := database.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&favorite)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"is_favorite": result.RowsAffected > 0,
		},
	})
}

func CreateMessage(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	message := models.MessageBoard{
		UserID:  userID,
		Content: req.Content,
		Status:  0,
	}

	if err := database.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "提交失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "提交成功",
	})
}

func GetMyMessages(c *gin.Context) {
	userID := middleware.GetUserID(c)
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
	database.DB.Model(&models.MessageBoard{}).Where("user_id = ?", userID).Count(&total)

	var messages []models.MessageBoard
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&messages)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       messages,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}
