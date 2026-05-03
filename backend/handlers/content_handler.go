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

func GetNewsList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.News{}).Where("status = 1")

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

func GetNewsDetail(c *gin.Context) {
	newsID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var news models.News
	if err := database.DB.Where("id = ? AND status = 1", newsID).First(&news).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "新闻不存在",
		})
		return
	}

	database.DB.Model(&news).Update("views", news.Views+1)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": news,
	})
}

func GetForumList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Forum{}).Where("status = 1")

	var total int64
	query.Count(&total)

	var forums []models.Forum
	query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&forums)

	var userIDs []uint
	for _, f := range forums {
		userIDs = append(userIDs, f.UserID)
	}

	var users []models.User
	if len(userIDs) > 0 {
		database.DB.Where("id IN ?", userIDs).Find(&users)
	}

	userMap := make(map[uint]models.User)
	for _, u := range users {
		userMap[u.ID] = u
	}

	type ForumWithUser struct {
		models.Forum
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}

	var result []ForumWithUser
	for _, f := range forums {
		user := userMap[f.UserID]
		result = append(result, ForumWithUser{
			Forum:    f,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
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

func GetForumDetail(c *gin.Context) {
	forumID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var forum models.Forum
	if err := database.DB.Where("id = ? AND status = 1", forumID).First(&forum).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "帖子不存在",
		})
		return
	}

	database.DB.Model(&forum).Update("views", forum.Views+1)

	var user models.User
	database.DB.First(&user, forum.UserID)

	var replies []models.ForumReply
	database.DB.Where("forum_id = ?", forumID).Order("created_at ASC").Find(&replies)

	var replyUserIDs []uint
	for _, r := range replies {
		replyUserIDs = append(replyUserIDs, r.UserID)
	}

	var replyUsers []models.User
	if len(replyUserIDs) > 0 {
		database.DB.Where("id IN ?", replyUserIDs).Find(&replyUsers)
	}

	replyUserMap := make(map[uint]models.User)
	for _, u := range replyUsers {
		replyUserMap[u.ID] = u
	}

	type ReplyWithUser struct {
		models.ForumReply
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}

	var replyList []ReplyWithUser
	for _, r := range replies {
		replyUser := replyUserMap[r.UserID]
		replyList = append(replyList, ReplyWithUser{
			ForumReply: r,
			Username:   replyUser.Username,
			Nickname:   replyUser.Nickname,
			Avatar:     replyUser.Avatar,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"forum":    forum,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"nickname": user.Nickname,
				"avatar":   user.Avatar,
			},
			"replies": replyList,
		},
	})
}

func CreateForum(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Images  string `json:"images"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	forum := models.Forum{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		Images:  req.Images,
		Status:  1,
	}

	if err := database.DB.Create(&forum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "发布失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发布成功",
		"data": forum,
	})
}

func ReplyForum(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		ForumID uint   `json:"forum_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var forum models.Forum
	if err := database.DB.Where("id = ? AND status = 1", req.ForumID).First(&forum).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "帖子不存在",
		})
		return
	}

	reply := models.ForumReply{
		ForumID: req.ForumID,
		UserID:  userID,
		Content: req.Content,
	}

	if err := database.DB.Create(&reply).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "回复失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "回复成功",
	})
}

func GetMyForums(c *gin.Context) {
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

	query := database.DB.Model(&models.Forum{}).Where("user_id = ?", userID)

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

func GetNotices(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := database.DB.Model(&models.Notice{}).Where("status = 1")

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

func GetNoticeDetail(c *gin.Context) {
	noticeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var notice models.Notice
	if err := database.DB.Where("id = ? AND status = 1", noticeID).First(&notice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "公告不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": notice,
	})
}

func GetMerchantStatistics(c *gin.Context) {
	merchantID := middleware.GetUserID(c)

	var products []models.Product
	database.DB.Model(&models.Product{}).Where("merchant_id = ?", merchantID).Find(&products)

	var productIDs []uint
	for _, p := range products {
		productIDs = append(productIDs, p.ID)
	}

	var orderItems []models.OrderItem
	if len(productIDs) > 0 {
		database.DB.Where("product_id IN ?", productIDs).Find(&orderItems)
	}

	var orderIDs []uint
	for _, item := range orderItems {
		orderIDs = append(orderIDs, item.OrderID)
	}

	var paidOrders []models.Order
	if len(orderIDs) > 0 {
		database.DB.Where("id IN ? AND status >= 1", orderIDs).Find(&paidOrders)
	}

	var todayOrders []models.Order
	today := time.Now().Format("2006-01-02")
	if len(orderIDs) > 0 {
		database.DB.Where("id IN ? AND status >= 1 AND DATE(created_at) = ?", orderIDs, today).Find(&todayOrders)
	}

	var todaySales float64 = 0
	for _, order := range todayOrders {
		todaySales += order.TotalAmount
	}

	var monthOrders []models.Order
	monthStart := time.Now().Format("2006-01") + "-01"
	if len(orderIDs) > 0 {
		database.DB.Where("id IN ? AND status >= 1 AND DATE(created_at) >= ?", orderIDs, monthStart).Find(&monthOrders)
	}

	var monthSales float64 = 0
	for _, order := range monthOrders {
		monthSales += order.TotalAmount
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"total_products":    len(products),
			"total_orders":      len(paidOrders),
			"today_orders":      len(todayOrders),
			"today_sales":       todaySales,
			"month_sales":       monthSales,
		},
	})
}
