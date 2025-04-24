package db

import (
	"context"
	"ecommerce-go-api/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {

	arg := CreateCustomerParams{

		CustomerID:   "BALLE",
		CompanyName:  util.RandomString(10),
		ContactName:  util.RandomPgText(8),
		ContactTitle: util.RandomPgText(6),
		Address:      util.RandomPgText(6),
		City:         util.RandomPgText(10),
		Region:       util.RandomPgText(6),
		PostalCode:   util.RandomPgText(6),
		Country:      util.RandomPgText(6),
		Phone:        util.RandomPhoneNumber(),
		Fax:          util.RandomPgText(6),
	}

	customer, err := testStore.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.CustomerID, customer.CustomerID)
	require.Equal(t, arg.CompanyName, customer.CompanyName)
	require.Equal(t, arg.ContactName, customer.ContactName)
	require.Equal(t, arg.ContactTitle, customer.ContactTitle)
	require.Equal(t, arg.Address, customer.Address)
	require.Equal(t, arg.City, customer.City)
	require.Equal(t, arg.Region, customer.Region)
	require.Equal(t, arg.PostalCode, customer.PostalCode)
	require.Equal(t, arg.Country, customer.Country)
	require.Equal(t, arg.Phone, customer.Phone)
	require.Equal(t, arg.Fax, customer.Fax)

}
