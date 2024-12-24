package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedsController struct {
	beego.Controller
}

type Breed struct {
	ID  string `json:"id"`
	Url string `json:"url"`
	// Origin      string `json:"origin"`
	// Description string `json:"description"`
}

// FetchCatImage serves the HTML template.
// func (c *BreedsController) FetchBreed() {
// 	c.TplName = "main.html"
// }

// GetBreeds fetches breed details from the Cat API
func (c *BreedsController) GetBreeds() {
	// Fetch API Key and URL
	apiKey, _ := beego.AppConfig.String("catapi::apikey")
	baseURL, _ := beego.AppConfig.String("catapi::apiurl")

	// Get the breed ID from query parameters
	breedID := c.GetString("breed_id")

	// Build the API URL
	apiURL := fmt.Sprintf("%s/images/search?limit=6&breed_ids=%s", baseURL, breedID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error creating request"}
		c.ServeJSON()
		return
	}

	// Set API Key in the header
	req.Header.Set("x-api-key", apiKey)

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

	var breeds []Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error parsing JSON"}
		c.ServeJSON()
		return
	}

	// Return the first breed (as the API returns a slice)
	if len(breeds) > 0 {
		c.Data["json"] = breeds
	} else {
		c.Data["json"] = map[string]string{"error": "No breed found"}
	}
	c.ServeJSON()
}

type BreedCatagory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	Wikipedia   string `json:"wikipedia_url"`
}

func (c *BreedsController) GetBreedsCatarory() {
	// Fetch API Key and URL
	apiKey, _ := beego.AppConfig.String("catapi::apikey")
	baseURL, _ := beego.AppConfig.String("catapi::apiurl")

	// Get the breed ID from query parameters
	breedID := c.GetString("breed_ids")

	fmt.Println("breedID", breedID)

	// Build the API URL
	apiURL := fmt.Sprintf("%s/breeds", baseURL)

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error creating request"}
		c.ServeJSON()
		return
	}

	// Set API Key in the header
	req.Header.Set("x-api-key", apiKey)

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

	var breedscatagory []BreedCatagory
	err = json.Unmarshal(body, &breedscatagory)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error parsing JSON"}
		c.ServeJSON()
		return
	}

	fmt.Println("req apiURL :", len(breedscatagory))
	// Return the first breed (as the API returns a slice)
	if len(breedscatagory) > 0 {
		c.Data["json"] = breedscatagory
	} else {
		c.Data["json"] = map[string]string{"error": "No breed found"}
	}
	c.ServeJSON()
}
