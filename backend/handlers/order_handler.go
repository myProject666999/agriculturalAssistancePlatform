package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/middleware"
	"agriculturalAssistancePlatform/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCartList(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var carts []models.Cart
	database.DB.Where("user_id = ?", userID).Find(&carts)

	var productIDs []uint
	for _, cart := range carts {
		productIDs = append(productIDs, cart.ProductID)
	}

	var products []models.Product
	if len(productIDs) > 0 {
		database.DB.Where("id IN ?", productIDs).Find(&products)
	}

	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	type CartWithProduct struct {
		ID        uint            `json:"id"`
		ProductID uint            `json:"product_id"`
		Quantity  int             `json:"quantity"`
		CreatedAt time.Time       `json:"created_at"`
		Product   *models.Product `json:"product,omitempty"`
	}

	var result []CartWithProduct
	var totalAmount float64 = 0

	for _, cart := range carts {
		product := productMap[cart.ProductID]
		if product.ID != 0 && product.Status == 1 {
			result = append(result, CartWithProduct{
				ID:        cart.ID,
				ProductID: cart.ProductID,
				Quantity:  cart.Quantity,
				CreatedAt: cart.CreatedAt,
				Product:   &product,
			})
			totalAmount += product.Price * float64(cart.Quantity)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"list":         result,
			"total_amount": totalAmount,
		},
	})
}

func AddToCart(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity"`
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

	if product.Status != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "商品已下架",
		})
		return
	}

	if req.Quantity <= 0 {
		req.Quantity = 1
	}

	var existingCart models.Cart
	database.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&existingCart)

	if existingCart.ID != 0 {
		existingCart.Quantity += req.Quantity
		if existingCart.Quantity > product.Stock && product.Stock > 0 {
			existingCart.Quantity = product.Stock
		}
		existingCart.UpdatedAt = time.Now()
		database.DB.Save(&existingCart)
	} else {
		if req.Quantity > product.Stock && product.Stock > 0 {
			req.Quantity = product.Stock
		}
		cart := models.Cart{
			UserID:    userID,
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
		}
		database.DB.Create(&cart)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func UpdateCart(c *gin.Context) {
	userID := middleware.GetUserID(c)
	cartID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("id = ? AND user_id = ?", cartID, userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "购物车项不存在",
		})
		return
	}

	if req.Quantity <= 0 {
		database.DB.Delete(&cart)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已移除",
		})
		return
	}

	var product models.Product
	database.DB.First(&product, cart.ProductID)

	if req.Quantity > product.Stock && product.Stock > 0 {
		req.Quantity = product.Stock
	}

	cart.Quantity = req.Quantity
	cart.UpdatedAt = time.Now()
	database.DB.Save(&cart)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func RemoveFromCart(c *gin.Context) {
	userID := middleware.GetUserID(c)
	cartID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	result := database.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&models.Cart{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "购物车项不存在",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "移除成功",
	})
}

func BatchRemoveFromCart(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if len(req.Ids) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请选择要删除的商品",
		})
		return
	}

	database.DB.Where("user_id = ? AND id IN ?", userID, req.Ids).Delete(&models.Cart{})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "移除成功",
	})
}

func GenerateOrderNo() string {
	now := time.Now()
	return fmt.Sprintf("ORD%s%06d", now.Format("20060102150405"), now.Nanosecond()%1000000)
}

