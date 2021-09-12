package userChildMeal

import (
	"context"
	"hungry-baby/businesses/userChildMeal"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/userChildMeal/request"
	"hungry-baby/controllers/userChildMeal/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserChildMealController struct {
	userChildMealUseCase userChildMeal.Usecase
}

func NewUserChildMealController(userChildMealUC userChildMeal.Usecase) *UserChildMealController {
	return &UserChildMealController{
		userChildMealUseCase: userChildMealUC,
	}
}

func (ctrl *UserChildMealController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userChildID, _ := strconv.Atoi(c.QueryParam("user_child_id"))
	search := c.QueryParam("search")

	resp, err := ctrl.userChildMealUseCase.FindAll(ctx, search, userChildID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.UserChildMeal{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *UserChildMealController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userChildID, _ := strconv.Atoi(c.QueryParam("user_child_id"))
	search := c.QueryParam("search")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.userChildMealUseCase.Find(ctx, search, userChildID, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.UserChildMeal{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *UserChildMealController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildMealUseCase.FindByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildMealController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChildMeal{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	resp, err := ctrl.userChildMealUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildMealController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChildMeal{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildMealUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *UserChildMealController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.UserChildMeal{}
	req.UserID = ctx.Value("userID").(int)

	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.userChildMealUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
