package user

import (
	"context"
	"hungry-baby/businesses/user"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/user/request"
	"hungry-baby/controllers/user/response"
	"net/http"
	"strconv"

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

func (ctrl *UserController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	status := c.QueryParam("status")

	resp, err := ctrl.userUseCase.FindAll(ctx, search, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.User{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *UserController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.userUseCase.Find(ctx, search, status, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.User{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *UserController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	status := c.QueryParam("status")
	resp, err := ctrl.userUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.userUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.User{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.User{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
