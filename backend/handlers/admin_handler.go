package handlers

import (
	"net/http"
	"strconv"
	"time"

	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/models"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.User{})

	if keyword != "" {
		query = query.Where("username LIKE ? OR nickname LIKE ? OR phone LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var users []models.User
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       users,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func UpdateUserStatus(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Status int `json:"status" binding:"required"`
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

	user.Status = req.Status
	user.UpdatedAt = time.Now()
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetMerchantList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	keyword := c.Query("keyword")
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Merchant{})

	if keyword != "" {
		query = query.Where("username LIKE ? OR store_name LIKE ? OR contact_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var merchants []models.Merchant
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&merchants)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       merchants,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateMerchant(c *gin.Context) {
	var req struct {
		Username     string `json:"username" binding:"required"`
		Password     string `json:"password" binding:"required"`
		StoreName    string `json:"store_name"`
		ContactName  string `json:"contact_name"`
		ContactPhone string `json:"contact_phone"`
		Address      string `json:"address"`
		Description  string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var existingMerchant models.Merchant
	database.DB.Where("username = ?", req.Username).First(&existingMerchant)
	if existingMerchant.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名已存在",
		})
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "密码加密失败",
		})
		return
	}

	merchant := models.Merchant{
		Username:     req.Username,
		Password:     hashedPassword,
		StoreName:    req.StoreName,
		ContactName:  req.ContactName,
		ContactPhone: req.ContactPhone,
		Address:      req.Address,
		Description:  req.Description,
		Status:       1,
	}

	if err := database.DB.Create(&merchant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": merchant,
	})
}

func UpdateMerchantStatus(c *gin.Context) {
	merchantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Status int `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var merchant models.Merchant
	if err := database.DB.First(&merchant, merchantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "商家不存在",
		})
		return
	}

	merchant.Status = req.Status
	merchant.UpdatedAt = time.Now()
	database.DB.Save(&merchant)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetAdminList(c *gin.Context) {
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
	database.DB.Model(&models.Admin{}).Count(&total)

	var admins []models.Admin
	database.DB.Order("id DESC").Offset(offset).Limit(pageSize).Find(&admins)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       admins,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateAdmin(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var existingAdmin models.Admin
	database.DB.Where("username = ?", req.Username).First(&existingAdmin)
	if existingAdmin.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名已存在",
		})
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "密码加密失败",
		})
		return
	}

	if req.Role == "" {
		req.Role = "admin"
	}

	admin := models.Admin{
		Username: req.Username,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Role:     req.Role,
		Status:   1,
	}

	if err := database.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": admin,
	})
}

func GetAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")
	orderNo := c.Query("order_no")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Order{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if orderNo != "" {
		query = query.Where("order_no LIKE ?", "%"+orderNo+"%")
	}

	var total int64
	query.Count(&total)

	var orders []models.Order
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       orders,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")
	keyword := c.Query("keyword")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Product{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
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

func UpdateProductStatus(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Status int `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
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

	product.Status = req.Status
	product.UpdatedAt = time.Now()
	database.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetBannerList(c *gin.Context) {
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
	database.DB.Model(&models.Banner{}).Count(&total)

	var banners []models.Banner
	database.DB.Order("sort_order ASC, id DESC").Offset(offset).Limit(pageSize).Find(&banners)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       banners,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateBanner(c *gin.Context) {
	var req struct {
		Title     string `json:"title"`
		Image     string `json:"image" binding:"required"`
		Link      string `json:"link"`
		SortOrder int    `json:"sort_order"`
		Status    int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	banner := models.Banner{
		Title:     req.Title,
		Image:     req.Image,
		Link:      req.Link,
		SortOrder: req.SortOrder,
		Status:    req.Status,
	}

	if err := database.DB.Create(&banner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": banner,
	})
}

func UpdateBanner(c *gin.Context) {
	bannerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Title     string `json:"title"`
		Image     string `json:"image"`
		Link      string `json:"link"`
		SortOrder int    `json:"sort_order"`
		Status    int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var banner models.Banner
	if err := database.DB.First(&banner, bannerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "轮播图不存在",
		})
		return
	}

	if req.Title != "" {
		banner.Title = req.Title
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Link != "" {
		banner.Link = req.Link
	}
	banner.SortOrder = req.SortOrder
	banner.Status = req.Status
	banner.UpdatedAt = time.Now()

	database.DB.Save(&banner)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteBanner(c *gin.Context) {
	bannerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Delete(&models.Banner{}, bannerID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "轮播图不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GetNewsListAdmin(c *gin.Context) {
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

	query := database.DB.Model(&models.News{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var news []models.News
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&news)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       news,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateNews(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Author  string `json:"author"`
		Image   string `json:"image"`
		Status  int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	news := models.News{
		Title:   req.Title,
		Content: req.Content,
		Author:  req.Author,
		Image:   req.Image,
		Status:  req.Status,
	}

	if err := database.DB.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": news,
	})
}

func UpdateNews(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Author  string `json:"author"`
		Image   string `json:"image"`
		Status  int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var news models.News
	if err := database.DB.First(&news, newsID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "新闻不存在",
		})
		return
	}

	if req.Title != "" {
		news.Title = req.Title
	}
	if req.Content != "" {
		news.Content = req.Content
	}
	if req.Author != "" {
		news.Author = req.Author
	}
	if req.Image != "" {
		news.Image = req.Image
	}
	news.Status = req.Status
	news.UpdatedAt = time.Now()

	database.DB.Save(&news)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteNews(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Delete(&models.News{}, newsID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "新闻不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GetForumListAdmin(c *gin.Context) {
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

	query := database.DB.Model(&models.Forum{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var forums []models.Forum
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&forums)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       forums,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func UpdateForumStatus(c *gin.Context) {
	forumID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Status int `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var forum models.Forum
	if err := database.DB.First(&forum, forumID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "帖子不存在",
		})
		return
	}

	forum.Status = req.Status
	forum.UpdatedAt = time.Now()
	database.DB.Save(&forum)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetMessageBoardList(c *gin.Context) {
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

	query := database.DB.Model(&models.MessageBoard{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var messages []models.MessageBoard
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&messages)

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

func ReplyMessage(c *gin.Context) {
	messageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Reply string `json:"reply" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var message models.MessageBoard
	if err := database.DB.First(&message, messageID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "留言不存在",
		})
		return
	}

	message.Reply = req.Reply
	message.Status = 1
	message.UpdatedAt = time.Now()
	database.DB.Save(&message)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "回复成功",
	})
}

func GetNoticeListAdmin(c *gin.Context) {
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

	query := database.DB.Model(&models.Notice{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var notices []models.Notice
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&notices)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       notices,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func CreateNotice(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content"`
		Status  int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	notice := models.Notice{
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
	}

	if err := database.DB.Create(&notice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": notice,
	})
}

func UpdateNotice(c *gin.Context) {
	noticeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Status  int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var notice models.Notice
	if err := database.DB.First(&notice, noticeID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "公告不存在",
		})
		return
	}

	if req.Title != "" {
		notice.Title = req.Title
	}
	if req.Content != "" {
		notice.Content = req.Content
	}
	notice.Status = req.Status
	notice.UpdatedAt = time.Now()

	database.DB.Save(&notice)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteNotice(c *gin.Context) {
	noticeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Delete(&models.Notice{}, noticeID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "公告不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GetProductCategoryList(c *gin.Context) {
	var categories []models.ProductCategory
	database.DB.Order("sort_order ASC, id DESC").Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": categories,
	})
}

