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

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.InvalidRequestError
// @Failure 401 {object} models.InvalidCredentialsError
// @Failure 500 {object} models.TokenGenerationError
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	// Placeholder implementation
	var req models.LoginRequest
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
	return c.JSON(models.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	})

}
