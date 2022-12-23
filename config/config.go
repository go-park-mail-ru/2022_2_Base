package BaseConfig

var Port = ":8080"
var BasePath = "/api/v1"
var PathLogin = BasePath + "/login"
var PathLogOut = BasePath + "/logout"
var PathSignUp = BasePath + "/signup"
var PathSessions = BasePath + "/session"
var PathDocs = BasePath + "/docs"
var PathMain = BasePath + "/products"
var PathProductsWithDiscount = BasePath + "/productswithdiscount"
var PathProductByID = BasePath + "/products/{id}"
var PathCategory = BasePath + "/category/{category}"
var PathComments = BasePath + "/products/comments/{id}"
var PathSeacrh = BasePath + "/search"
var PathSuggestions = BasePath + "/suggestions"
var PathRecommendations = BasePath + "/recommendations/{id}"
var PathBestProductCategory = BasePath + "/bestproduct/{category}"

var PathProfile = "/profile"
var PathAvatar = "/avatar"
var PathPassword = "/password"
var PathMakeComment = "/makecomment"
var PathFavorites = "/favorites"
var PathInsertIntoFavorites = "/insertintofav"
var PathDeleteFromFavorites = "/deletefromfav"

var PathCart = "/cart"
var PathAddItemToCart = "/insertintocart"
var PathDeleteItemFromCart = "/deletefromcart"
var PathMakeOrder = "/makeorder"
var PathGetOrders = "/orders"
var PathPromo = "/setpromocode"

var Headers = map[string]string{
	//"Access-Control-Allow-Origin":      "http://89.208.198.137:8081",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}

var Promos = []string{
	"SALE500",
	"SALE1500",
}
