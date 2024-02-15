package api

import (
	db "github.com/alifdwt/jedai/server/db/sqlc"
	"github.com/alifdwt/jedai/server/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/categories", server.createCategory)
	router.GET("/categories/:id", server.getCategory)
	router.GET("/categories", server.listCategories)

	router.POST("/courses", server.createCourse)
	router.GET("/courses/:id/:user_id", server.getCourse)
	router.GET("/courses", server.listCourses)

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.GET("/users", server.listUsers)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
