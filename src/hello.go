package main

import (
	"errors"
	"hello/db"
	docs "hello/docs"
	"hello/enc"
	"hello/vks"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rsc.io/quote"
)

type URI struct {
	Details string `json:"name" uri:"details"`
}

func main() {
	docs.SwaggerInfo.BasePath = "/api/v1"

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println(quote.Go())
	engine := gin.New()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	engine.GET("/test/:details", test_route)

	engine.POST("/encrypt", encrypt_route)

	engine.POST("/decrypt", decrypt_route)

	engine.GET("/keys", get_keys_route)

	engine.POST("/user", add_user_route)
	// TODO - change to RunTLS
	engine.Run(":3000")
}

func get_keys_route(context *gin.Context) {
	vks.GetValues()
	context.JSON(http.StatusAccepted, "OK")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /test [get]
func test_route(context *gin.Context) {
	uri := URI{}
	log.Println("test")
	if err := context.BindUri(&uri); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println(uri)
	context.JSON(http.StatusAccepted, &uri)
}

// @Accept json
// @Produce json
// @Param object body enc.REQUEST true "values to encrypt"
// @Success 200 {object} enc.RESPONSE
// @Router /encrypt [post]
func encrypt_route(context *gin.Context) {
	log.Println("encrypt")
	request, err := context_to_request(context)
	if err != nil {
		return
	}
	response := request.Encrypt()
	context.JSON(http.StatusAccepted, &response)
}

// @Accept json
// @Produce json
// @Param object body enc.REQUEST true "values to decrypt"
// @Success 200 {object} enc.RESPONSE
// @Failure      400  {string}  "Bad Request Error"
// @Router /decrypt [post]
func decrypt_route(context *gin.Context) {
	log.Println("decrypt")
	request, err := context_to_request(context)
	if err != nil {
		return
	}
	response := request.Decrypt()
	context.JSON(http.StatusAccepted, &response)
}

// @Accept json
// @Produce json
// @Param object body user.USER true "values to encrypt"
// @Success 200 {object} db.USER
// @Router /encrypt [post]
func add_user_route(context *gin.Context) {
	log.Println("create user")
	user := db.USER{}

	if err := context.BindJSON(&user); err != nil {
		context.AbortWithError(http.StatusBadRequest, errors.New("cannot read value"))
		return
	}

	if user.PlainTextPassword == "" {
		context.AbortWithError(http.StatusBadRequest, errors.New("must provide plain text password"))
		return
	}
	if user.FullName == "" {
		context.AbortWithError(http.StatusBadRequest, errors.New("must provide Full Name"))
		return
	}
	if user.UserName == "" {
		context.AbortWithError(http.StatusBadRequest, errors.New("must provide user name"))
		return
	}

	db.Create_user(&user)
	context.JSON(http.StatusAccepted, &user)

}

func context_to_request(context *gin.Context) (enc.REQUEST, *gin.Error) {
	request := enc.REQUEST{}

	if err := context.BindJSON(&request); err != nil {
		e := context.AbortWithError(http.StatusBadRequest, errors.New("cannot read value"))
		return request, e
	}
	if len(strings.TrimSpace(request.Text)) == 0 {
		e := context.AbortWithError(http.StatusBadRequest, errors.New("no 'text' value provided"))
		return request, e
	}
	if len(strings.TrimSpace(request.Key)) == 0 {
		e := context.AbortWithError(http.StatusBadRequest, errors.New("no 'key' value provided"))
		return request, e
	}
	return request, nil
}

func add(x int, y int) int {
	return x + y
}
