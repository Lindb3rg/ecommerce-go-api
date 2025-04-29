package api

import (
	db "ecommerce-go-api/db/sqlc"
	"ecommerce-go-api/util"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
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
	err := router.SetTrustedProxies(server.config.TrustedProxies)
	if err != nil {
		log.Fatal("Could not establish proxy connection: ", errorResponse(err))
	}

	router.POST("/customer", server.createCustomer)
	router.GET("/customers/list", server.listCustomers)
	router.GET("/customer/:customer_id", server.getCustomer)
	router.GET("/customer/company", server.searchCustomersByCompanyName)
	router.GET("/customer/city", server.listCustomersByCity)
	router.POST("/customer/delete/:customer_id", server.deleteCustomer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
