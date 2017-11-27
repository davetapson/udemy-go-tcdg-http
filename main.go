package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type Reader interface{
	Read([]byte) (int, error)
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	//fmt.Println(resp.Status)
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)

	// create byte slice
	//bs := make([]byte, 99999)
	//// read body into byte slice
	//resp.Body.Read(bs)
	//
	//fmt.Println(string(bs))

	// OR
	io.Copy(os.Stdout, resp.Body)

}

//func () Reader(){
//
//}


