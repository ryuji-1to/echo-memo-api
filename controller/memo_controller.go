package controller

import (
	"echo-rest-api/model"
	"echo-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IMemoController interface {
	GetAllMemos(c echo.Context) error
	GetMemoById(c echo.Context) error
	CreateMemo(c echo.Context) error
	UpdateMemo(c echo.Context) error
	DeleteMemo(c echo.Context) error
}

type memoController struct {
	mu usecase.IMemoUsecase
}

func NewMemoController(mu usecase.IMemoUsecase) IMemoController {
	return &memoController{mu}
}

func (mc *memoController) GetAllMemos(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	memoRes, err := mc.mu.GetAllMemos(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, memoRes)
}

func (mc *memoController) GetMemoById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("memoId")
	memoId, _ := strconv.Atoi(id)
	memoRes, err := mc.mu.GetMemoById(uint(userId.(float64)), uint(memoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, memoRes)
}

func (mc *memoController) CreateMemo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	memo := model.Memo{}
	if err := c.Bind(&memo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	memo.UserId = uint(userId.(float64))
	memoRes, err := mc.mu.CreateMemo(memo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, memoRes)
}

func (mc *memoController) UpdateMemo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("memoId")
	memoId, _ := strconv.Atoi(id)

	memo := model.Memo{}
	if err := c.Bind(&memo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	memoRes, err := mc.mu.UpdateMemo(memo, uint(userId.(float64)), uint(memoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, memoRes)
}

func (mc *memoController) DeleteMemo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("memoId")
	memoId, _ := strconv.Atoi(id)

	err := mc.mu.DeleteMemo(uint(userId.(float64)), uint(memoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
