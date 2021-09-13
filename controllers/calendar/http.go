package calendar

import (
	"context"
	"hungry-baby/businesses/calendar"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/calendar/request"
	"hungry-baby/controllers/calendar/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type CalendarController struct {
	calendarUseCase calendar.Usecase
}

func NewCalendarController(calendarUC calendar.Usecase) *CalendarController {
	return &CalendarController{
		calendarUseCase: calendarUC,
	}
}

func (ctrl *CalendarController) FindAll(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	search := c.QueryParam("search")
	startAt := c.QueryParam("start_at")
	endAt := c.QueryParam("end_at")
	pageToken := c.QueryParam("page_token")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	resp, err := ctrl.calendarUseCase.FindAll(ctx, search, startAt, endAt, pageToken, limit)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Calendar{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController, 0)
}

func (ctrl *CalendarController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id := c.Param("id")
	resp, err := ctrl.calendarUseCase.FindByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CalendarController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.Calendar{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.calendarUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *CalendarController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id := c.Param("id")
	err := ctrl.calendarUseCase.Delete(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, nil, 0)
}
