package controllers

import (
	"authen-system/internal/auth"
	"authen-system/internal/config"
	"authen-system/internal/database"
	"authen-system/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

type signUpRequest struct {
	FullName    string    `json:"fullName,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Email       string    `json:"email,omitempty"`
	UserName    string    `json:"userName,omitempty"`
	PassWord    string    `json:"password,omitempty"`
	Birthday    time.Time `json:"birthday,omitempty"`
	LatestLogin time.Time `json:"latestLogin,omitempty"`
}

type loginRequest struct {
	UserName    string `json:"fullName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `json:"email,omitempty"`
	PassWord    string `json:"password,omitempty"`
}

type Handler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type handler struct {
	userRepo   database.UserRepository
	authConfig config.Authentication
}

func NewHandler(
	userRepo database.UserRepository,
	authConfig config.Authentication,
) Handler {
	return &handler{
		userRepo:   userRepo,
		authConfig: authConfig,
	}
}

func (h *handler) SignUp(c *gin.Context) {
	signUpReq := signUpRequest{}
	if err := c.ShouldBindJSON(&signUpReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if !signUpReq.isRequestValid() {
		c.JSON(http.StatusBadRequest, gin.H{"err": "email or username or phone should be required"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpReq.PassWord), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "failed to hash password",
		})
		return
	}
	user := models.User{
		FullName:    signUpReq.FullName,
		PhoneNumber: signUpReq.PhoneNumber,
		Email:       signUpReq.Email,
		UserName:    signUpReq.UserName,
		PassWord:    string(hash),
		Birthday:    signUpReq.Birthday,
	}
	if err := h.userRepo.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *handler) Login(c *gin.Context) {
	loginReq := loginRequest{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if !loginReq.isRequestValid() {
		c.JSON(http.StatusBadRequest, gin.H{"err": "email or username or phone should be required"})
		return
	}
	userInfo, err := h.userRepo.Find(loginReq.getQueries())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "failed to find user"})
		return
	}
	if userInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid username or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.PassWord), []byte(loginReq.PassWord)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "invalid username or password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken": auth.GenerateToken(strconv.Itoa(int(userInfo.ID)), h.authConfig.SecretKey),
	})
}

func (r signUpRequest) isRequestValid() bool {
	if r.PassWord == "" {
		return true
	}
	if r.Email == "" && r.UserName == "" && r.PhoneNumber == "" {
		return false
	}
	return true
}

func (r loginRequest) isRequestValid() bool {
	if r.PassWord == "" {
		return true
	}
	if r.Email == "" && r.UserName == "" && r.PhoneNumber == "" {
		return false
	}
	return true
}

func (r loginRequest) getQueries() map[string]string {
	queries := make(map[string]string, 4)
	queries["email"] = r.Email
	queries["userName"] = r.UserName
	queries["phoneNumber"] = r.PhoneNumber
	queries["password"] = r.PassWord
	return queries
}
