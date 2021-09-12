package userChild

import (
	"context"
	"hungry-baby/businesses/userChild"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/userChild/request"
	"hungry-baby/controllers/userChild/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserChildController struct {
	userChildUseCase userChild.Usecase
}

func NewUserChildController(userChildUC userChild.Usecase) *UserChildController {
	return &UserChildController{
		userChildUseCase: userChildUC,
	}
}

func (ctrl *UserChildController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userID := ctx.Value("userID").(int)
	search := c.QueryParam("search")

	resp, err := ctrl.userChildUseCase.FindAll(ctx, search, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.UserChild{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *UserChildController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userID := ctx.Value("userID").(int)
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.userChildUseCase.Find(ctx, search, userID, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.UserChild{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *UserChildController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildUseCase.FindByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChild{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	resp, err := ctrl.userChildUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChild{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChild{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
