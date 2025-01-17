package router

import (
	"echo-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, mc controller.IMemoController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	csrfConfig := middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
	}
	if os.Getenv("GO_ENV") == "dev" {
		csrfConfig.CookieSameSite = http.SameSiteDefaultMode
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/memos")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", mc.GetAllMemos)
	t.GET("/:memoId", mc.GetMemoById)
	t.POST("", mc.CreateMemo)
	t.PUT("/:memoId", mc.UpdateMemo)
	t.DELETE("/:memoId", mc.DeleteMemo)
	return e
}
