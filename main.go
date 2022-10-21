package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var data = make(map[string]string)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func shorturl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		param := r.FormValue("param")

		for k := range data {
			if k == param {
				result := map[string]interface{}{
					"url":       param,
					"short_url": data[param],
				}
				response, err := json.Marshal(result)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Write(response)
				return
			}
		}

		gen := RandStringBytesRmndr(5)
		data[param] = fmt.Sprintf("atma.ly/%s", gen)

		result := map[string]interface{}{
			"url":       param,
			"short_url": data[param],
		}

		response, err := json.Marshal(result)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(response)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func url(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		param := r.FormValue("param")

		for key, value := range data {
			if value == param {
				result := map[string]interface{}{
					"url":       key,
					"short_url": param,
				}
				response, err := json.Marshal(result)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Write(response)
				return
			}
		}

	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/shorturl", shorturl)
	http.HandleFunc("/url", url)

	fmt.Println("----Program Starting----")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
