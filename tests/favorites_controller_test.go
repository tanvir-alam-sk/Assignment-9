package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"cat-image-viewer/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/adapter/context"
	"github.com/stretchr/testify/assert"
)

func TestSaveFavorite(t *testing.T) {
	// Initialize Beego app
	beego.TestAppInit()

	// Mock API request payload
	favoriteRequest := controllers.FavoriteRequest{
		ImageID: "cat123",
		SubID:   "sub123",
	}

	payload, _ := json.Marshal(favoriteRequest)

	// Create a mock HTTP request
	req := httptest.NewRequest("POST", "/favorite/save", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	// Initialize the controller
	controller := &controllers.FavoritesController{}
	controller.Ctx = &context.Context{
		Request:        req,
		ResponseWriter: rr,
	}

	// Call the SaveFavorite method
	controller.SaveFavorite()

	// Assert the response status code
	assert.Equal(t, 200, rr.Code)

	// Assert the response JSON body
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "Favorite saved successfully", response["message"])
}

func TestGetFavorites(t *testing.T) {
	// Initialize Beego app
	beego.TestAppInit()

	// Create a mock HTTP request for the GetFavorites method
	req := httptest.NewRequest("GET", "/favorite/list", nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	// Initialize the controller
	controller := &controllers.FavoritesController{}
	controller.Ctx = &context.Context{
		Request:        req,
		ResponseWriter: rr,
	}

	// Call the GetFavorites method
	controller.GetFavorites()

	// Assert the response status code
	assert.Equal(t, 200, rr.Code)

	// Assert the response JSON body
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "Fetched favorites successfully", response["message"])
}

func TestDeleteFavorite(t *testing.T) {
	// Initialize Beego app
	beego.TestAppInit()

	// Create a mock HTTP DELETE request
	req := httptest.NewRequest("DELETE", "/favorite/delete/1", nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	// Initialize the controller
	controller := &controllers.FavoritesController{}
	controller.Ctx = &context.Context{
		Request:        req,
		ResponseWriter: rr,
	}

	// Call the DeleteFavorite method
	controller.DeleteFavourite()

	// Assert the response status code
	assert.Equal(t, 200, rr.Code)

	// Assert the response JSON body
	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
	assert.Equal(t, "Favorite deleted successfully", response["message"])
}
