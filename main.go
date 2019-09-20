package main

import (
  "fmt"
  "net/http"

  "ytm_downloader/api"
  "ytm_downloader/search"
)

func main() {
  port := "8080"

  search.SearchVideos("Missio dizzy")

  fmt.Printf("Starting on Port %s ... \n", port)

  http.HandleFunc("/song/get/", api.HandleDownload)
  http.HandleFunc("/song/info/", api.HandleInfo)
  http.HandleFunc("/search/videos", api.HandleSearchVideos)
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
