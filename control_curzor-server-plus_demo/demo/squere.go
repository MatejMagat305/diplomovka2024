package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	paramValue := []string{"down", "left", "up", "right"} // nastavenie hodnoty pre parameter "d"

	for _, s := range paramValue {

		urlValues := url.Values{}
		urlValues.Add("d", s)
		url0 := fmt.Sprintf("http://localhost:8080/direction?%s", urlValues.Encode())

		for i := 0; i < 100; i++ {
			req, err := http.NewRequest("GET", url0, nil)
			if err != nil {
				panic(err)
			}

			client := http.DefaultClient
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			fmt.Println(resp.Body)

			resp.Body.Close()
		}
	}
}
