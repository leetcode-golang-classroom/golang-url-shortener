package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/leetcode-golang-classroom/golang-url-shorter/internal/shortener"
	"github.com/vmihailenco/msgpack"
)

func main() {
	address := fmt.Sprintf("http://localhost%s", httpPort())
	redirect := shortener.Redirect{}
	redirect.URL = "https://github.com/leetcode-golang-classroom?tab=repositories"
	body, err := msgpack.Marshal(&redirect)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(address, "application/x-msgpack", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	msgpack.Unmarshal(body, &redirect)
	log.Printf("%v\n", redirect)
}
func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