func CreateOrder(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req struct {
		CartIds     []uint `json:"cart_ids" binding:"required"`
		AddressID   uint   `json:"address_id" binding:"required"`
		Remark      string `json:"remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if len(req.CartIds) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "请选择要购买的商品",
		})
		return
	}

	var address models.Address
	if err := database.DB.Where("id = ? AND user_id = ?", req.AddressID, userID).First(&address).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "收货地址不存在",
		})
		return
	}

	var carts []models.Cart
	database.DB.Where("user_id = ? AND id IN ?", userID, req.CartIds).Find(&carts)

	if len(carts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "购物车为空",
		})
		return
	}

	var productIDs []uint
	for _, cart := range carts {
		productIDs = append(productIDs, cart.ProductID)
	}

	var products []models.Product
	database.DB.Where("id IN ?", productIDs).Find(&products)

	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	var totalAmount float64 = 0
	var orderItems []models.OrderItem
	var stockErrors []string

	for _, cart := range carts {
		product := productMap[cart.ProductID]
		if product.ID == 0 {
			stockErrors = append(stockErrors, fmt.Sprintf("商品ID:%d 不存在", cart.ProductID))
			continue
		}
		if product.Status != 1 {
			stockErrors = append(stockErrors, fmt.Sprintf("商品:%s 已下架", product.Name))
			continue
		}
		if cart.Quantity > product.Stock && product.Stock > 0 {
			stockErrors = append(stockErrors, fmt.Sprintf("商品:%s 库存不足", product.Name))
			continue
		}

		totalAmount += product.Price * float64(cart.Quantity)

		orderItems = append(orderItems, models.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     product.Price,
		})
	}

	if len(stockErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  strings.Join(stockErrors, "；"),
		})
		return
	}

	if len(orderItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "没有可购买的商品",
		})
		return
	}

	addressInfo := fmt.Sprintf("%s %s %s %s %s", address.Name, address.Phone, address.Province, address.City, address.Detail)

	tx := database.DB.Begin()

	order := models.Order{
		OrderNo:     GenerateOrderNo(),
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      0,
		AddressID:   req.AddressID,
		AddressInfo: addressInfo,
		Remark:      req.Remark,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建订单失败: " + err.Error(),
		})
		return
	}

	for i := range orderItems {
		orderItems[i].OrderID = order.ID
	}

	if err := tx.Create(&orderItems).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建订单详情失败: " + err.Error(),
		})
		return
	}

	for _, cart := range carts {
		product := productMap[cart.ProductID]
		if err := tx.Model(&models.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock-cart.Quantity).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "更新库存失败: " + err.Error(),
			})
			return
		}
	}

	if err := tx.Where("user_id = ? AND id IN ?", userID, req.CartIds).Delete(&models.Cart{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "清空购物车失败: " + err.Error(),
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建订单成功",
		"data": gin.H{
			"order_id": order.ID,
			"order_no": order.OrderNo,
			"total_amount": order.TotalAmount,
		},
	})
}

func GetOrderList(c *gin.Context) {
	userID := middleware.GetUserID(c)
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

	query := database.DB.Model(&models.Order{}).Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	query.Count(&total)

	var orders []models.Order
	query.Order("id DESC").Offset(offset).Limit(pageSize).Find(&orders)

	var orderIDs []uint
	for _, order := range orders {
		orderIDs = append(orderIDs, order.ID)
	}

	var orderItems []models.OrderItem
	if len(orderIDs) > 0 {
		database.DB.Where("order_id IN ?", orderIDs).Find(&orderItems)
	}

	orderItemMap := make(map[uint][]models.OrderItem)
	for _, item := range orderItems {
		orderItemMap[item.OrderID] = append(orderItemMap[item.OrderID], item)
	}

	var productIDs []uint
	for _, item := range orderItems {
		productIDs = append(productIDs, item.ProductID)
	}

	var products []models.Product
	if len(productIDs) > 0 {
		database.DB.Where("id IN ?", productIDs).Find(&products)
	}

	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	type OrderWithItems struct {
		models.Order
		Items []struct {
			models.OrderItem
			Product *models.Product `json:"product"`
		} `json:"items"`
	}

	var result []OrderWithItems
	for _, order := range orders {
		orderWithItems := OrderWithItems{
			Order: order,
		}

		for _, item := range orderItemMap[order.ID] {
			product := productMap[item.ProductID]
			orderWithItems.Items = append(orderWithItems.Items, struct {
				models.OrderItem
				Product *models.Product `json:"product"`
			}{
				OrderItem: item,
				Product:   &product,
			})
		}

		result = append(result, orderWithItems)
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

func GetOrderDetail(c *gin.Context) {
	userID := middleware.GetUserID(c)
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var order models.Order
	if err := database.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	var orderItems []models.OrderItem
	database.DB.Where("order_id = ?", orderID).Find(&orderItems)

	var productIDs []uint
	for _, item := range orderItems {
		productIDs = append(productIDs, item.ProductID)
	}

	var products []models.Product
	if len(productIDs) > 0 {
		database.DB.Where("id IN ?", productIDs).Find(&products)
	}

	productMap := make(map[uint]models.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	type OrderItemWithProduct struct {
		models.OrderItem
		Product *models.Product `json:"product"`
	}

	var items []OrderItemWithProduct
	for _, item := range orderItems {
		product := productMap[item.ProductID]
		items = append(items, OrderItemWithProduct{
			OrderItem: item,
			Product:   &product,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": gin.H{
			"order": order,
			"items": items,
		},
	})
}

func CancelOrder(c *gin.Context) {
	userID := middleware.GetUserID(c)
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var order models.Order
	if err := database.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	if order.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "只能取消待付款的订单",
		})
		return
	}

	tx := database.DB.Begin()

	var orderItems []models.OrderItem
	tx.Where("order_id = ?", orderID).Find(&orderItems)

	for _, item := range orderItems {
		tx.Model(&models.Product{}).Where("id = ?", item.ProductID).UpdateColumn("stock", gorm.Expr("stock + ?", item.Quantity))
	}

	order.Status = -1
	order.UpdatedAt = time.Now()
	tx.Save(&order)

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "取消成功",
	})
}

func ConfirmOrder(c *gin.Context) {
	userID := middleware.GetUserID(c)
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}

	var order models.Order
	if err := database.DB.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	if order.Status != 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "只能确认已发货的订单",
		})
		return
	}

	order.Status = 3
	order.UpdatedAt = time.Now()
	database.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "确认收货成功",
	})
}

func GetMerchantOrders(c *gin.Context) {
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

	var products []models.Product
	database.DB.Model(&models.Product{}).Where("merchant_id = ?", merchantID).Find(&products)

	var productIDs []uint
	for _, p := range products {
		productIDs = append(productIDs, p.ID)
	}

	if len(productIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": gin.H{
				"list":       []interface{}{},
				"total":      0,
				"page":       page,
				"page_size":  pageSize,
			},
		})
		return
	}

	var orderItems []models.OrderItem
	database.DB.Where("product_id IN ?", productIDs).Find(&orderItems)

	var orderIDs []uint
	orderProductMap := make(map[uint][]uint)
	for _, item := range orderItems {
		orderIDs = append(orderIDs, item.OrderID)
		orderProductMap[item.OrderID] = append(orderProductMap[item.OrderID], item.ProductID)
	}

	if len(orderIDs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": gin.H{
				"list":       []interface{}{},
				"total":      0,
				"page":       page,
				"page_size":  pageSize,
			},
		})
		return
	}

	query := database.DB.Model(&models.Order{}).Where("id IN ?", orderIDs)

	if status != "" {
		query = query.Where("status = ?", status)
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

func UpdateMerchantOrderStatus(c *gin.Context) {
	merchantID := middleware.GetUserID(c)
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
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

	var products []models.Product
	database.DB.Model(&models.Product{}).Where("merchant_id = ?", merchantID).Find(&products)

	var productIDs []uint
	for _, p := range products {
		productIDs = append(productIDs, p.ID)
	}

	var orderItems []models.OrderItem
	database.DB.Where("order_id = ? AND product_id IN ?", orderID, productIDs).Find(&orderItems)

	if len(orderItems) == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "无权操作此订单",
		})
		return
	}

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "订单不存在",
		})
		return
	}

	if req.Status == 2 && order.Status == 1 {
		order.Status = 2
		order.UpdatedAt = time.Now()
		database.DB.Save(&order)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "发货成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "状态更新失败",
		})
	}
}
