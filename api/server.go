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

	router.POST("/api/customer", server.createCustomer)
	router.GET("/api/customers/list", server.listCustomers)
	router.GET("/api/customer/:customer_id", server.getCustomer)
	router.PUT("/api/customer/:customer_id", server.updateCustomer)
	router.GET("/api/customer/company", server.searchCustomersByCompanyName)
	router.GET("/api/customer/city", server.listCustomersByCity)
	router.DELETE("/api/customer/:customer_id", server.deleteCustomer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
