
package controllers

import (
	"bytes"
	"encoding/json"
	"example-beego/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	// "github.com/beego/beego/v2/server/web/context"
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

	// Initialize controller
	controller := &controllers.FavoritesController{}

	tests := []struct {
		name            string
		requestBody     controllers.FavoriteRequest
		mockAPIStatus   int
		mockAPIResponse string
		expectedStatus  int
		expectedError   string
	}{
		{
			name: "successful_save",
			requestBody: controllers.FavoriteRequest{
				ImageID: "test123",
				SubID:   "user123",
			},
			mockAPIStatus:   http.StatusOK,
			mockAPIResponse: `{"id": 1, "image_id": "test123", "status": "success"}`,
			expectedStatus:  http.StatusOK,
		},
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
			expectedError:   "The Cat API Error: error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create test context
			// w := httptest.NewRecorder()
			reqBody, _ := json.Marshal(tc.requestBody)
			r, _ := http.NewRequest("POST", "/favorites", bytes.NewBuffer(reqBody))

			var w *httptest.ResponseRecorder
			// Set up test context
			controller.Ctx = &context.Context{
				ResponseWriter: w,
				Request:        r,
			}

			// Call the handler
			controller.SaveFavorite()

			// Parse response
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

// -----------------------------------------
// func TestGetFavorites(t *testing.T) {
// 	// Initialize controller
// 	controller := &FavoritesController{}

// 	tests := []struct {
// 		name            string
// 		mockAPIStatus   int
// 		mockAPIResponse string
// 		expectedStatus  int
// 		expectedError   string
// 	}{
// 		{
// 			name:            "successful_fetch",
// 			mockAPIStatus:   http.StatusOK,
// 			mockAPIResponse: `[{"id": 1, "image": {"id": "test123", "url": "http://example.com/cat.jpg"}}]`,
// 			expectedStatus:  http.StatusOK,
// 		},
// 		{
// 			name:            "api_error",
// 			mockAPIStatus:   http.StatusInternalServerError,
// 			mockAPIResponse: `{"message": "Internal server error"}`,
// 			expectedStatus:  http.StatusOK,
// 			expectedError:   "Error parsing API response",
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create test context
// 			w := httptest.NewRecorder()
// 			r, _ := http.NewRequest("GET", "/favorites", nil)

// 			// Set up test context
// 			controller.Ctx = &beego.Context{
// 				ResponseWriter: w,
// 				Request:        r,
// 			}

// 			// Call the handler
// 			controller.GetFavorites()

// 			// Parse response
// 			var response map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &response)

// 			// Assert response status
// 			assert.Equal(t, tc.expectedStatus, w.Code)

// 			// Check for expected error
// 			if tc.expectedError != "" {
// 				assert.Equal(t, tc.expectedError, response["error"])
// 			} else {
// 				assert.Nil(t, response["error"])
// 				assert.NotNil(t, response["data"])
// 			}
// 		})
// 	}
// }

// func TestDeleteFavorite(t *testing.T) {
// 	// Initialize controller
// 	controller := &FavoritesController{}

// 	tests := []struct {
// 		name            string
// 		favoriteID      string
// 		mockAPIStatus   int
// 		mockAPIResponse string
// 		expectedStatus  int
// 		expectedError   string
// 	}{
// 		{
// 			name:           "successful_delete",
// 			favoriteID:     "123",
// 			mockAPIStatus:  http.StatusOK,
// 			expectedStatus: http.StatusOK,
// 		},
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
// 			r, _ := http.NewRequest("DELETE", "/favorites/"+tc.favoriteID, nil)

// 			// Set up test context
// 			controller.Ctx = &beego.Context{
// 				ResponseWriter: w,
// 				Request:        r,
// 				Input:          &beego.BeegoInput{},
// 			}
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
