package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math"
	"net/smtp"
	"strconv"
	"time"
)

var (
	month = []string{"january", "febuary", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"}
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

type Pagination struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalElement int
	TotalPage    int
	Account      []model.Accounts
}

func AccountAdminFindByPageable(context *gin.Context) {
	perPage := context.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(perPage)
	pageStr := context.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * size
	var accounts []model.Accounts
	initializers.DB.Where("status=?", 1).Limit(size).Offset(offset).Find(&accounts)
	var totalAccount int64
	initializers.DB.Raw("select count(*) from accounts where status=1 ").Scan(&totalAccount)
	fmt.Println(totalAccount)
	totalPage := math.Ceil(float64(totalAccount/int64(size) + 1))
	previousPage := page - 1
	if previousPage < 0 {
		previousPage = 0
	}
	nextPage := page + 1
	if totalPage == 1 && page == 1 {
		nextPage = 1
	}
	config.CustomResponse(context, 200, "none",
		Pagination{
			NextPage:     nextPage,
			PreviousPage: previousPage,
			CurrentPage:  page,
			TotalElement: int(totalAccount),
			TotalPage:    int(math.Ceil(totalPage)),
			Account:      accounts,
		})
}
func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
func SendMailSimple(subject string, code string, email string) {
	auth := smtp.PlainAuth("", "duydvph09704@fpt.edu.vn", "iepmdsgpxwfclimh", "smtp.gmail.com")
	msg := subject + "\n this is your verification code " + code

	err := smtp.SendMail("smtp.gmail.com:587", auth, "duydvph09704@fpt.edu.vn", []string{email}, []byte(msg))
	if err != nil {
		fmt.Println(err)
	}
}
func CreateAccountAdmin(ctx *gin.Context) {
	var account model.Accounts
	ctx.ShouldBindJSON(&account)
	if initializers.DB.Raw("select * from accounts where username=? and status=0 and TIMESTAMPDIFF(SECOND,time_valid,NOW())>0",
		account.Username).Scan(&account); account.ID != 0 {
		initializers.DB.Delete(&account, account.ID)
	}
	if initializers.DB.Raw("select * from accounts where username=? ", account.Username).Scan(&account); account.ID != 0 {
		config.CustomResponse(ctx, 200, "username used ", nil)
		return
	}
	if initializers.DB.Raw("select * from accounts where phone_number=?", account.PhoneNumber).Scan(&account); account.ID != 0 {
		config.CustomResponse(ctx, 200, "phone number used ", nil)
		return
	}
	if initializers.DB.Raw("select * from accounts where email=?", account.Email).Scan(&account); account.ID != 0 {
		config.CustomResponse(ctx, 200, "email used ", nil)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), 10)
	account.Password = string(hash)
	account.CreateDate = time.Now()
	account.Status = false
	code := EncodeToString(6)
	account.VerificationCode = code
	account.TimeValid = time.Now().Local().Add(time.Minute * time.Duration(30))
	initializers.DB.Save(&account)
	SendMailSimple("Verification Account", code, account.Email)
	config.CustomResponse(ctx, 200, "success", account)

}

func Verify(ctx *gin.Context) {
	code := ctx.Query("code")
	phoneNumber := ctx.Query("phone_number")
	var account model.Accounts
	if err := initializers.DB.Raw("select * from accounts where phone_number=?", phoneNumber).Scan(&account); err != nil {
		config.CustomResponse(ctx, 400, "not found", nil)
	}
	if account.VerificationCode == code && account.TimeValid.After(time.Now()) {
		account.Status = true
		account.VerificationCode = ""
		account.TimeValid = time.Time{}
		config.CustomResponse(ctx, 200, "success", nil)
	}

}

type accountAdminChangePass struct {
	username    string
	oldPassword string
	password    string
	repPassword string
}

func AdminAccountChangePass(ctx *gin.Context) {
	var account model.Accounts
	//var accountChangePass accountAdminChangePass
	//ctx.ShouldBindJSON(&accountChangePass)
	username := ctx.Query("username")
	curPass := ctx.Query("oldPassword")
	password := ctx.Query("password")
	repPassword := ctx.Query("repPassword")
	if initializers.DB.Where("username=?", username).Find(&account); account.ID == 0 {
		fmt.Println(username, account.ID)
		config.CustomResponse(ctx, 200, "not found", nil)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(curPass)); err != nil {
		config.CustomResponse(ctx, 200, "password is incorrect", nil)
		return
	}
	if repPassword != password {
		config.CustomResponse(ctx, 400, "password and rep password not equal", nil)
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	account.Password = string(hash)
	initializers.DB.Save(&account)
	config.CustomResponse(ctx, 200, "success", account)
	return

}
func AdminAccountDeleteById(ctx *gin.Context) {
	var account model.Accounts
	id := ctx.Query("id")
	if err := initializers.DB.Delete(&account, id); err != nil {
		config.CustomResponse(ctx, 400, "can't delete", err)
		return
	} else {
		config.CustomResponse(ctx, 200, "success", err)
		return
	}

}
