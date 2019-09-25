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
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
