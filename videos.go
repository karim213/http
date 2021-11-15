package main

import (
	"io/ioutil"
	"encoding/json"
)

type video struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageUrl string `json:"imageUrl"`
	Url string `json:"url"`
}

func getVideos()(videos []video){
	// Reading videos list from `videos.json`
	fileBytes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}
	// decode json format to video type array
	err = json.Unmarshal(fileBytes, &videos) 
	if err != nil {
		panic(err)
	}

	return videos;
}

func saveVideos(videos []video)(){
	// Parse data to json format
	videoBytes,err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}
	// save data on file
	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}

/**
*@TODO : check if ID is already exist 
*/
func validateVideo(vid video)(bool){
	if(vid.Id == "" || vid.Title == "" || vid.Description == "" || vid.ImageUrl == "" || vid.Url == "" ){
		return false
	}
	return true
}
