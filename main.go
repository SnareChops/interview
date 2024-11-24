package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("GET /{$}", Index)
	http.HandleFunc("GET /demo", Demo)
	http.HandleFunc("GET /secret", Secret)
	http.HandleFunc("GET /data", GetData)
	http.HandleFunc("GET /data/v2", GetDataV2)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

var data = []string{
	"Mango",
	"Apple",
	"Pineapple",
	"Banana",
	"Kiwi",
	"Orange",
	"Watermelon",
	"Grapes",
	"Papaya",
	"Strawberry",
}

func Index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.html")
}

func Demo(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "demo.html")
}

func Secret(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "secret.html")
}

func GetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	for i := range result {
		if i == ',' {
			result = append(result[:i], result[i+1:]...)
			break
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(result)
}

func GetDataV2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(result)
}
