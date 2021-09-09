package country

import (
	"context"
	"hungry-baby/businesses/country"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/country/request"
	"hungry-baby/controllers/country/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type CountryController struct {
	countryUseCase country.Usecase
}

func NewCountryController(countryUC country.Usecase) *CountryController {
	return &CountryController{
		countryUseCase: countryUC,
	}
}

func (ctrl *CountryController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	status := c.QueryParam("status")

	resp, err := ctrl.countryUseCase.FindAll(ctx, search, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Country{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *CountryController) Find(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	status := c.QueryParam("status")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, total, err := ctrl.countryUseCase.Find(ctx, search, status, page, perpage)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Country{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, total)
}

func (ctrl *CountryController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	status := c.QueryParam("status")
	resp, err := ctrl.countryUseCase.FindByID(ctx, id, status)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CountryController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Country{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.countryUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CountryController) Update(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Country{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.countryUseCase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CountryController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Country{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.countryUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
