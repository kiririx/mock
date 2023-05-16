package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var apiConf = map[string]string{}

type MockAPI struct {
	Path     string `json:"path"`
	RespFile string `json:"resp_file"`
}

func init() {
	data, err := ioutil.ReadFile("apis.json")
	if err != nil {
		log.Fatal(err)
	}
	var apis []MockAPI

	err = json.Unmarshal(data, &apis)
	if err != nil {
		log.Fatal(err)
	}
	for _, api := range apis {
		apiConf[api.Path] = api.RespFile
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := apiConf[r.RequestURI[:strings.Index(r.RequestURI, "?")]]
		if resp != "" {
			http.ServeFile(w, r, resp)
		}
	})

	port := os.Getenv("port")
	if port == "" {
		port = "10042"
	}
	addr := ":" + port

	log.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
