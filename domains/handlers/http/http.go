package http

import (
	interfaces "learning-backend/domains"
	"learning-backend/domains/models/requests"
	"learning-backend/shared/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Http struct {
	uc interfaces.UseCase
}

func NewHttp(uc interfaces.UseCase) *Http {
	return &Http{uc: uc}
}

func (handler *Http) Login(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.LoginRequest{}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.StructCtx(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	result, err := handler.uc.Login(ctx, requestBody)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.BasicResponse{
		Data: result,
	})
}

func (handler *Http) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.SignUpRequest{}

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.StructCtx(ctx, requestBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	result, err := handler.uc.SignUp(ctx, requestBody)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, result)
}

func (handler *Http) AddMember(c *gin.Context) {
	ctx := c.Request.Context()
	requestBody := &requests.MemberRegistration{}

	// Bind JSON request to struct
	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	// Validate input
	validate := validator.New()
	if err := validate.StructCtx(ctx, requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	// Call the use case
	result, err := handler.uc.AddMember(ctx, requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, result)
}

func (handler *Http) GetDashboardData(c *gin.Context) {
	ctx := c.Request.Context()
	
	// Prepare the request object
	requestBody := &requests.GetDashboardData{}
	
	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Call use case
	result, err := handler.uc.GetDashboardData(ctx, requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, result)
}

func (handler *Http) GetRekapitulasi(c *gin.Context) {
	ctx := c.Request.Context()

	req := &requests.RekapitulasiRequest{}

	// Bind query parameters (if sent via query) or JSON (if sent via body)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: "Invalid request: " + err.Error(),
		})
		return
	}

	// Validate required fields
	// fmt.Println(req)
	if req.Level == "" || req.Page < 1 || req.RowsPerPage < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.BasicResponse{
			Error: "Missing or invalid required fields: level, wilayah, page, rowsPerPage",
		})
		return
	}

	// Call use case
	result, err := handler.uc.GetRekapitulasi(ctx, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.BasicResponse{
			Error: "Failed to fetch rekapitulasi: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *Http) GetAllUser(c *gin.Context) {
	ctx := c.Request.Context()

	// Call the use case
	result, err := handler.uc.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BasicResponse{
			Error: err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, result)
}





