package api

import (
	"errors"
	"net/http"

	db "ecommerce-go-api/db/sqlc"
	"ecommerce-go-api/util"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createCustomerRequest struct {
	CustomerID   string      `json:"customerID" binding:"required"`
	CompanyName  string      `json:"company_name" binding:"required"`
	ContactName  pgtype.Text `json:"contact_name" binding:"required"`
	ContactTitle pgtype.Text `json:"contact_title"`
	Address      pgtype.Text `json:"address"`
	City         pgtype.Text `json:"city"`
	Region       pgtype.Text `json:"region"`
	PostalCode   pgtype.Text `json:"postal_code"`
	Country      pgtype.Text `json:"country"`
	Phone        pgtype.Text `json:"phone"`
	Fax          pgtype.Text `json:"fax"`
}

func (server *Server) createCustomer(ctx *gin.Context) {
	var req createCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCustomerParams{

		CustomerID:   req.CustomerID,
		CompanyName:  req.CompanyName,
		ContactName:  req.ContactName,
		ContactTitle: req.ContactTitle,
		Address:      req.Address,
		City:         req.City,
		Region:       req.Region,
		PostalCode:   req.PostalCode,
		Country:      req.Country,
		Phone:        req.Phone,
		Fax:          req.Fax,
	}

	customer, err := server.store.CreateCustomer(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

type getCustomerRequest struct {
	ID string `uri:"customer_id" binding:"required"`
}

func (server *Server) getCustomer(ctx *gin.Context) {
	var req getCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	customer, err := server.store.GetCustomer(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

type deleteCustomerRequest struct {
	ID string `uri:"customer_id" binding:"required"`
}

func (server *Server) deleteCustomer(ctx *gin.Context) {
	var req deleteCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCustomer(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Deleted customer")
}

type listCustomerRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCustomers(ctx *gin.Context) {
	var req listCustomerRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCustomersParams{

		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	customers, err := server.store.ListCustomers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

type searchCustomersByCompanyNameRequest struct {
	CompanyName string `form:"company_name" binding:"required"`
}

func (server *Server) searchCustomersByCompanyName(ctx *gin.Context) {
	var req searchCustomersByCompanyNameRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	companyNamePgType := util.FormatIntoPgTypeText(req.CompanyName)

	customers, err := server.store.SearchCustomersByCompanyName(ctx, companyNamePgType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

type listCustomersByCityRequest struct {
	City string `form:"city" binding:"required"`
}

func (server *Server) listCustomersByCity(ctx *gin.Context) {
	var req listCustomersByCityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cityPgType := util.FormatIntoPgTypeText(req.City)

	customers, err := server.store.ListCustomersByCity(ctx, cityPgType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}
