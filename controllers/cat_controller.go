package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type CatController struct {
	beego.Controller
}

type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// FetchCatImage serves the HTML template.
func (c *CatController) FetchCatImage() {
	c.TplName = "getcatimage.html"
}

// GetCatImage fetches a random cat image from the API and returns JSON.
func (c *CatController) GetCatImage() {
	apiKey, _ := beego.AppConfig.String("catapi::apikey")
	responseChannel := make(chan []CatImage)

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			responseChannel <- nil
			return
		}

		if apiKey != "" {
			req.Header.Set("x-api-key", apiKey)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making API call:", err)
			responseChannel <- nil
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			responseChannel <- nil
			return
		}

		var catImages []CatImage
		err = json.Unmarshal(body, &catImages)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			responseChannel <- nil
			return
		}

		responseChannel <- catImages
	}()

	catImages := <-responseChannel

	if catImages != nil {
		c.Data["json"] = catImages
	} else {
		c.Data["json"] = map[string]string{"error": "Unable to fetch cat image"}
	}
	c.ServeJSON()
}
