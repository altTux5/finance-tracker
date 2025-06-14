package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/altTux5/expense-tracker/internal/helpers"
	"github.com/altTux5/expense-tracker/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"omitempty,email"`
	Phone     string `json:"phone" binding:"omitempty"`
	Password  string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty"`
	Password string `json:"password" binding:"required,min=8"`
}

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)
	r.POST("/logout", logoutHandler)
}

func registerHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that either Email or Phone is provided
	if req.Email == "" && req.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "either email or phone must be provided"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Create user model
	user := models.User{
		ID:        uuid.New(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC(), // Use UTC for consistency
		UpdatedAt: time.Now().UTC(),
	}

	// Only set email/phone if they're provided
	if req.Email != "" {
		user.Email = &req.Email
	}
	if req.Phone != "" {
		user.Phone = &req.Phone
	}

	// Insert user
	if err := models.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Generate and return token
	token, err := helpers.GenerateJWT(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"token":   token,
		"user_id": user.ID.String(), // Return user ID for reference
	})
}

func loginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that either Email or Phone is provided
	if req.Email == "" && req.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "either email or phone must be provided"})
		return
	}

	fmt.Println(req)

	user, err := models.GetUserByEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	fmt.Println(user)

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// Generate JWT
	token, err := helpers.GenerateJWT(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"user_id": user.ID.String(),
	})
}

func logoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully. Please delete your token on the client."})
}
