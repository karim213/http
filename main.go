package main

import (
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/videos", Handler) // call function `handler` whene the server receive request on `/videos`
	http.ListenAndServe(":8080", nil) // Listen on port 8080
}

/**
* This function is called when we listen on http request on `/videos`
* @param http.ResponseWriter is used to return the response
* @param http.Request used to read the request
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // switch between http method type
		case http.MethodGet: // if GET Method


			// fetching the parameters
			id :=  r.FormValue("id") 
			all := r.FormValue("all")

			// customize header to send data as json format
			w.Header().Set("Content-Type", "application/json")

			if all== "true" { // if client want all videos
				json.NewEncoder(w).Encode(getVideos())
			}else if id != "" { // if client want one videos by id
				for _, video := range getVideos() { 
					if video.Id == id {
						json.NewEncoder(w).Encode(video)
						break
					}
				}
			}

		case http.MethodPost: // if POST Method

			// Fetching data
			decoder := json.NewDecoder(r.Body)
			var data video
			err := decoder.Decode(&data)
			if err != nil {
				panic(err)
			}

			// construct new data as video Type
			video := video{
				Id: data.Id,
				Title: data.Title,
				Description: data.Description,
				ImageUrl: data.ImageUrl, 
				Url: data.Url, 
			}
			videos := getVideos() // get all videos
			videos = append(videos,video) // append the new video to the list
		
			saveVideos(videos) // save videos
		
		default: // if the request type is not managed
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		
	}
}
