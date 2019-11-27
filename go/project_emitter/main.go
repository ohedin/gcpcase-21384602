package main

import (
	"fmt"
	"github.com/adeo/project_http/project_emitter/rest"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := rest.NewHandler()
	http.HandleFunc("/", handler.HelloWorld)
	if err := http.ListenAndServe(":" + port, nil); err !=nil{
		fmt.Printf("Aie AIE aAIE [%v]", err)
	}
}


