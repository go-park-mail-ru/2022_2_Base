package BaseConfig

var Port = ":8080"
var BasePath = "/api/v1"
var PathLogin = BasePath + "/login"
var PathLogOut = BasePath + "/logout"
var PathSignUp = BasePath + "/signup"
var PathGetUser = BasePath + "/getuser/{username}"
var PathSessions = BasePath + "/session"
var PathDocs = BasePath + "/docs"
var PathMain = BasePath + "/"

var Headers = map[string]string{
	"Access-Control-Allow-Origin":      "http://89.208.198.137:8081",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept",
	//"accept":                           "application/json",
	"Content-Type": "application/json",
}
