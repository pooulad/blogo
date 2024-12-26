package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pooulad/blogo/internal/app"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type logLevel int8

const (
	LogLevelDebug logLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	logLevelFatal
)

var (
	logg = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

type api struct {
	engine *gin.Engine
	app    app.App
}

func New(app app.App) *api {
	a := &api{
		engine: gin.New(),
		app:    app,
	}
	a.setupRoutes()
	return a
}

func (a *api) setupRoutes() {
	a.engine.Use(gin.Recovery())

	api := a.engine.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", a.Login)
			auth.POST("/register", a.Register)
		}
		users := api.Group("/users")
		// protect routes here with jwtMiddleware
		users.Use(jwtMiddleware())
		{
			users.GET("", a.GetAllUsers)
			// create user by admin in panel
			users.POST("/create", a.CreateUser)
			users.GET("/get/:id", a.GetUserByID)
			users.PATCH("/update/:id", a.UpdateUserByID)
			users.DELETE("/delete/:id", a.DeleteUserByID)
			users.POST("/follow", a.FollowUserByID)
			users.POST("/unfollow", a.UnFollowUserByID)
			users.GET("/followers/:id", a.GetFollowersByID)
			users.GET("/following/:id", a.GetFollowingByID)
		}
		posts := api.Group("/posts")
		// protect routes here with jwtMiddleware
		posts.Use(jwtMiddleware())
		{
			posts.GET("", a.GetAllPosts)
			posts.POST("/create", a.CreatePost)
			posts.PATCH("/update/:id", a.UpdatePostByID)
			posts.DELETE("/delete/:id", a.DeletePostByID)
			posts.GET("/get/:id", a.GetPostByID)
			posts.POST("/like", a.LikePostByID)
			posts.POST("/unlike", a.UnLikePostByID)
		}
	}

	// Swagger endpoint
	a.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (a *api) Start() error {
	cfg := a.app.GetConfig()
	addr := fmt.Sprintf("%s:%s", cfg.AppUrl, cfg.Port)
	// Set all allowed proxies. I just set it to nil for now
	a.engine.SetTrustedProxies(nil)
	return a.engine.Run(addr)
}

// writeJSON transforms the data into json format
func writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	json = append(json, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Date", time.Now().Format(http.TimeFormat))

	w.WriteHeader(status)
	w.Write(json)
	return nil
}

// errorResponse prints the error details to the log
func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := message

	err := writeJSON(w, status, env, nil)
	if err != nil {
		logError(r, LogLevelError, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse prints the error details to the log
func serverErrorResponse(w http.ResponseWriter, r *http.Request, status int, err interface{}, logErr error) {
	logError(r, LogLevelError, logErr)
	errorResponse(w, r, status, err)
}

// logger is a custom logger which prints the details to the standard output
func logger(level logLevel, message map[string]string) {
	var messages string
	for k, v := range message {
		messages += fmt.Sprintf("%s: '%s'; ", k, v)
	}

	switch level {
	case LogLevelDebug:
		logg.Printf("\t[DEBUG] %s\n", messages)
	case LogLevelInfo:
		logg.Printf("\t[INFO] %s\n", messages)
	case LogLevelWarn:
		logg.Printf("\t[WARN] %s\n", messages)
	case LogLevelError:
		logg.Printf("\t[ERROR] %s\n", messages)
	case logLevelFatal:
		logg.Printf("\t[FATAL] %s\n", messages)
		os.Exit(1)
	}
}

func logError(r *http.Request, logLevel logLevel, err error) {
	logger(logLevel, map[string]string{
		"error":  err.Error(),
		"method": r.Method,
		"url":    r.URL.String(),
	})
}
