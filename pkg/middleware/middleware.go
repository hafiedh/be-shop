package middleware

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres"
	"be-shop/internal/app/service/utils"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

const (
	UserData userDataKey = "user_data"
)

type (
	UserCtxReq struct {
		UserID   int    `json:"user_id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	MiddleWareImpl struct {
		dig.In

		UserRepo postgres.UserRepo
	}

	MiddleWare interface {
		AuthUser(next echo.HandlerFunc) echo.HandlerFunc
	}

	userDataKey string
)

func NewMiddleWare(impl MiddleWareImpl) MiddleWare {
	return &impl
}

func (m *MiddleWareImpl) AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		errResponse := models.DefaultResponse{Code: http.StatusUnauthorized}
		ctx := c.Request().Context()
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			errResponse.Message = "Token is required"
			return c.JSON(http.StatusUnauthorized, errResponse)
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		userData, err := utils.Verify(token)
		if err != nil {
			errResponse.Message = err.Error()
			return c.JSON(http.StatusUnauthorized, errResponse)
		}

		var userCtx UserCtxReq
		bt, err := json.Marshal(userData.Data)
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			return c.JSON(http.StatusUnauthorized, errResponse)
		}

		err = json.Unmarshal(bt, &userCtx)
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			return c.JSON(http.StatusUnauthorized, errResponse)
		}

		user, err := m.UserRepo.GetUserByID(ctx, userCtx.UserID)
		if err != nil {
			errResponse.Message = "Unauthorized"
			errResponse.Error = err.Error()
			return c.JSON(http.StatusUnauthorized, errResponse)
		}

		userCtx.Email = user.Email
		userCtx.Username = user.Username
		userCtx.UserID = user.ID

		ctx = context.WithValue(ctx, UserData, userCtx)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
