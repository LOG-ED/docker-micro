package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var (
	APIAddress = "http://localhost:8080"
)

func ExampleGetSuccessResponseIfRequestMethodIsPost() {
	entity := url.Values{}
	rsp, err := http.PostForm(APIAddress+"/task/run", entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Printf("UnSuccessful Response with status: %s", rsp.Status)
	} else {
		fmt.Printf("Successful Response with status: %s", rsp.Status)
	}

	//Output: Successful Response with status: 200 OK
}

func ExampleGetUnSuccessfulResponseIfRequestMethodIsGet() {
	rsp, err := http.Get(APIAddress + "/task/run")
	if err != nil {
		fmt.Println(err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Printf("UnSuccessful Response with status: %s", rsp.Status)
	} else {
		fmt.Printf("Successful Response with status: %s", rsp.Status)
	}

	//Output: UnSuccessful Response with status: 500 Internal Server Error
}
