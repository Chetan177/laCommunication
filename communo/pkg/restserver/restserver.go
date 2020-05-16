package restserver

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	db "pkg/pkg/database"
	"strings"
)

type Rest struct {
	key    string
	server *echo.Echo
	Port   string
}

func (r *Rest) StartServer() {
	r.server = echo.New()

	r.loadAPI()

	err := r.server.Start(":" + r.Port)
	if err != nil {
		log.Fatal("error starting Echo server")
	}
}

func (r *Rest) loadAPI() {
	login := r.server.Group("/login")
	login.POST("/", r.handleLogin)

	signup := r.server.Group("/signup")
	signup.POST("/", r.handleSignup)

	api := r.server.Group("/api")
	api.POST("/", r.handleAPI)
}

// Login API
func (r *Rest) handleLogin(c echo.Context) error {
	loginDetails := login{}
	c.Bind(&loginDetails)

	_, userPass, err := db.GetData(loginDetails.Email)
	if err != nil{
		if strings.ContainsAny(err.Error(),"sql: no rows in result set"){
			log.Println("User not found ",loginDetails.Email)
			return c.JSON(http.StatusForbidden,response{
				Message: "user not found",
				Status:  403,
			})
		}else {
			log.Println("Database Error ",err)
			return c.JSON(http.StatusInternalServerError,response{
				Message: "database error",
				Status:  500,
			})
		}
	}

	if loginDetails.Password !=userPass{
		log.Println("Incorrect password ",loginDetails.Email)
		return c.JSON(http.StatusBadRequest,response{
			Message: "incorrect password",
			Status:  400,
		})
	}

	log.Println("Logged in", loginDetails.Email)

	return c.JSON(http.StatusOK, response{
		Message: "login successfull",
		Status:  http.StatusOK,
	})
}

func (r *Rest) handleSignup(c echo.Context) error {

	loginDetails := login{}
	c.Bind(&loginDetails)

	err:= db.InsertData(loginDetails.Email,loginDetails.Password)
	if err != nil{
		if strings.ContainsAny(err.Error(),"Duplicate entry"){
			log.Println("User exists ",loginDetails.Email)
			return c.JSON(http.StatusForbidden,response{
				Message: "user exists",
				Status:  403,
			})
		}else {
			log.Println("Database Error ",err)
			return c.JSON(http.StatusInternalServerError,response{
				Message: "database error",
				Status:  500,
			})
		}
	}

	log.Println("New user signup", loginDetails.Email)

	return c.JSON(http.StatusOK, response{
		Message: "sign up successfull",
		Status:  http.StatusOK,
	})
}

func (r *Rest) handleAPI(c echo.Context) error {

	return c.String(http.StatusOK, "Successfull")
}
