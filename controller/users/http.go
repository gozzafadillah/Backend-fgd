package users

import (
	"charum/business/users"
	"charum/controller/users/request"
	"charum/controller/users/response"
	"charum/helper"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(userUC users.UseCase) *UserController {
	return &UserController{
		userUseCase: userUC,
	}
}

/*
Create
*/

func (userCtrl *UserController) UserRegister(c echo.Context) error {
	userInput := request.UserRegister{}

	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "fill all the required fields",
			Data:    nil,
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "validation failed",
			Data:    err,
		})
	}

	user, token, err := userCtrl.userUseCase.UserRegister(userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusConflict, helper.BaseResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Status:  http.StatusCreated,
		Message: "success to register",
		Data: map[string]interface{}{
			"token": token,
			"user":  response.FromDomain(user),
		},
	})
}

/*
Read
*/

func (userCtrl *UserController) Login(c echo.Context) error {
	userInput := request.Login{}

	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "fill all the required fields",
			Data:    nil,
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "validation failed",
			Data:    err,
		})
	}

	user, token, err := userCtrl.userUseCase.Login(userInput.ToDomain())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to login",
		Data: map[string]interface{}{
			"token": token,
			"user":  response.FromDomain(user),
		},
	})
}

func (userCtrl *UserController) GetUsersWithPagination(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "page must be a number",
			Data:    nil,
		})
	} else if page < 1 {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "page must be greater than 0",
			Data:    nil,
		})
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "25"
	}
	limitNumber, err := strconv.Atoi(limit)
	if err != nil || limitNumber < 1 {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "limit must be a number and greater than 0",
			Data:    nil,
		})
	}

	sort := c.QueryParam("sort")
	if sort == "" {
		sort = "createdAt"
	} else if !(sort == "id" || sort == "email" || sort == "userName" || sort == "displayName" || sort == "createdAt" || sort == "updatedAt") {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "sort must be id, email, userName, displayName, createdAt, or updatedAt",
			Data:    nil,
		})
	}

	order := c.QueryParam("order")
	if order == "" {
		order = "desc"
	} else if !(order == "asc" || order == "desc") {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "order must be asc or desc",
			Data:    nil,
		})
	}

	users, totalPage, err := userCtrl.userUseCase.GetUsersWithSortAndOrder(page, limitNumber, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to get all users",
		Data: map[string]interface{}{
			"totalPage": totalPage,
			"users":     response.FromDomainArray(users),
		},
	})
}

func (userCtrl *UserController) GetUserByID(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid id",
			Data:    nil,
		})
	}

	user, err := userCtrl.userUseCase.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.BaseResponse{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to get user by id",
		Data: map[string]interface{}{
			"user": response.FromDomain(user),
		},
	})
}

/*
Update
*/

func (userCtrl *UserController) AdminUpdate(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid id",
			Data:    nil,
		})
	}

	userInput := request.AdminUpdate{}
	if c.Bind(&userInput) != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "fill all the required fields",
			Data:    nil,
		})
	}

	if err := userInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "validation failed",
			Data:    err,
		})
	}

	user, err := userCtrl.userUseCase.Update(userID, userInput.ToDomain())

	var statusCode int
	if err == errors.New("failed to get user") {
		statusCode = http.StatusNotFound
	} else if !(err == errors.New("username is already used") || err == errors.New("email is already used")) {
		statusCode = http.StatusConflict
	} else {
		statusCode = http.StatusInternalServerError
	}

	if err != nil {
		return c.JSON(statusCode, helper.BaseResponse{
			Status:  statusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to update user",
		Data: map[string]interface{}{
			"user": response.FromDomain(user),
		},
	})
}

/*
Delete
*/
