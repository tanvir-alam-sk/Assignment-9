package controllers

import (
	"bytes"
	"encoding/json"
	"example-beego/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

func setupApp() {
	// Register the routes with Beego
	beego.Router("/favorites", &controllers.FavoritesController{}, "post:SaveFavorite")
	beego.Router("/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
	beego.Router("/favorites/:id", &controllers.FavoritesController{}, "delete:DeleteFavourite")
}

func TestSaveFavorite(t *testing.T) {

	// Set up the routes
	setupApp()

	tests := []struct {
		name            string
		requestBody     controllers.FavoriteRequest
		mockAPIStatus   int
		mockAPIResponse string
		expectedStatus  int
		expectedError   string
	}{
		{
			name: "missing_image_id",
			requestBody: controllers.FavoriteRequest{
				ImageID: "",
				SubID:   "user123",
			},
			expectedStatus: http.StatusOK,
			expectedError:  "image_id and sub_id are required",
		},
		{
			name: "missing_sub_id",
			requestBody: controllers.FavoriteRequest{
				ImageID: "test123",
				SubID:   "",
			},
			expectedStatus: http.StatusOK,
			expectedError:  "image_id and sub_id are required",
		},
		{
			name: "api_error",
			requestBody: controllers.FavoriteRequest{
				ImageID: "test123",
				SubID:   "user123",
			},
			mockAPIStatus:   http.StatusBadRequest,
			mockAPIResponse: `{"status": "error", "message": "Invalid request"}`,
			expectedStatus:  http.StatusOK,
			expectedError:   "Error making API request",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create test context
			w := httptest.NewRecorder()
			reqBody, _ := json.Marshal(tc.requestBody)
			r, _ := http.NewRequest("POST", "/favorites", bytes.NewBuffer(reqBody))

			beego.BeeApp.Handlers.ServeHTTP(w, r)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			// Assert response status
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Check for expected error
			if tc.expectedError != "" {
				assert.Equal(t, tc.expectedError, response["error"])
			} else {
				assert.Nil(t, response["error"])
				assert.NotNil(t, response["data"])
			}
		})
	}
}

// --------------------------
func TestGetFavorites(t *testing.T) {
	// Initialize controller
	controller := &controllers.FavoritesController{}

	tests := []struct {
		name            string
		mockAPIStatus   int
		mockAPIResponse string
		expectedStatus  int
		expectedError   string
	}{
		{
			name:            "api_error",
			mockAPIStatus:   http.StatusInternalServerError,
			mockAPIResponse: `{"message": "Internal server error"}`,
			expectedStatus:  http.StatusOK,
			expectedError:   "Error parsing API response",
		},
		{
			name:            "api_success",
			mockAPIStatus:   http.StatusOK,
			mockAPIResponse: `{"data": [{"id": 1, "name": "Favorite Item"}]}`,
			expectedStatus:  http.StatusOK,
			expectedError:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a test context and recorder
			w := httptest.NewRecorder()

			// Mock external HTTP request to simulate the API behavior
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			// Simulate the external API's behavior
			httpmock.RegisterResponder("GET", "https://api.thecatapi.com/v1/favourites",
				httpmock.NewStringResponder(tc.mockAPIStatus, tc.mockAPIResponse))

			// Create the request to the Beego handler
			r, _ := http.NewRequest("GET", "/favorites", nil)

			// Serve the HTTP request through the Beego app
			beego.BeeApp.Handlers.ServeHTTP(w, r)

			// Call the controller method (this will trigger the external API request)
			controller.GetFavorites()

			// Parse the response from the controller
			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Error unmarshalling response: %v", err)
			}

			// Assert response status
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Check for expected error
			if tc.expectedError != "" {
				assert.Equal(t, tc.expectedError, response["error"])
			} else {
				assert.Nil(t, response["error"])
				assert.NotNil(t, response["data"])
			}
		})
	}
}

// --------------------------

// func TestDeleteFavorite(t *testing.T) {
// 	// Initialize controller
// 	controller := &controllers.FavoritesController{}

// 	tests := []struct {
// 		name            string
// 		requestBody     controllers.FavoriteRequest
// 		favoriteID      string
// 		mockAPIStatus   int
// 		mockAPIResponse string
// 		expectedStatus  int
// 		expectedError   string
// 	}{

// 		{
// 			name:           "missing_id",
// 			favoriteID:     "",
// 			expectedStatus: http.StatusOK,
// 			expectedError:  "Favorite ID is required",
// 		},
// 		{
// 			name:            "api_error",
// 			favoriteID:      "123",
// 			mockAPIStatus:   http.StatusNotFound,
// 			mockAPIResponse: `{"message": "Favorite not found"}`,
// 			expectedStatus:  http.StatusOK,
// 			expectedError:   "Failed to delete favorite",
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create test context
// 			w := httptest.NewRecorder()
// 			reqBody, _ := json.Marshal(tc.requestBody)
// 			r, _ := http.NewRequest("DELETE", "/favourite/"+tc.favoriteID, bytes.NewBuffer(reqBody))

// 			beego.BeeApp.Handlers.ServeHTTP(w, r)

// 			// Set up test context
// 			// controller.Ctx = &beego.Context{
// 			// 	ResponseWriter: w,
// 			// 	Request:        r,
// 			// 	Input:          &beego.BeegoInput{},
// 			// }

// 			controller.Ctx.Input.SetParam(":id", tc.favoriteID)

// 			// Call the handler
// 			controller.DeleteFavourite()

// 			// Parse response
// 			var response map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &response)

// 			// Assert response status
// 			assert.Equal(t, tc.expectedStatus, w.Code)

// 			// Check for expected error
// 			if tc.expectedError != "" {
// 				assert.Contains(t, response["error"], tc.expectedError)
// 			} else {
// 				assert.Nil(t, response["error"])
// 				assert.NotNil(t, response["message"])
// 			}
// 		})
// 	}
// }
