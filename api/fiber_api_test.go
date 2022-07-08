package api

import (
	"net/http"
	"payment_full/db"
	"payment_full/models"
	"payment_full/utils"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func randomAccount() models.Account {
	return models.Account{
		AccountId: uuid.New().String(),
		Owner:     utils.RandomOwner(),
		Balance:   int64(utils.RandomMoney()),
		Currency:  utils.RandomCurrency(),
	}
}

func TestGetAccountAPI(t *testing.T) {
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		// expectedError bool
		// success       bool
		expectedCode int
	}{
		{
			description: "OK",
			route:       "/account/9edfa30c-5ce4-4a84-b53f-97358ee8cc34",
			// success:       true,
			// expectedError: false,
			expectedCode: 200,
		},
		{
			description: "InvalidID",
			route:       "/account/someid",
			// success:       false,
			// expectedError: true,
			expectedCode: 400,
		},
		{
			description: "Id required",
			route:       "/account/",
			// success:       false,
			// expectedError: true,
			expectedCode: 405,
		},
		{
			description: "Not Found",
			route:       "/account/not-found",
			// success:       false,
			// expectedError: true,
			expectedCode: 400,
		},
	}
	// db migrate
	db.ConnectDb()
	app := fiber.New()
	app.Use(logger.New())
	// INITIAL ROUTE
	// InitRouter(app)

	// app.Listen(":8080")
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		resp, _ := app.Test(req)

		// assert.Equalf(t, test.expectedError, resp., test.description)

		// if test.expectedError {
		// 	continue
		// }
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		// body, err := ioutil.ReadAll(resp.Body)
		// assert.Equalf(t, test.success, body.success, test.description)
	}
}

// func TestCreateAccountAPI(t *testing.T) {
// 	account := randomAccount()
// 	tests := []struct {
// 		description string

// 		// Test input
// 		route string

// 		// Expected output
// 		expectedStatusCode int
// 	}{
// 		{
// 			description:        "OK",
// 			route:              "/accounts",
// 			expectedStatusCode: 200,
// 		},
// 		{
// 			description: "Not Found",
// 			route:       "/accountsadasda",
// 			// success:       false,
// 			// expectedError: true,
// 			expectedStatusCode: 404,
// 		},
// 	}
// 	// app := Setup()
// 	// for _, test := range tests {
// 	// 	req, _ := http.NewRequest(
// 	// 		"GET",
// 	// 		test.route,
// 	// 		nil,
// 	// 	)
// 	// 	resp, err := app.Test(req)

// 	// 	assert.Equalf(t, test.expectedError, err != nil, test.description)

// 	// 	if test.expectedError {
// 	// 		continue
// 	// 	}
// 	// 	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
// 	// }

// }
