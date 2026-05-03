package routers

import (
	"agriculturalAssistancePlatform/handlers"
	"agriculturalAssistancePlatform/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/user/register", handlers.UserRegister)
			auth.POST("/user/login", handlers.UserLogin)
			auth.POST("/merchant/login", handlers.MerchantLogin)
			auth.POST("/admin/login", handlers.AdminLogin)
		}

		api.GET("/banners", handlers.GetBanners)
		api.GET("/notices", handlers.GetNotices)
		api.GET("/notices/:id", handlers.GetNoticeDetail)

		api.GET("/categories", handlers.GetProductCategories)
		api.GET("/products", handlers.GetProductList)
		api.GET("/products/:id", handlers.GetProductDetail)
		api.GET("/products/:id/comments", handlers.GetProductComments)
		api.GET("/products/recommend", handlers.GetRecommendProducts)
		api.GET("/products/hot", handlers.GetHotProducts)
		api.GET("/products/new", handlers.GetNewProducts)

		api.GET("/news", handlers.GetNewsList)
		api.GET("/news/:id", handlers.GetNewsDetail)

		api.GET("/forums", handlers.GetForumList)
		api.GET("/forums/:id", handlers.GetForumDetail)

		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware("user"))
		{
			user.GET("/info", handlers.GetUserInfo)
			user.PUT("/password", handlers.UpdateUserPassword)

			user.GET("/profile", handlers.GetUserProfile)
			user.PUT("/profile", handlers.UpdateUserProfile)

			user.GET("/addresses", handlers.GetAddressList)
			user.GET("/addresses/default", handlers.GetDefaultAddress)
			user.GET("/addresses/:id", handlers.GetAddressDetail)
			user.POST("/addresses", handlers.CreateAddress)
			user.PUT("/addresses/:id", handlers.UpdateAddress)
			user.DELETE("/addresses/:id", handlers.DeleteAddress)
			user.POST("/addresses/:id/default", handlers.SetDefaultAddress)

			user.GET("/cart", handlers.GetCartList)
			user.POST("/cart", handlers.AddToCart)
			user.PUT("/cart/:id", handlers.UpdateCart)
			user.DELETE("/cart/:id", handlers.RemoveFromCart)
			user.POST("/cart/batch-delete", handlers.BatchRemoveFromCart)

			user.POST("/orders", handlers.CreateOrder)
			user.GET("/orders", handlers.GetOrderList)
			user.GET("/orders/:id", handlers.GetOrderDetail)
			user.POST("/orders/:id/pay", handlers.OrderPay)
			user.POST("/orders/:id/cancel", handlers.CancelOrder)
			user.POST("/orders/:id/confirm", handlers.ConfirmOrder)

			user.GET("/favorites", handlers.GetFavoriteList)
			user.POST("/favorites", handlers.AddFavorite)
			user.DELETE("/favorites/:product_id", handlers.RemoveFavorite)
			user.GET("/favorites/check/:product_id", handlers.CheckFavorite)

			user.POST("/comments", handlers.CreateComment)

			user.POST("/forums", handlers.CreateForum)
			user.POST("/forums/reply", handlers.ReplyForum)
			user.GET("/forums/my", handlers.GetMyForums)

			user.POST("/messages", handlers.CreateMessage)
			user.GET("/messages/my", handlers.GetMyMessages)

			user.GET("/balance", handlers.GetBalance)
			user.POST("/recharge", handlers.CreateRecharge)
			user.POST("/recharge/pay", handlers.RechargePay)
			user.GET("/recharge/records", handlers.GetRechargeList)
		}

		merchant := api.Group("/merchant")
		merchant.Use(middleware.AuthMiddleware("merchant"))
		{
			merchant.GET("/info", handlers.GetUserInfo)
			merchant.PUT("/password", handlers.UpdateUserPassword)

			merchant.GET("/statistics", handlers.GetMerchantStatistics)

			merchant.GET("/products", handlers.GetMerchantProducts)
			merchant.POST("/products", handlers.CreateMerchantProduct)
			merchant.PUT("/products/:id", handlers.UpdateMerchantProduct)
			merchant.DELETE("/products/:id", handlers.DeleteMerchantProduct)

			merchant.GET("/orders", handlers.GetMerchantOrders)
			merchant.PUT("/orders/:id/status", handlers.UpdateMerchantOrderStatus)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware("admin"))
		{
			admin.GET("/info", handlers.GetUserInfo)
			admin.PUT("/password", handlers.UpdateUserPassword)

			admin.GET("/statistics", handlers.GetAdminStatistics)

			admin.GET("/users", handlers.GetUserList)
			admin.PUT("/users/:id/status", handlers.UpdateUserStatus)

			admin.GET("/merchants", handlers.GetMerchantList)
			admin.POST("/merchants", handlers.CreateMerchant)
			admin.PUT("/merchants/:id/status", handlers.UpdateMerchantStatus)

			admin.GET("/admins", handlers.GetAdminList)
			admin.POST("/admins", handlers.CreateAdmin)

			admin.GET("/orders", handlers.GetAllOrders)

			admin.GET("/products", handlers.GetAllProducts)
			admin.PUT("/products/:id/status", handlers.UpdateProductStatus)

			admin.GET("/categories", handlers.GetProductCategoryList)
			admin.POST("/categories", handlers.CreateProductCategory)
			admin.PUT("/categories/:id", handlers.UpdateProductCategory)
			admin.DELETE("/categories/:id", handlers.DeleteProductCategory)

			admin.GET("/banners", handlers.GetBannerList)
			admin.POST("/banners", handlers.CreateBanner)
			admin.PUT("/banners/:id", handlers.UpdateBanner)
			admin.DELETE("/banners/:id", handlers.DeleteBanner)

			admin.GET("/news", handlers.GetNewsListAdmin)
			admin.POST("/news", handlers.CreateNews)
			admin.PUT("/news/:id", handlers.UpdateNews)
			admin.DELETE("/news/:id", handlers.DeleteNews)

			admin.GET("/forums", handlers.GetForumListAdmin)
			admin.PUT("/forums/:id/status", handlers.UpdateForumStatus)

			admin.GET("/messages", handlers.GetMessageBoardList)
			admin.POST("/messages/:id/reply", handlers.ReplyMessage)

			admin.GET("/notices", handlers.GetNoticeListAdmin)
			admin.POST("/notices", handlers.CreateNotice)
			admin.PUT("/notices/:id", handlers.UpdateNotice)
			admin.DELETE("/notices/:id", handlers.DeleteNotice)
		}
	}

	return r
}
