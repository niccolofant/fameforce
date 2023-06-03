package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/pkg/http"
	"github.com/ocintnaf/fameforce/usecases"
)

type userController struct {
	router      fiber.Router
	userUsecase usecases.UserUsecase
}

type UserController interface {
	GetAll(ctx *fiber.Ctx) error
}

func NewUserController(
	router fiber.Router,
	userUsecase usecases.UserUsecase,
) *userController {
	return &userController{
		router:      router,
		userUsecase: userUsecase,
	}
}

func (uc *userController) GetAll(ctx *fiber.Ctx) error {
	userDTOs, err := uc.userUsecase.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(http.NewErrorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(http.NewSuccessResponse(fiber.Map{
		"users": userDTOs,
	}))
}

func (ic *userController) RegisterRoutes() {
	ic.router.Get("/", ic.GetAll)
}