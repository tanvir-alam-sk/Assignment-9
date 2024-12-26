package routers

import (
	"example-beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	// beego.Router("/cat-image", &controllers.CatController{}, "get:GetCatImage")

	// Route for rendering the HTML page
	beego.Router("/", &controllers.CatController{}, "get:FetchCatImage")

	// Route for the JSON API
	beego.Router("/cat/get-image", &controllers.CatController{}, "get:GetCatImage")

	beego.Router("/breeds/catagory", &controllers.BreedsController{}, "get:GetBreedsCatarory")
	beego.Router("/breeds/get", &controllers.BreedsController{}, "get:GetBreeds")

	beego.Router("/api/favorites", &controllers.FavoritesController{}, "get:GetFavorites")
	beego.Router("/api/favorites", &controllers.FavoritesController{}, "post:SaveFavorite")
	beego.Router("/favourite/:id", &controllers.FavoritesController{}, "delete:DeleteFavourite")
}
