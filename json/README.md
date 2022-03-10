1. Create your Dockerfile with a dev target allowing you to run golang 1.17 x
2. Define go module, json 
3. Create your main.go with a main function entry point
4. Create a go file that deals specifically with videos
5. Create a video struct: id, title, desc, imageurl, url

6. Create getVidios function that returns a slice of videos (1st iteration)
    - create two video structs (video1 & video2) return them as a slice of type video 
7. Call this dummy getVideos iteration in the main function, printing the results to the server

8. Use ioutil to pull data from json file
    - golang maps fields from the struct into json automatically
    - Use ReadFile() to grap data (this returns a slice of bytes and an error)
    - handle potential error with the panic function

9. Convert bytes to a string using string() (2nd iteration)
    - This is still bytes 

10. use encoding/json to unmarshall json to go type (video struct) - use a pointer to the function defined videos slice 
11. Handle the error associated with json.Unmarshall()
12. Return videos slice
13. Add a dummy new line of text to the list of videos you are printing in your main function using a for loop
    - silence loop vars you don't need
14. Create a saveVideo function that takes an input of a slice of videos
    - does not return any output
15. Use json.Marsahll() to convert go object into json data
    - handle error
16. Write your videobytes to a json file with ioutil.WriteFile
17. Use your main() to write the updated description to a jsonfile using by calling the saveVideo function
18. Use mapping to handle capitalizations in your go struct (json does not capitalize)
