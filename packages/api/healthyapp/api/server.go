package api

import (
	db "healthyapp/db/sqlc"
	"healthyapp/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config utils.Config
	store  db.Store
	router *gin.Engine
}

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

var (
	DEFAULT_LIMIT = 10
	DEFAULT_OFFET = 0
)

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/login", server.loginUser)

	router.GET("/diary/:id", server.getUserDiary)
	router.POST("/diary", server.createDiary)

	router.GET("/exercise-history/:id", server.getExerciseHistoryOfUser)

	router.GET("/news", server.listNews)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return NewSuccessResponse(data, nil, nil)
}
