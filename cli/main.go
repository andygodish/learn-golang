package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// get videos 'subcommand'
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// Inputs/fields for videos get subcommand
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Video by ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "", "YouTube video ID")
	addTitle := addCmd.String("title", "", "YouTube video Title")
	addUrl := addCmd.String("url", "", "YouTube video URL")
	addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
	addDesc := addCmd.String("desc", "", "YouTube video description")

	// Read input that the user has sent to the application for validation (os package)
	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommand")
		os.Exit(1)
	}

	// Run a switch statement on the second value of the os.Args array (as indicated by the '[1]')
	switch os.Args[1] {
	  case "get":
		HandleGet(getCmd, getAll, getID)
	  case "add":
		HandleAdd(addCmd, addID, addTitle, addDesc, addImageUrl, addUrl)
	  default:
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	// 'all' is false (user did not pass anything in for the 'all' flag) and the id is blank, print back a message to the user
	// describing their errors
	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// if all is true (user passed in all flag) return all videos
	if *all {
		//return all videos
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}

		return
	}

	if *id != "" {
		videos := getVideos()
		id := *id
		for _, video := range videos {
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)

			}
		}
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, description *string, imageurl *string, url *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, description)

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		Imageurl:    *imageurl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, description *string, imageUrl *string, url *string) {
	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *description == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

//////////////

type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

func getVideos() (videos []video) {

	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

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
