package category

import (
	"hungry-baby/businesses/category"
	"hungry-baby/controllers/category/response"
	"net/http"

	controller "hungry-baby/controllers"

	echo "github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase category.Usecase
}

func NewCategoryController(cu category.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: cu,
	}
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.categoryUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Category{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}
