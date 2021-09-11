package city

import (
	"context"
	"hungry-baby/businesses/city"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/city/request"
	"hungry-baby/controllers/city/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type CityController struct {
	cityUseCase city.Usecase
}

func NewCityController(cityUC city.Usecase) *CityController {
	return &CityController{
		cityUseCase: cityUC,
	}
}

func (ctrl *CityController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	provinceID, _ := strconv.Atoi(c.QueryParam("province_id"))
	status := c.QueryParam("status")

	resp, err := ctrl.cityUseCase.FindAll(ctx, search, provinceID, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.City{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *CityController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	provinceID, _ := strconv.Atoi(c.QueryParam("province_id"))
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.cityUseCase.Find(ctx, search, provinceID, status, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.City{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *CityController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	status := c.QueryParam("status")
	resp, err := ctrl.cityUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CityController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.cityUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CityController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.cityUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CityController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.City{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.cityUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
