package file

import (
	"context"
	"hungry-baby/businesses/file"
	controller "hungry-baby/controllers"
	"hungry-baby/controllers/file/request"
	"hungry-baby/controllers/file/response"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type FileController struct {
	fileUseCase file.Usecase
}

func NewFileController(fileUC file.Usecase) *FileController {
	return &FileController{
		fileUseCase: fileUC,
	}
}

func (ctrl *FileController) FindByID(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := ctrl.fileUseCase.FindByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *FileController) Store(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.File{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.fileUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *FileController) Upload(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	// Read file type
	fileType := c.FormValue("type")

	// Upload file to local temporary
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	resp, err := ctrl.fileUseCase.Upload(ctx, fileType, fileType+"/"+strconv.Itoa(ctx.Value("userID").(int)), fileHeader)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}

func (ctrl *FileController) Delete(c echo.Context) error {
	ctx := c.Get("ctx").(context.Context)

	req := request.File{}
	domainReq := req.ToDomain()
	domainReq.ID, _ = strconv.Atoi(c.Param("id"))
	resp, err := ctrl.fileUseCase.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp), 0)
}
