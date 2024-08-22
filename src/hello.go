package main

import (
	"errors"
	"hello/enc"
	"hello/vks"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

type URI struct {
	Details string `json:"name" uri:"details"`
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println(quote.Go())
	engine := gin.New()
	engine.GET("/test/:details", test_route)

	engine.POST("/encrypt", encrypt_route)

	engine.POST("/decrypt", decrypt_route)

	engine.GET("/keys", get_keys_route)

	engine.Run(":3000")
}

func get_keys_route(context *gin.Context) {
	vks.GetValues()
	context.JSON(http.StatusAccepted, "OK")
}

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

func encrypt_route(context *gin.Context) {
	log.Println("encrypt")
	request, err := context_to_request(context)
	if err != nil {
		return
	}
	response := request.Encrypt()
	context.JSON(http.StatusAccepted, &response)
}

func decrypt_route(context *gin.Context) {
	log.Println("decrypt")
	request, err := context_to_request(context)
	if err != nil {
		return
	}
	response := request.Decrypt()
	context.JSON(http.StatusAccepted, &response)
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
