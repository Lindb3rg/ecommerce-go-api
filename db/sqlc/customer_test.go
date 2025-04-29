package db

import (
	"context"
	"ecommerce-go-api/util"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

type ManualControlParams struct {
	City    *pgtype.Text
	Country *pgtype.Text
}

func createRandomCustomer(t *testing.T, manualControl *ManualControlParams) Customer {

	address := util.RandomAddress()

	arg := CreateCustomerParams{

		CustomerID:   util.RandomString(5, true),
		CompanyName:  util.RandomCompanyName(),
		ContactName:  util.RandomContactName(),
		ContactTitle: util.RandomContactTitle(),
		Address:      util.FormatIntoPgTypeText(address.Street),
		Region:       util.RandomRegion(),
		PostalCode:   util.FormatIntoPgTypeText(address.Zip),
		Phone:        util.RandomPhoneNumber(),
		Fax:          util.RandomPhoneNumber(),
	}

	if manualControl != nil {

		if manualControl.City != nil {
			arg.City = *manualControl.City
		} else {
			arg.City = util.FormatIntoPgTypeText(address.City)
		}

		if manualControl.Country != nil {
			arg.Country = *manualControl.Country
		} else {
			arg.Country = util.FormatIntoPgTypeText(address.Country)
		}
	} else {

		arg.City = util.FormatIntoPgTypeText(address.City)
		arg.Country = util.FormatIntoPgTypeText(address.Country)
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
	createRandomCustomer(t, nil)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t, nil)
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

	customer1 := createRandomCustomer(t, nil)
	address := util.RandomAddress()
	arg := UpdateCustomerParams{

		CustomerID:   customer1.CustomerID,
		CompanyName:  util.RandomCompanyName(),
		ContactName:  util.RandomContactName(),
		ContactTitle: util.RandomContactTitle(),
		Address:      util.FormatIntoPgTypeText(address.Street),
		City:         util.FormatIntoPgTypeText(address.City),
		Region:       util.RandomRegion(),
		PostalCode:   util.FormatIntoPgTypeText(address.Zip),
		Country:      util.FormatIntoPgTypeText(address.Country),
		Phone:        util.RandomPhoneNumber(),
		Fax:          util.RandomPhoneNumber(),
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

func TestDeleteCustomer(t *testing.T) {
	customer := createRandomCustomer(t, nil)

	err := testStore.DeleteCustomer(context.Background(), customer.CustomerID)
	require.NoError(t, err, "Failed to delete customer")

	checkIfCustomerExists, err := testStore.GetCustomer(context.Background(), customer.CustomerID)
	fmt.Println(checkIfCustomerExists)

	require.Error(t, err, "Failed to check if customer has been deleted")
	require.Empty(t, checkIfCustomerExists, "Failed to check if customer exists")
}

func TestSearchCustomersByCompanyName(t *testing.T) {

	customer := createRandomCustomer(t, nil)
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

func TestSearchCustomersByContactName(t *testing.T) {

	customer := createRandomCustomer(t, nil)

	customers, err := testStore.SearchCustomersByContactName(context.Background(), customer.ContactName)

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

		createRandomCustomer(t, nil)

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

func TestListCustomersByCountry(t *testing.T) {

	sqlStore, ok := testStore.(*SQLStore)
	require.True(t, ok, "testStore is not a *SQLStore")

	country := "USA"
	searchTerm := util.FormatIntoPgTypeText(country)

	rows, err := sqlStore.connPool.Query(context.Background(),
		"SELECT * FROM customers WHERE country = $1 ORDER BY company_name", searchTerm)
	require.NoError(t, err)
	defer rows.Close()

	var expectedCustomers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.CustomerID,
			&customer.CompanyName,
			&customer.ContactName,
			&customer.ContactTitle,
			&customer.Address,
			&customer.City,
			&customer.Region,
			&customer.PostalCode,
			&customer.Country,
			&customer.Phone,
			&customer.Fax,
			&customer.CreatedAt,
			&customer.Active,
		)
		require.NoError(t, err)
		expectedCustomers = append(expectedCustomers, customer)
	}
	require.NoError(t, rows.Err())

	result, err := testStore.ListCustomersByCountry(context.Background(), searchTerm)
	require.NoError(t, err)

	require.Equal(t, len(expectedCustomers), len(result), "Result count doesn't match expected count")

	for i, expected := range expectedCustomers {
		actual := result[i]
		require.Equal(t, expected.CustomerID, actual.CustomerID)
		require.Equal(t, expected.CompanyName, actual.CompanyName)
		require.Equal(t, expected.ContactName, actual.ContactName)
		require.Equal(t, expected.ContactTitle, actual.ContactTitle)
		require.Equal(t, expected.Address, actual.Address)
		require.Equal(t, expected.City, actual.City)
		require.Equal(t, expected.Region, actual.Region)
		require.Equal(t, expected.PostalCode, actual.PostalCode)
		require.Equal(t, expected.Country, actual.Country)
		require.Equal(t, expected.Phone, actual.Phone)
		require.Equal(t, expected.Fax, actual.Fax)
		require.Equal(t, expected.CreatedAt, actual.CreatedAt)
		require.Equal(t, expected.Active, actual.Active)

	}

}

func TestListCustomersByCity(t *testing.T) {

	sqlStore, ok := testStore.(*SQLStore)
	require.True(t, ok, "testStore is not a *SQLStore")

	city := "London"
	searchTerm := util.FormatIntoPgTypeText(city)

	rows, err := sqlStore.connPool.Query(context.Background(),
		"SELECT * FROM customers WHERE city = $1 ORDER BY company_name", searchTerm)
	require.NoError(t, err)
	defer rows.Close()

	var expectedCustomers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.CustomerID,
			&customer.CompanyName,
			&customer.ContactName,
			&customer.ContactTitle,
			&customer.Address,
			&customer.City,
			&customer.Region,
			&customer.PostalCode,
			&customer.Country,
			&customer.Phone,
			&customer.Fax,
			&customer.CreatedAt,
			&customer.Active,
		)
		require.NoError(t, err)
		expectedCustomers = append(expectedCustomers, customer)
	}
	require.NoError(t, rows.Err())

	result, err := testStore.ListCustomersByCity(context.Background(), searchTerm)
	require.NoError(t, err)

	require.Equal(t, len(expectedCustomers), len(result), "Result count doesn't match expected count")

	for i, expected := range expectedCustomers {
		actual := result[i]
		require.Equal(t, expected.CustomerID, actual.CustomerID)
		require.Equal(t, expected.CompanyName, actual.CompanyName)
		require.Equal(t, expected.ContactName, actual.ContactName)
		require.Equal(t, expected.ContactTitle, actual.ContactTitle)
		require.Equal(t, expected.Address, actual.Address)
		require.Equal(t, expected.City, actual.City)
		require.Equal(t, expected.Region, actual.Region)
		require.Equal(t, expected.PostalCode, actual.PostalCode)
		require.Equal(t, expected.Country, actual.Country)
		require.Equal(t, expected.Phone, actual.Phone)
		require.Equal(t, expected.Fax, actual.Fax)
		require.Equal(t, expected.CreatedAt, actual.CreatedAt)
		require.Equal(t, expected.Active, actual.Active)

	}

}

func TestToggleCustomerActiveStatus(t *testing.T) {
	// Create a random customer
	customer := createRandomCustomer(t, nil)

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
