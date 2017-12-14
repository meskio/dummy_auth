package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var keys = make(map[string]string)

func main() {
	log.Println("Start dummy auth")
	http.HandleFunc("/isvalid", isvalid)
	http.HandleFunc("/exists", exists)
	http.HandleFunc("/add", add)
	http.HandleFunc("/remove", remove)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isvalid(w http.ResponseWriter, r *http.Request) {
	result := map[string]bool{"isvalid": false}

	user := r.FormValue("user")
	key := r.FormValue("key")
	if keys[user] == key {
		result["isvalid"] = true
	}
	log.Println(result)

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(resultBytes)
}

func exists(w http.ResponseWriter, r *http.Request) {
	result := map[string]bool{"exists": false}

	user := r.FormValue("user")
	if keys[user] != "" {
		result["exists"] = true
	}
	log.Println(result)

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(resultBytes)
}

func add(w http.ResponseWriter, r *http.Request) {
	result := map[string]bool{"add": true}

	user := r.FormValue("user")
	key := r.FormValue("key")
	keys[user] = key
	log.Println(result)

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(resultBytes)
}

func remove(w http.ResponseWriter, r *http.Request) {
	result := map[string]bool{"remove": true}

	user := r.FormValue("user")
	delete(keys, user)
	log.Println(result)

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Write(resultBytes)
}
