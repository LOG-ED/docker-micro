package main

import (
	"fmt"
	"net/http"
	"net/url"
)

var (
	APIAddress = "http://localhost:8080"
)

func ExampleGetSuccessfulResponseIfRequestMethodIsGet() {
	entity := url.Values{}
	rsp, err := http.PostForm(APIAddress+"/task/run", entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch rsp.StatusCode {
	case 200, 201:
		fmt.Printf("Successful Response with status: %s", rsp.Status)
	default:
		fmt.Printf("UnSuccessful Response with status: %s", rsp.Status)
	}

	//Output: Successful Response with status: 200 OK
}

func ExampleGetCreatedResponseIfRequestMethodIsPost() {
	entity := url.Values{}
	rsp, err := http.PostForm(APIAddress+"/task/run", entity)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch rsp.StatusCode {
	case 200, 201:
		fmt.Printf("Successful Response with status: %s", rsp.Status)
	default:
		fmt.Printf("UnSuccessful Response with status: %s", rsp.Status)
	}

	//Output: Successful Response with status: 201 Created
}
