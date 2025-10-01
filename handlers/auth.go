package handlers

import (
	"help_desk/database"
	"help_desk/models"
	"help_desk/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func Login(c *fiber.Ctx) error {
	// Placeholder implementation
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User

	// Username ile kullanıcıyı bul
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		log.Printf("Failed login attempt for username: %s (user not found) at %s", req.Username, time.Now().Format("2006-01-02 15:04:05"))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Şifreyi doğrula
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		log.Printf("Failed login attempt for username: %s (invalid password) at %s", req.Username, time.Now().Format("2006-01-02 15:04:05"))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// JWT token oluştur
	token, expiresAt, err := utils.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	// Başarılı login log kaydı
	log.Printf("User %s logged in successfully at %s", user.Username, time.Now().Format("2006-01-02 15:04:05"))

	// Yanıtı döndür
	return c.JSON(LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	})

}
