package main

import (
  "fmt"
  "net/http"

  "ytm_downloader/api"
)

func main() {
  port := "8080"

  fmt.Printf("Starting on Port %s ... \n", port)

  http.HandleFunc("/song/get/", api.HandleDownload)
  http.HandleFunc("/song/info/", api.HandleInfo)
  http.HandleFunc("/search/videos", api.HandleSearchVideos)
  http.HandleFunc("/import/playlist", api.HandlePlaylistImport)
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
