package handlers

import (
	_ "restapi/docs"
	"restapi/internal"
	"restapi/internal/app/commands"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	service  *internal.Service
	voteChan *commands.VoteService
}

func NewHandler(service *internal.Service, voteChan *commands.VoteService) *Handler {
	return &Handler{service: service, voteChan: voteChan}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		users := api.Group("/users")
		{
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}

		posts := users.Group("/posts")
		{
			posts.GET("/:page", h.ListPosts)
			posts.POST("/", h.CreatePost)
			posts.PUT("/:id", h.UpdatePost)
			posts.DELETE("/:id", h.DeletePost)
			posts.POST("/:id/vote", h.UserVote)
		}
	}

	return router
}
