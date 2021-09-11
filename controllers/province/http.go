package province

import (
	"context"
	"hungry-baby/businesses/province"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/province/request"
	"hungry-baby/controllers/province/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ProvinceController struct {
	provinceUseCase province.Usecase
}

func NewProvinceController(provinceUC province.Usecase) *ProvinceController {
	return &ProvinceController{
		provinceUseCase: provinceUC,
	}
}

func (ctrl *ProvinceController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	countryID, _ := strconv.Atoi(c.QueryParam("country_id"))
	status := c.QueryParam("status")

	resp, err := ctrl.provinceUseCase.FindAll(ctx, search, countryID, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Province{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *ProvinceController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	countryID, _ := strconv.Atoi(c.QueryParam("country_id"))
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.provinceUseCase.Find(ctx, search, countryID, status, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Province{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *ProvinceController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	status := c.QueryParam("status")
	resp, err := ctrl.provinceUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *ProvinceController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Province{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.provinceUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *ProvinceController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Province{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.provinceUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *ProvinceController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Province{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.provinceUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
