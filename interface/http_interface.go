package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs)) 				// or

	// io.Copy(os.Stdout, resp.Body)		//or

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("No. of bytes wrote now : ", len(bs))
	return len(bs), nil
}
