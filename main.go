package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // have a router
	router.GET("/", myFunc) // if GET request sent to rounter, carry out myFunc

	router.Run(":8080")
}

func myFunc(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	str := c.Query("str") // Look for the query params named "str"

	data := url.Values{}
	data.Add("str", str) // Add that value of the str query params to a data

	resp, err := http.PostForm("http://192.168.56.103:1880/app", data)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body) // Need this ioutil tool to read the resp
	if err != nil {
		log.Fatal(err)
	}

	res := string(body) // convert the response into string

	c.IndentedJSON(http.StatusOK, res) // convert into json

}
