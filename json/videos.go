package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"decription"`
	ImageUrl    string `json:"imageurl"`
	Url         string `json:"url"`
}

func getVideosDumb() (videos []video) {
	vid1 := video{
		Id:          "test",
		Title:       "test",
		Description: "test",
		ImageUrl:    "Test",
		Url:         "Test",
	}

	return []video{vid1}
}

func getVideos() (videos []video) {
	videoBytes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(videoBytes, &videos)
	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
