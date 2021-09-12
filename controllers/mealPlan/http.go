package mealPlan

import (
	"context"
	"hungry-baby/businesses/mealPlan"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/mealPlan/request"
	"hungry-baby/controllers/mealPlan/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type MealPlanController struct {
	mealPlanUseCase mealPlan.Usecase
}

func NewMealPlanController(mealPlanUC mealPlan.Usecase) *MealPlanController {
	return &MealPlanController{
		mealPlanUseCase: mealPlanUC,
	}
}

func (ctrl *MealPlanController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userID := ctx.Value("userID").(int)
	search := c.QueryParam("search")
	status := c.QueryParam("status")

	resp, err := ctrl.mealPlanUseCase.FindAll(ctx, search, userID, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.MealPlan{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *MealPlanController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	userID := ctx.Value("userID").(int)
	search := c.QueryParam("search")
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.mealPlanUseCase.Find(ctx, search, userID, status, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.MealPlan{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *MealPlanController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	status := c.QueryParam("status")
	resp, err := ctrl.mealPlanUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *MealPlanController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.MealPlan{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	resp, err := ctrl.mealPlanUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *MealPlanController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.MealPlan{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = ctx.Value("userID").(int)

	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.mealPlanUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *MealPlanController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.MealPlan{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.mealPlanUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
