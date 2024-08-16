package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/go-redis/redis"
	t "auth/api/token"
	"auth/genproto/auth"

	"github.com/gin-gonic/gin"
	md "auth/api/middleware"
	"auth/internal/email"
)

var (
	from     = "qodirovazizbek1129@gmail.com"
	password = "jkzt mtab wvaq ewlm "
)

// RegisterUser handles user registration
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.RegisterReq true "Register User Request"
// @Success 200 {object} string "User registered successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /register [post]
func (h *Handlers) RegisterUser(c *gin.Context) {
	var body auth.RegisterReq
	if err := c.BindJSON(&body); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	password, err := t.HashPassword(body.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	req := &auth.RegisterReq{
		Username:    body.Username,
		Email:       body.Email,
		Password:    password,
		FullName:    body.FullName,
		DateOfBirth: body.DateOfBirth,
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("reg-user", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.LoginReq true "Login Request"
// @Success 200 {string} auth.LoginRes
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /login [post]
func (h *Handlers) LoginUser(c *gin.Context) {
	var req auth.LoginReq
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	res, err := h.Auth.Login(context.Background(), &req)
	if err != nil {
		log.Printf("failed to login user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	token, refToken := t.GenerateJWTToken(res)

	_, err = h.Auth.SaveRefreshToken(context.Background(), &auth.RefToken{UserId: res.Id, Token: refToken})
	if err != nil {
		log.Printf("failed to refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

// ForgotPassword handles forgot password functionality
// @Summary Forgot password
// @Description Request to reset user's password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.GetByEmail true "Email Request"
// @Success 200 {string} string "Password reset email sent successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /forgot-password [post]
func (h *Handlers) ForgotPassword(c *gin.Context) {
	var req auth.GetByEmail
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	_, err := h.Auth.ForgotPassword(context.Background(), &req)
	if err != nil {
		log.Printf("failed to send password reset email: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	forgotPasswordCode := email.GenForgotPassword()

	err = h.RDB.Set(context.Background(), forgotPasswordCode, req.Email, 15*time.Minute).Err()
	if err != nil {
		log.Printf("failed to store forgot password code in Redis: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	err = email.SendVerificationCode(&auth.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification:%s", req.Email, forgotPasswordCode),
		Code:     forgotPasswordCode,
	})

	if err != nil {
		log.Printf("Could not send an email: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent successfully"})
}

// ResetPassword handles password reset
// @Summary Reset password
// @Description Reset user's password with a reset code
// @Tags auth
// @Accept json
// @Produce json
// @Param user body auth.ResetPassReqBody true "Password Reset Request"
// @Success 200 {string} string "Password reset successfully"
// @Failure 400 {string} string "invalid request"
// @Failure 500 {string} string "internal server error"
// @Router /reset-password [post]
func (h *Handlers) ResetPassword(c *gin.Context) {
	var req auth.ResetPassReq
	if err := c.BindJSON(&req); err != nil {
		log.Printf("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	password, err := t.HashPassword(req.NewPassword)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	req.NewPassword = password

	email, err := h.RDB.Get(context.Background(), req.ResetToken).Result()
	if err == redis.Nil {
		log.Printf("forgot password code not found in Redis: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.Email = email

	_, err = h.Auth.ResetPassword(context.Background(), &req)
	if err != nil {
		log.Printf("failed to reset password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// RefreshToken handles refresh token functionality
// @Summary Refresh Token
// @Description Request to reset user's access token
// @Tags auth
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} auth.LoginRes
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /refresh-token [get]
func (h *Handlers) RefreshToken(c *gin.Context) {
	email := c.Query("email")
	req := auth.GetByEmail{
		Email: email,
	}

	res, err := h.Auth.RefreshToken(context.Background(), &req)
	if err != nil {
		log.Printf("failed to refresh token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// ChangeRole godoc
// @Summary Change user role
// @Description Update the role of a user with the specified ID
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body auth.Role true "Change role request"
// @Success 200 {object} string "Role Changed successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /change-role [put]
func (h *Handlers) ChangeRole(c *gin.Context) {
	var req auth.Role
    if err := c.BindJSON(&req); err!= nil {
        log.Printf("failed to bind JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    _, err := h.Auth.ChangeRole(context.Background(), &auth.Role{
        Id:    req.Id,
        Role:   req.Role,
    })
    if err!= nil {
        log.Printf("failed to change role: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func isValidEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegexPattern)
	return re.MatchString(email)
}

func getuserId(ctx *gin.Context) string {
	user_id, err := md.GetUserId(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}
	return user_id
}