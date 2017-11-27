package main

import (
	"net/http"
	"fmt"
	"os"
)

type Reader struct{
	
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	fmt.Println(resp.Body.Read(100, error))
}


