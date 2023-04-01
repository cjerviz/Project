package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/cjerviz/Project/Project/divrhino-test/database"
	"github.com/gofiber/fiber/v2"
	
)

func TestMain(t *testing.T) {
	// initialize the app
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	// create a mock request for testing
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// send the request to the app
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// check the response status code is 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}