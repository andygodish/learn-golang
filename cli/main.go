package main

import (
	"flag"
)

func main() {
	
	// get videos 'subcommand'
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// Inputs/flags for videos get subcommand
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Video by ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "", "YouTube video ID")
  addTitle := addCmd.String("title", "", "YouTube video Title")
  addUrl := addCmd.String("url", "", "YouTube video URL")
  addImageUrl := addCmd.String("imageurl", "", "YouTube video Image URL")
  addDesc := addCmd.String("desc", "", "YouTube video description")
}
