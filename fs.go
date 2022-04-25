package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed files
var f embed.FS

func checkNilError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	fmt.Println("Testting file server with embed")
	fmt.Println("Running on port :9000")

	fs := http.FileServer(http.FS(f))
	http.Handle("/", fs)

	err := http.ListenAndServe(":9000", nil)
	checkNilError(err)

}
