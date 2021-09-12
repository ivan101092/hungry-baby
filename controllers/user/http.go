package user

import (
	"context"
	"hungry-baby/businesses/user"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/user/request"
	"hungry-baby/controllers/user/response"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase user.Usecase
}

func NewUserController(userUC user.Usecase) *UserController {
	return &UserController{
		userUseCase: userUC,
	}
}

func (ctrl *UserController) FindByToken(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id := ctx.Value("userID").(int)
	status := c.QueryParam("status")
	resp, err := ctrl.userUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserController) UpdateByToken(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainReq := req.ToDomain()
	domainReq.ID = ctx.Value("userID").(int)
	resp, err := ctrl.userUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
