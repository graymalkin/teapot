package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port         int
	Path         string
	ResponseFile string

	UseSSL   bool
	CertPath string
	KeyPath  string
}

func main() {
	fmt.Printf("Teapot.\n")

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf(" Serving: %s\n", configuration.ResponseFile)

	http.HandleFunc(configuration.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusTeapot)

		log.Println(configuration.Path)
		http.ServeFile(w, r, configuration.ResponseFile)
	})

	if configuration.UseSSL {
		log.Fatal(http.ListenAndServeTLS(
			fmt.Sprintf(":%d", configuration.Port),
			configuration.CertPath,
			configuration.KeyPath,
			nil))
	} else {
		log.Fatal(http.ListenAndServe(
			fmt.Sprintf(":%d", configuration.Port),
			nil))
	}

}
