package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", http.HandlerFunc(HandleGetVideos))
	http.HandleFunc("/update", HandleUpdateVideos)
	http.ListenAndServe(":8080", nil)
}

func HandleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			panic(err)
			fmt.Fprint(w, "Bad request")
		}

		saveVideos(videos)

	} else {
		w.WriteHeader(405)
		fmt.Fprint(w, "Method not supported.")
	}
}

func HandleGetVideos(w http.ResponseWriter, r *http.Request) {

	videos := getVideos()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Write(videoBytes)

	// for header, value := range r.Header {
	// 	fmt.Printf("key: %v, Value: %v", header, value)
	// }

	// w.Header().Add("TestHeader", "TestValue") // writes a "response" header

	// w.Write([]byte(`hello world`))
}
