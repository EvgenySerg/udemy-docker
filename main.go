package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Sum(a int, b int) int {
	return a + b
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello darling%d\n", Sum(1, 2))
}

func home(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("home.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	f, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(f))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	log.Println("starting")
	port := "3000"
	fmt.Println(fmt.Sprintf("server started on port %s", port))
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", home)

	http.ListenAndServe(":"+port, nil)
}
