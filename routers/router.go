package routers

import (
	"example-beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/cat-image", &controllers.CatController{}, "get:GetCatImage")

	// Route for rendering the HTML page
	beego.Router("/cat/image", &controllers.CatController{}, "get:FetchCatImage")

	// Route for the JSON API
	beego.Router("/cat/get-image", &controllers.CatController{}, "get:GetCatImage")

	beego.Router("/breeds", &controllers.BreedsController{}, "get:FetchBreed")
	beego.Router("/breeds/get", &controllers.BreedsController{}, "get:GetBreeds")
}