func CreateProductCategory(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
		Status      int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	category := models.ProductCategory{
		Name:        req.Name,
		Description: req.Description,
		SortOrder:   req.SortOrder,
		Status:      req.Status,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": category,
	})
}

func UpdateProductCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
		Status      int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var category models.ProductCategory
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "分类不存在",
		})
		return
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	category.SortOrder = req.SortOrder
	category.Status = req.Status
	category.UpdatedAt = time.Now()

	database.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteProductCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var count int64
	database.DB.Model(&models.Product{}).Where("category_id = ?", categoryID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "该分类下有商品，无法删除",
		})
		return
	}

	result := database.DB.Delete(&models.ProductCategory{}, categoryID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "分类不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func GetAdminStatistics(c *gin.Context) {
	var userCount int64
	database.DB.Model(&models.User{}).Count(&userCount)

	var merchantCount int64
	database.DB.Model(&models.Merchant{}).Count(&merchantCount)

	var productCount int64
	database.DB.Model(&models.Product{}).Count(&productCount)

	var orderCount int64
	database.DB.Model(&models.Order{}).Where("status >= 1").Count(&orderCount)

	var todayOrders []models.Order
	today := time.Now().Format("2006-01-02")
	database.DB.Model(&models.Order{}).Where("status >= 1 AND DATE(created_at) = ?", today).Find(&todayOrders)

	var todaySales float64 = 0
	for _, order := range todayOrders {
		todaySales += order.TotalAmount
	}

	var monthOrders []models.Order
	monthStart := time.Now().Format("2006-01") + "-01"
	database.DB.Model(&models.Order{}).Where("status >= 1 AND DATE(created_at) >= ?", monthStart).Find(&monthOrders)

	var monthSales float64 = 0
	for _, order := range monthOrders {
		monthSales += order.TotalAmount
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"user_count":     userCount,
			"merchant_count": merchantCount,
			"product_count":  productCount,
			"order_count":    orderCount,
			"today_orders":   len(todayOrders),
			"today_sales":    todaySales,
			"month_sales":    monthSales,
		},
	})
}
