package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	uri := "https://api.bitflyer.jp/v1/gethealth"
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
