package common

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"os"
	"log"
)


/**
return document
 */
func UrlResponse(url string) (*goquery.Document){
	response := GetResponse(url)
	if response.StatusCode == 403 {
		fmt.Println("403:Forbidden")
		os.Exit(1)
	}
	if response.StatusCode == 200 {
		fmt.Println("Response is success:"+url)
		dom, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatalf("Error msg: %s", err.Error())
		}
		return dom
	}else{
		log.Fatalf("Response fail.Status Code:%d", response.StatusCode)
		return nil
	}
}


/**
* return response
*/
func GetResponse(url string) *http.Response {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	response, _ := client.Do(request)
	return response
}

