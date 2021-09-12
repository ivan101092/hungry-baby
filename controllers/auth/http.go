package auth

import (
	"hungry-baby/businesses/auth"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/auth/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase auth.Usecase
}

func NewAuthController(authUC auth.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) GetGoogleLoginURL(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.authUseCase.GetGoogleLoginURL(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp, 0)
}

func (ctrl *AuthController) VerifyGoogleCode(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.QueryParam("code")
	resp, err := ctrl.authUseCase.VerifyGoogleCode(ctx, code)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
