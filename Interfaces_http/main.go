package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

)

type logWritter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	lw := logWritter{}
	io.Copy(lw, resp.Body)
	//fmt.Println(resp)
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//io.Copy(os.Stdout, resp.Body)

}
func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Jsut write", len(bs))
	return len(bs), nil
}
