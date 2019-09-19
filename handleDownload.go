package main

import (
  "strings"
  "strconv"
  "net/http"

  "ytm_downloader/loader"
)

func handleDownload(w http.ResponseWriter, r *http.Request) {
  url := r.URL.Path
  urlParts := strings.Split(url, "/")
  id := urlParts[len(urlParts) - 1]

  if id == "" || id == "get" || len(id) != 11 {
    w.WriteHeader(400)

    return
  }

  content, err := loader.LoadVideo(id)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.Header().Set("Content-Length", strconv.Itoa(len(content)))
  w.Write(content)
}
