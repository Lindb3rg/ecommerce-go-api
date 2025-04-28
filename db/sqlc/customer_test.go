package db

import (
	"context"
	"ecommerce-go-api/util"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {

	arg := CreateCustomerParams{

		CustomerID:   util.RandomString(5, true),
		CompanyName:  util.RandomString(10, false),
		ContactName:  util.RandomPgTypeString(8, false),
		ContactTitle: util.RandomPgTypeString(6, false),
		Address:      util.RandomPgTypeString(6, false),
		City:         util.RandomPgTypeString(10, false),
		Region:       util.RandomPgTypeString(6, false),
		PostalCode:   util.RandomPgTypeString(6, false),
		Country:      util.GetRandomCountry(),
		Phone:        util.RandomPhoneNumber(),
		Fax:          util.RandomPgTypeString(6, false),
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

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testStore.GetCustomer(context.Background(), customer1.CustomerID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.CustomerID, customer2.CustomerID)
	require.Equal(t, customer1.CompanyName, customer2.CompanyName)
	require.Equal(t, customer1.ContactName, customer2.ContactName)
	require.Equal(t, customer1.ContactTitle, customer2.ContactTitle)
	require.Equal(t, customer1.Address, customer2.Address)
	require.Equal(t, customer1.City, customer2.City)
	require.Equal(t, customer1.Region, customer2.Region)
	require.Equal(t, customer1.PostalCode, customer2.PostalCode)
	require.Equal(t, customer1.Country, customer2.Country)
	require.Equal(t, customer1.Phone, customer2.Phone)
	require.Equal(t, customer1.Fax, customer2.Fax)

}

func TestUpdateCustomer(t *testing.T) {

	customer1 := createRandomCustomer(t)
	arg := UpdateCustomerParams{

		CustomerID:   customer1.CustomerID,
		CompanyName:  util.RandomString(10, false),
		ContactName:  util.RandomPgTypeString(8, false),
		ContactTitle: util.RandomPgTypeString(6, false),
		Address:      util.RandomPgTypeString(6, false),
		City:         util.RandomPgTypeString(10, false),
		Region:       util.RandomPgTypeString(6, false),
		PostalCode:   util.RandomPgTypeString(6, false),
		Country:      util.RandomPgTypeString(6, false),
		Phone:        util.RandomPhoneNumber(),
		Fax:          util.RandomPgTypeString(6, false),
	}

	customer2, err := testStore.UpdateCustomer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, arg.CustomerID, customer2.CustomerID)
	require.Equal(t, arg.CompanyName, customer2.CompanyName)
	require.Equal(t, arg.ContactName, customer2.ContactName)
	require.Equal(t, arg.ContactTitle, customer2.ContactTitle)
	require.Equal(t, arg.Address, customer2.Address)
	require.Equal(t, arg.City, customer2.City)
	require.Equal(t, arg.Region, customer2.Region)
	require.Equal(t, arg.PostalCode, customer2.PostalCode)
	require.Equal(t, arg.Country, customer2.Country)
	require.Equal(t, arg.Phone, customer2.Phone)
	require.Equal(t, arg.Fax, customer2.Fax)

}

func TestSearchCustomersByName(t *testing.T) {

	customer := createRandomCustomer(t)
	searchTerm := util.FormatIntoPgTypeText(customer.CompanyName)

	customers, err := testStore.SearchCustomersByCompanyName(context.Background(), searchTerm)

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	found := false

	for _, c := range customers {
		if c.CustomerID == customer.CustomerID {
			found = true

			require.Equal(t, customer.CompanyName, c.CompanyName)
			require.Equal(t, customer.ContactName, c.ContactName)
			require.Equal(t, customer.ContactTitle, c.ContactTitle)
			require.Equal(t, customer.ContactTitle, c.ContactTitle)
			require.Equal(t, customer.Address, c.Address)
			require.Equal(t, customer.City, c.City)
			require.Equal(t, customer.Region, c.Region)
			require.Equal(t, customer.PostalCode, c.PostalCode)
			require.Equal(t, customer.Country, c.Country)
			require.Equal(t, customer.Phone, c.Phone)
			require.Equal(t, customer.Fax, c.Fax)

		}
	}
	require.True(t, found, "Created customer not found in search results")

}

func TestCountAllCustomers(t *testing.T) {

	sqlStore, ok := testStore.(*SQLStore)
	require.True(t, ok, "testStore is not a *SQLStore")

	var count int64
	err := sqlStore.connPool.QueryRow(context.Background(), "SELECT COUNT(*) FROM customers").Scan(&count)
	require.NoError(t, err)

	result, err := testStore.CountAllCustomers(context.Background())
	require.NoError(t, err)

	// Compare the results
	require.Equal(t, count, result, "Customer count from function doesn't match actual count in database")
}

func TestCountCustomersByCountry(t *testing.T) {
	sqlStore, ok := testStore.(*SQLStore)
	require.True(t, ok, "testStore is not a *SQLStore")

	rows, err := sqlStore.connPool.Query(context.Background(),
		"SELECT country, COUNT(*) as customer_count FROM customers GROUP BY country ORDER BY COUNT(*) DESC")
	require.NoError(t, err)
	defer rows.Close()

	expectedResults := make(map[string]int64)
	for rows.Next() {
		var country pgtype.Text
		var count int64
		err := rows.Scan(&country, &count)
		require.NoError(t, err)
		expectedResults[country.String] = count
	}
	require.NoError(t, rows.Err())

	result, err := testStore.CountCustomersByCountry(context.Background())
	require.NoError(t, err)

	require.Equal(t, len(expectedResults), len(result), "Result count doesn't match expected count")

	for _, r := range result {
		expectedCount, exists := expectedResults[r.Country.String]
		require.True(t, exists, "Country %s not found in expected results", r.Country.String)
		require.Equal(t, expectedCount, r.CustomerCount, "Count for country %s doesn't match", r.Country.String)
	}
}

func TestListCustomers(t *testing.T) {

	for i := 0; i < 10; i++ {

		createRandomCustomer(t)

	}

	arg := ListCustomersParams{

		Limit:  5,
		Offset: 5,
	}

	customers, err := testStore.ListCustomers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, customers, 5)

	for _, account := range customers {

		require.NotEmpty(t, account)

	}

}

