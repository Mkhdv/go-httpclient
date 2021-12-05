package main



import (
	"fmt"
	"github.com/mkhdv/go-httpclient.git/gohttp"
)

var (
	githubHttpClient = gohttp.New()
)

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}


func getUrls() {

	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}