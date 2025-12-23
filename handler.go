package main

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RegisterHandler(c *fiber.Ctx) error {
	registerRequest := new(UserRegisterRequest)
	err := c.BodyParser(registerRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	hashedPassword, err := HashPassword(registerRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed hash password",
		})
	}

	var user User
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	user.Password = hashedPassword

	err = CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(UserLoginRequest)
	err := c.BodyParser(loginRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	loginUser, err := GetUser(loginRequest.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	}

	isLoginValid := VerifyPassword(loginUser.Password, loginRequest.Password)
	if !isLoginValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	}

	accessToken, err := GenerateAccessToken(loginUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	refreshToken, err := GenerateRefreshToken(loginUser.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshTokenHandler(c *fiber.Ctx) error {
	refreshTokenRequest := new(RefreshTokenRequest)
	err := c.BodyParser(refreshTokenRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	verifiedToken, err := VerifyToken(refreshTokenRequest.RefreshToken, os.Getenv("JWT_REFRESH_SECRET"), "refresh")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	jwtClaims := verifiedToken.Claims.(jwt.MapClaims)
	userId := int(jwtClaims["userId"].(float64))

	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": accessToken,
	})
}

func CreateHandler(c *fiber.Ctx) error {
	expenseTracker := new(ExpenseRequest)
	err := c.BodyParser(expenseTracker)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)

	var expense Expense
	expense.Title = expenseTracker.Title
	expense.Description = expenseTracker.Description
	expense.Amount = expenseTracker.Amount
	expense.Category = expenseTracker.Category
	expense.UserID = int(authenticatedUser["userId"].(float64))

	createedExpense, err := CreateExpense(expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createedExpense)
}

func UpdateHandler(c *fiber.Ctx) error {
	expenseRequest := new(ExpenseRequest)
	err := c.BodyParser(expenseRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))

	expenseId, _ := c.ParamsInt("id")
	existingExpense, err := GetExpense(expenseId)
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "expense not found",
		})
	}

	if existingExpense.UserID != authenticatedUserId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	existingExpense.Description = expenseRequest.Description
	existingExpense.Amount = expenseRequest.Amount
	existingExpense.Category = expenseRequest.Category
	existingExpense.UserID = int(authenticatedUser["userId"].(float64))

	updatedExpense, err := UpdateExpense(existingExpense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(updatedExpense)
}

func DeleteHandler(c *fiber.Ctx) error {
	expenseId, _ := c.ParamsInt("id")
	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))

	err := DeleteExpense(expenseId, authenticatedUserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

func GetHandler(c *fiber.Ctx) error {
	dateStart := c.Query("dateStart", "")
	dateEnd := c.Query("dateEnd", "")
	filterType := c.Query("filterType", "")
	category := c.Query("category", "")

	var dateStartConverted, dateEndConverted time.Time

	currentTime := time.Now()
	switch filterType {
	case "lastWeek":
		dateStartConverted = currentTime.AddDate(0, 0, -7)
		dateEndConverted = currentTime
	case "lastMonth":
		dateStartConverted = currentTime.AddDate(0, -1, 0)
		dateEndConverted = currentTime
	case "lastThreeMonth":
		dateStartConverted = currentTime.AddDate(0, -3, 0)
		dateEndConverted = currentTime
	default:
		if dateStart != "" {
			dateStartConverted, _ = time.Parse("2006-01-02", dateStart)
		}

		if dateEnd != "" {
			dateEndConverted, _ = time.Parse("2006-01-02", dateEnd)
		}
	}

	authenticatedUser := c.Locals("authenticatedUser").(jwt.MapClaims)
	authenticatedUserId := int(authenticatedUser["userId"].(float64))
	expenses, total, err := GetAllExpenses(authenticatedUserId, dateStartConverted, dateEndConverted, category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  expenses,
		"total": total,
	})
}
