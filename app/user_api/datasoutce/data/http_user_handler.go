package data

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	entities_user "github.com/pepsi/go-fiber/app/user_api/entities"
	userUseCase "github.com/pepsi/go-fiber/app/user_api/usecase"
	"github.com/pepsi/go-fiber/utils"
	"golang.org/x/crypto/bcrypt"
)

type HttpUserHandler struct {
	usecase userUseCase.UserUseCase
}

func NewUserHttpHandler(useCase userUseCase.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{
		usecase: useCase,
	}
}

func (h *HttpUserHandler) CreateUserRegister(c *fiber.Ctx) error {
	var userReq entities_user.User

	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	userReq.Password = string(hashedPassword)

	if err := h.usecase.CreateUserLogin(&userReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(userReq)
}

func (h *HttpUserHandler) UserLogin(c *fiber.Ctx) error {
	var userReq entities_user.User

	// Parse the request body
	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get the user from the database
	user, err := h.usecase.LoginUser(&userReq)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Compare the hashed password (assuming password is hashed in database)
	if userReq.Email != user.Email || userReq.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Email or password is invalid",
		})
	}

	// Generate the JWT token
	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	// Set the token in the user object
	user.Token = tokenString

	// Return the user and token
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *HttpUserHandler) GetALlUser(c *fiber.Ctx) error {
	users, err := h.usecase.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *HttpUserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := h.usecase.GetUserById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *HttpUserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	var updatedUser entities_user.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	updatedUser.ID = uint(id)
	err = h.usecase.UpdateUser(updatedUser.ID, &updatedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedUser)
}
