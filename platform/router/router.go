package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/willprice/mentor-matcher/platform/authenticator"
	"github.com/willprice/mentor-matcher/platform/middleware"
	"github.com/willprice/mentor-matcher/web/app/callback"	
	"github.com/willprice/mentor-matcher/web/app/healthcheck"
	"github.com/willprice/mentor-matcher/web/app/home"
	"github.com/willprice/mentor-matcher/web/app/login"
	"github.com/willprice/mentor-matcher/web/app/logout"
	"github.com/willprice/mentor-matcher/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)
	router.GET("/healthcheck", healthcheck.Handler)

	return router
}
