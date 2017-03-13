package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var (
	APIAddress = "http://localhost:8080"
)

func ExampleGetSuccessResponseIfACommandIsGive() {
	entity := url.Values{
		"Command": []string{"ANY"},
	}
	rsp, err := http.PostForm(APIAddress+"/task/run", entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Printf("Wrong Response with status: %s", rsp.Status)
		return
	}

	fmt.Printf("Successful Response with status: %s", rsp.Status)
	//Output: Successful Response with status: 200
}
