package BaseConfig

var Port = ":8080"
var BasePath = "/api/v1"
var PathLogin = BasePath + "/login"
var PathLogOut = BasePath + "/logout"
var PathSignUp = BasePath + "/signup"
var PathSessions = BasePath + "/session"
var PathDocs = BasePath + "/docs"
var PathMain = BasePath + "/products"
var PathCategory = BasePath + "/products/{category}"
var PathProfile = BasePath + "/profile"
var PathAvatar = BasePath + "/avatar"
var PathCart = BasePath + "/cart"
var PathAddItemToCart = BasePath + "/insertintocart"
var PathDeleteItemFromCart = BasePath + "/deletefromcart"
var PathMakeOrder = BasePath + "/makeorder"

var Headers = map[string]string{
	//"Access-Control-Allow-Origin":      "http://127.0.0.1:8081",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
