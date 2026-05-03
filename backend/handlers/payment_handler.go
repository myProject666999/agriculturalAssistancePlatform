package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/middleware"
	"agriculturalAssistancePlatform/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateRechargeNo() string {
	now := time.Now()
	return fmt.Sprintf("RC%s%06d", now.Format("20060102150405"), now.Nanosecond()%1000000)
}

func CreateRecharge(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Amount float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "充值金额必须大于0",
		})
		return
	}

	recharge := models.RechargeRecord{
		UserID:  userID,
		Amount:  req.Amount,
		OrderNo: GenerateRechargeNo(),
		Status:  0,
	}

	if err := database.DB.Create(&recharge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建充值订单失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": gin.H{
			"recharge_id": recharge.ID,
			"order_no":    recharge.OrderNo,
			"amount":      recharge.Amount,
		},
	})
}

func RechargePay(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		RechargeID uint `json:"recharge_id" binding:"required"`
		PayMethod  int  `json:"pay_method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var recharge models.RechargeRecord
	if err := database.DB.Where("id = ? AND user_id = ?", req.RechargeID, userID).First(&recharge).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "充值订单不存在",
		})
		return
	}

	if recharge.Status == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "充值已成功",
		})
		return
	}

	tx := database.DB.Begin()

	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}

	user.Balance += recharge.Amount
	user.UpdatedAt = time.Now()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新余额失败",
		})
		return
	}

	now := time.Now()
	recharge.Status = 1
	recharge.PaymentTime = &now
	if err := tx.Save(&recharge).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新充值状态失败",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "充值成功",
		"data": gin.H{
			"new_balance": user.Balance,
		},
	})
}

func GetRechargeList(c *gin.Context) {
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
	database.DB.Model(&models.RechargeRecord{}).Where("user_id = ?", userID).Count(&total)

	var records []models.RechargeRecord
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&records)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":       records,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

func OrderPay(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		OrderID       uint `json:"order_id" binding:"required"`
		PaymentMethod int  `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var order models.Order
	if err := database.DB.Where("id = ? AND user_id = ?", req.OrderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	if order.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "订单状态不正确",
		})
		return
	}

	tx := database.DB.Begin()

	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}

	if user.Balance < order.TotalAmount {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "余额不足",
		})
		return
	}

	user.Balance -= order.TotalAmount
	user.UpdatedAt = time.Now()
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新余额失败",
		})
		return
	}

	now := time.Now()
	order.Status = 1
	order.PaymentTime = &now
	order.PaymentMethod = req.PaymentMethod
	order.UpdatedAt = now
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新订单状态失败",
		})
		return
	}

	payment := models.PaymentRecord{
		UserID:        userID,
		OrderID:       order.ID,
		Amount:        order.TotalAmount,
		PaymentMethod: req.PaymentMethod,
		Status:        1,
		PaymentTime:   &now,
	}
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建支付记录失败",
		})
		return
	}

	var orderItems []models.OrderItem
	tx.Where("order_id = ?", order.ID).Find(&orderItems)

	for _, item := range orderItems {
		tx.Model(&models.Product{}).Where("id = ?", item.ProductID).UpdateColumn("sales", gorm.Expr("sales + ?", item.Quantity))
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "支付成功",
		"data": gin.H{
			"new_balance": user.Balance,
		},
	})
}

func GetBalance(c *gin.Context) {
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
			"balance": user.Balance,
		},
	})
}
