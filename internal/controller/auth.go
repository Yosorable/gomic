package controller

import (
	"time"

	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model"
	"github.com/Yosorable/gomic/internal/model/database"
	"github.com/Yosorable/gomic/internal/model/request"
	"github.com/Yosorable/gomic/internal/model/response"
	"github.com/Yosorable/gomic/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type authController struct{}

func (*authController) Login(c *gin.Context) {
	var req request.LoginRequest
	err := c.BindJSON(&req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	if req.UserName == "" || req.Password == "" {
		response.FailWithMessage("Please input correct username or password", c)
		return
	}
	var user database.User
	err = global.DB.Where("name = ?", req.UserName).Limit(1).Find(&user).Error
	if err != nil {
		response.FailWithMessage("Incorrect username or password", c)
		return
	}
	valid := utils.CheckPasswordHash(req.Password, user.PWDHash)
	if !valid {
		response.FailWithMessage("Incorrect username or password", c)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTClaims{
		ID:       user.ID,
		UserName: user.Name,
		IsAdmin:  user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(global.CONFIG.Secret))
	if err != nil {
		response.FailWithMessage("Server error", c)
		return
	}

	response.OkWithData(response.UserResponse{
		ID:       user.ID,
		UserName: user.Name,
		IsAdmin:  user.IsAdmin,
		JWTToken: tokenString,
	}, c)
}

func (*authController) User(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.FailWithMessage("Server JWT error", c)
		return
	}
	u, success := user.(*model.JWTClaims)
	if !success {
		response.FailWithMessage("Server JWT error", c)
		return
	}
	response.OkWithData(response.UserResponse{
		ID:       u.ID,
		UserName: u.UserName,
		IsAdmin:  u.IsAdmin,
	}, c)
}