func TestToggleCustomerActiveStatus(t *testing.T) {
	// Create a random customer
	customer := createRandomCustomer(t)

	// Store the original active status
	originalActive := customer.Active

	// Toggle the active status
	updatedCustomer, err := testStore.ToggleCustomerActiveStatus(context.Background(), customer.CustomerID)

	// Check that there was no error
	require.NoError(t, err)

	// Verify the active status was toggled (flipped)
	require.NotEqual(t, originalActive, updatedCustomer.Active,
		"Active status should be toggled from %v to %v", originalActive, updatedCustomer.Active)

	// Optional: Toggle again and verify it goes back to the original value
	toggledAgainCustomer, err := testStore.ToggleCustomerActiveStatus(context.Background(), customer.CustomerID)
	require.NoError(t, err)
	require.Equal(t, originalActive, toggledAgainCustomer.Active,
		"Active status should be toggled back to original value %v", originalActive)
}

func TestListDistinctCountries(t *testing.T) {
	sqlStore, ok := testStore.(*SQLStore)
	require.True(t, ok, "testStore is not a *SQLStore")

	rows, err := sqlStore.connPool.Query(context.Background(),
		"SELECT DISTINCT country FROM customers WHERE country IS NOT NULL ORDER BY country")
	require.NoError(t, err)
	defer rows.Close()

	var expectedResults []string
	for rows.Next() {
		var country string
		err := rows.Scan(&country)
		require.NoError(t, err)
		expectedResults = append(expectedResults, country)
		fmt.Println(country)
	}
	require.NoError(t, rows.Err())

	result, err := testStore.ListDistinctCountries(context.Background())
	require.NoError(t, err)

	require.Equal(t, len(expectedResults), len(result), "Result count doesn't match expected count")

	for i, r := range result {

		require.Equal(t, expectedResults[i], r.String,
			"Country at index %d doesn't match expected", i)
	}
}
