package handlers

import (
	"net/http"
	"time"

	"agriculturalAssistancePlatform/database"
	"agriculturalAssistancePlatform/middleware"
	"agriculturalAssistancePlatform/models"
	"agriculturalAssistancePlatform/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var existingUser models.User
	database.DB.Where("username = ?", req.Username).First(&existingUser)
	if existingUser.ID != 0 {
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

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Balance:  0,
		Status:   1,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "注册失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func UserLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	if user.Status != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "账号已被禁用",
		})
		return
	}

	if !CheckPassword(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成Token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
			"user_info": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"nickname": user.Nickname,
				"phone":    user.Phone,
				"email":    user.Email,
				"avatar":   user.Avatar,
				"balance":  user.Balance,
			},
		},
	})
}

func MerchantLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var merchant models.Merchant
	if err := database.DB.Where("username = ?", req.Username).First(&merchant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	if merchant.Status != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "账号已被禁用",
		})
		return
	}

	if !CheckPassword(req.Password, merchant.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(merchant.ID, "merchant")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成Token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
			"merchant_info": gin.H{
				"id":           merchant.ID,
				"username":     merchant.Username,
				"store_name":   merchant.StoreName,
				"contact_name": merchant.ContactName,
				"contact_phone": merchant.ContactPhone,
			},
		},
	})
}

func AdminLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	var admin models.Admin
	if err := database.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	if admin.Status != 1 {
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "账号已被禁用",
		})
		return
	}

	if !CheckPassword(req.Password, admin.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenerateToken(admin.ID, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成Token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
			"admin_info": gin.H{
				"id":       admin.ID,
				"username": admin.Username,
				"nickname": admin.Nickname,
				"role":     admin.Role,
			},
		},
	})
}

func GetUserInfo(c *gin.Context) {
	userID := middleware.GetUserID(c)
	role := middleware.GetRole(c)

	switch role {
	case "user":
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
				"role":     "user",
			},
		})
	case "merchant":
		var merchant models.Merchant
		if err := database.DB.First(&merchant, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "商家不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": gin.H{
				"id":           merchant.ID,
				"username":     merchant.Username,
				"store_name":   merchant.StoreName,
				"contact_name": merchant.ContactName,
				"contact_phone": merchant.ContactPhone,
				"address":      merchant.Address,
				"role":         "merchant",
			},
		})
	case "admin":
		var admin models.Admin
		if err := database.DB.First(&admin, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "管理员不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取成功",
			"data": gin.H{
				"id":       admin.ID,
				"username": admin.Username,
				"nickname": admin.Nickname,
				"role":     admin.Role,
			},
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "未知的用户角色",
		})
	}
}

func UpdateUserPassword(c *gin.Context) {
	userID := middleware.GetUserID(c)
	role := middleware.GetRole(c)

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误: " + err.Error(),
		})
		return
	}

	if role == "user" {
		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "用户不存在",
			})
			return
		}

		if !CheckPassword(req.OldPassword, user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "原密码错误",
			})
			return
		}

		hashedPassword, err := HashPassword(req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "密码加密失败",
			})
			return
		}

		user.Password = hashedPassword
		user.UpdatedAt = time.Now()
		database.DB.Save(&user)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "密码修改成功",
		})
	} else if role == "merchant" {
		var merchant models.Merchant
		if err := database.DB.First(&merchant, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "商家不存在",
			})
			return
		}

		if !CheckPassword(req.OldPassword, merchant.Password) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "原密码错误",
			})
			return
		}

		hashedPassword, err := HashPassword(req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "密码加密失败",
			})
			return
		}

		merchant.Password = hashedPassword
		merchant.UpdatedAt = time.Now()
		database.DB.Save(&merchant)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "密码修改成功",
		})
	} else if role == "admin" {
		var admin models.Admin
		if err := database.DB.First(&admin, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "管理员不存在",
			})
			return
		}

		if !CheckPassword(req.OldPassword, admin.Password) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "原密码错误",
			})
			return
		}

		hashedPassword, err := HashPassword(req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "密码加密失败",
			})
			return
		}

		admin.Password = hashedPassword
		admin.UpdatedAt = time.Now()
		database.DB.Save(&admin)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "密码修改成功",
		})
	}
}
