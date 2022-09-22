package controller

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var SignUpRequest struct {
	Email       string
	Username    string
	Password    string
	PhoneNumber string
}
var LoginRequest struct {
	Username string
	Password string
}

func SignUp(ctx *gin.Context) {
	if ctx.ShouldBindJSON(&SignUpRequest) != nil {
		config.CustomResponse(ctx, 400, "can't load body", "")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(SignUpRequest.Password), 10)
	if err != nil {
		config.CustomResponse(ctx, 400, "can't load body", "")
		return
	}
	var account model.Accounts
	account.Password = string(hash)
	account.CreateDate = time.Now()
	account.Status = true
	account.Username = SignUpRequest.Username
	account.PhoneNumber = SignUpRequest.PhoneNumber
	account.Email = SignUpRequest.Email
	result := initializers.DB.Create(&account)
	if result != nil {
		config.CustomResponse(ctx, 400, "can't create", "")
		return
	}
	config.CustomResponse(ctx, 200, "", "")
}
func Login(ctx *gin.Context) {
	var account model.Accounts
	initializers.DB.First(&account, "username=?", LoginRequest.Username)
	if account.ID == 0 {
		config.CustomResponse(ctx, 400, "can't find", "")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(LoginRequest.Password))
	if err != nil {
		config.CustomResponse(ctx, 400, "invalid user or pass", "")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix()})
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		config.CustomResponse(ctx, 400, "can't generate", "")

	}
	config.CustomResponse(ctx, 200, "ok", tokenString)
}
