package follow_threads

import (
	followthreads "charum/business/follow_threads"
	"charum/helper"
	"charum/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowThreadController struct {
	followThreadUseCase followthreads.UseCase
}

func NewFollowThreadController(followThreadUC followthreads.UseCase) *FollowThreadController {
	return &FollowThreadController{
		followThreadUseCase: followThreadUC,
	}
}

/*
Create
*/

func (ftc *FollowThreadController) Create(c echo.Context) error {
	threadID, err := primitive.ObjectIDFromHex(c.Param("thread-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid thread id",
			Data:    nil,
		})
	}

	userID, err := util.GetUIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
	}

	domain := followthreads.Domain{
		UserID:   userID,
		ThreadID: threadID,
	}

	thread, err := ftc.followThreadUseCase.Create(&domain)
	if err != nil {
		statusCode := http.StatusBadRequest
		if err.Error() == "user already follow this thread" {
			statusCode = http.StatusConflict
		}

		return c.JSON(statusCode, helper.BaseResponse{
			Status:  statusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseThread, err := ftc.followThreadUseCase.DomainToResponse(thread)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, helper.BaseResponse{
		Status:  http.StatusCreated,
		Message: "success to follow thread",
		Data: map[string]interface{}{
			"followThread": responseThread,
		},
	})
}

/*
Read
*/

func (ftc *FollowThreadController) GeyAllByUserID(c echo.Context) error {
	userID, err := util.GetUIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
	}

	domains, err := ftc.followThreadUseCase.GetAllByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseThreads, err := ftc.followThreadUseCase.DomainToResponseArray(domains)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to get all follow threads",
		Data: map[string]interface{}{
			"followThreads": responseThreads,
		},
	})
}

func (ftc *FollowThreadController) GetFollowedThreadByUserID(c echo.Context) error {
	userID, err := primitive.ObjectIDFromHex(c.Param("user-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid user id",
			Data:    nil,
		})
	}

	domains, err := ftc.followThreadUseCase.GetAllByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseThreads, err := ftc.followThreadUseCase.DomainToResponseArray(domains)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to get all followed threads",
		Data: map[string]interface{}{
			"followThreads": responseThreads,
		},
	})
}

func (ftc *FollowThreadController) GetFollowedThreadByToken(c echo.Context) error {
	userID, err := util.GetUIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
	}

	domains, err := ftc.followThreadUseCase.GetAllByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseThreads, err := ftc.followThreadUseCase.DomainToResponseArray(domains)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to get all followed threads",
		Data: map[string]interface{}{
			"followThreads": responseThreads,
		},
	})
}

/*
Update
*/

/*
Delete
*/

func (ftc *FollowThreadController) Delete(c echo.Context) error {
	threadID, err := primitive.ObjectIDFromHex(c.Param("thread-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid thread id",
			Data:    nil,
		})
	}

	userID, err := util.GetUIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
	}

	domain := followthreads.Domain{
		UserID:   userID,
		ThreadID: threadID,
	}

	result, err := ftc.followThreadUseCase.Delete(&domain)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	responseThread, err := ftc.followThreadUseCase.DomainToResponse(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helper.BaseResponse{
		Status:  http.StatusOK,
		Message: "success to unfollow thread",
		Data: map[string]interface{}{
			"unfollowThread": responseThread,
		},
	})
}
