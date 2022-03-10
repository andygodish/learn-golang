package main

import "fmt"

func main() {
	videos := getVideosDumb()

	fmt.Println(videos)

	videos = getVideos()

	for i, _ := range videos {
		videos[i].Description = videos[i].Description + "\n this is a new line"
	}

	fmt.Println(videos)

	saveVideos(videos)
}
