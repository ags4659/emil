package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", myFunc)

	router.Run(":8080")
}

func myFunc(c *gin.Context) {
	str := c.Query("str")

	data := url.Values{}
	data.Add("str", str)

	resp, err := http.PostForm("http://192.168.56.103:1880/app", data)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := string(body)

	c.IndentedJSON(http.StatusOK, res)

	// c.IndentedJSON(http.StatusOK, "asdf")
}
