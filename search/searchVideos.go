package search

import (
  "flag"
  "net/http"

  "google.golang.org/api/googleapi/transport"
  "google.golang.org/api/youtube/v3"

  "ytm_downloader/config"
)

var (
  maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func SearchVideos(searchTerm string) (SearchResult, error) {
  flag.Parse()

  config, err := config.LoadConfig()
  if err != nil {
    return SearchResult{}, err
  }

  client := &http.Client{
    Transport: &transport.APIKey{Key: config.YouTubeApiToken},
  }

  service, err := youtube.New(client)
  if err != nil {
    return SearchResult{}, err
  }

  // Make the API call to YouTube.
  call := service.Search.List("id,snippet").
            Q(searchTerm).
            MaxResults(*maxResults)
  response, err := call.Do()
  if err != nil {
    return SearchResult{}, err
  }

  result := SearchResult{
    Videos: make([]SearchVideo, 0),
  }

  // Iterate through each item and add it to the correct list.
  for _, item := range response.Items {
    switch item.Id.Kind {
      case "youtube#video":
        vid := SearchVideo{
          ID: item.Id.VideoId,
          Title: item.Snippet.Title,
          Channel: item.Snippet.ChannelTitle,
        }
        result.Videos = append(result.Videos, vid)
    }
  }

  return result, nil
}
