package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type context struct {
	DataOne string
}

func main() {
	tmpl, err := template.ParseFiles("./tmpl/main.gohtml",
		"./tmpl/block_one.gohtml")

	if err != nil {
		log.Fatalf("error template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dataToSend := context{DataOne: "main service called..."}
		tmpl.Execute(w, dataToSend)
	})

	http.HandleFunc("/call1", func(w http.ResponseWriter, r *http.Request) {
		dataToSend := context{DataOne: "call1 service called..."}
		tmpl.Execute(w, dataToSend)
	})

	http.HandleFunc("/getStateData", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["state"]

		if !ok || len(keys) != 1 {
			log.Printf("parameter error")
			return
		}

		stateNum := keys[0]

		fmt.Println("selected state is ", stateNum)

		dataToSend := context{DataOne: stateNum}

		jsonRes, err := json.Marshal(dataToSend)

		if err != nil {
			log.Fatalf("JSON encoding: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRes)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
