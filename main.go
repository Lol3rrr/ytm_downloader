package main

import (
  "fmt"
  "net/http"
)

func main() {
  port := "8080"

  fmt.Printf("Starting on Port %s ... \n", port)

  http.HandleFunc("/song/get/", handleDownload)
  if err := http.ListenAndServe(":" + port, nil); err != nil {
    panic(err)
  }
}
