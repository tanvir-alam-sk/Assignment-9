package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// FavoritesController handles saving favorite images
type FavoritesController struct {
	beego.Controller
}

// FavoriteRequest represents the request body for saving a favorite
type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

// FavoriteResponse represents the response from the Cat API
type FavoriteResponse struct {
	ID     int    `json:"id"`
	Image  string `json:"image_id"`
	Status string `json:"status"`
}

func (c *FavoritesController) FetchFavorite() {
	c.TplName = "favorite.html"
}

// SaveFavorite handles saving a favorite image
func (c *FavoritesController) SaveFavorite() {
	// Fetch API Key and URL
	apiKey, _ := beego.AppConfig.String("catapi::apikey")
	baseURL, _ := beego.AppConfig.String("catapi::apiurl")

	reqBodydata, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	var reqBody FavoriteRequest
	if err := json.Unmarshal(reqBodydata, &reqBody); err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	// Validate input
	if reqBody.ImageID == "" || reqBody.SubID == "" {
		c.Data["json"] = map[string]string{"error": "image_id and sub_id are required"}
		c.ServeJSON()
		return
	}

	// Build the API URL for saving favorites
	apiURL := fmt.Sprintf("%s/favourites", baseURL)

	// Prepare the API request payload
	payload, _ := json.Marshal(reqBody)
	fmt.Println(payload)

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error creating request"}
		c.ServeJSON()
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Make the request to the Cat API
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error making API request"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error reading API response"}
		c.ServeJSON()
		return
	}

	// Parse the response from the Cat API
	var favoriteResponse FavoriteResponse
	err = json.Unmarshal(body, &favoriteResponse)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Already save as a favorite"}
		c.ServeJSON()
		return
	}

	// Handle response
	if resp.StatusCode != http.StatusOK {
		c.Data["json"] = map[string]string{"error": fmt.Sprintf("The Cat API Error: %s", favoriteResponse.Status)}
		c.ServeJSON()
		return
	}

	// Return success response
	c.Data["json"] = map[string]interface{}{
		"message": "Favorite saved successfully",
		"data":    favoriteResponse,
	}
	c.ServeJSON()
}

type Favorite struct {
	ID     int    `json:"id"`
	Image  string `json:"image_id"`
	Status string `json:"status"`
}

// GetFavorites handles the GET request to fetch favorite items
func (c *FavoritesController) GetFavorites() {
	apiKey, _ := beego.AppConfig.String("catapi::apikey")
	baseURL, _ := beego.AppConfig.String("catapi::apiurl")

	apiURL := fmt.Sprintf("%s/favourites", baseURL)

	// Create a new GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error creating request"}
		c.ServeJSON()
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error making API request"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error reading API response"}
		c.ServeJSON()
		return
	}

	// Parse the JSON response
	var responseData interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error parsing API response"}
		c.ServeJSON()
		return
	}

	// Send the data as JSON response
	c.Data["json"] = map[string]interface{}{
		"message": "Fetched favorites successfully",
		"data":    responseData,
	}
	c.ServeJSON()
}
