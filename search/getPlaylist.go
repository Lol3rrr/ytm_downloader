package search

import (
  "google.golang.org/api/youtube/v3"
)

// Retrieve playlistItems in the specified playlist
func playlistItemsList(service *youtube.Service, part string, playlistId string, pageToken string) *youtube.PlaylistItemListResponse {
  call := service.PlaylistItems.List(part)
  call = call.PlaylistId(playlistId)
  call = call.MaxResults(50)
  if pageToken != "" {
    call = call.PageToken(pageToken)
  }
  response, err := call.Do()
  if err != nil {
    return nil
  }
  return response
}

func GetPlaylist(playlistId string) (SearchResult, error) {
  service, err := getYTService()
  if err != nil {
    return SearchResult{}, err
  }

  result := SearchResult{
    Videos: make([]SearchVideo, 0),
  }

  nextPageToken := ""
  for {
    // Retrieve next set of items in the playlist.
    playlistResponse := playlistItemsList(service, "snippet", playlistId, nextPageToken)

    for _, playlistItem := range playlistResponse.Items {
      vidID := playlistItem.Snippet.ResourceId.VideoId

      tmpInfo, err := GetVideo(vidID)
      if err != nil {
        continue
      }

      tmpResult := SearchVideo{
        ID: tmpInfo.ID,
        Title: tmpInfo.Title,
        Channel: tmpInfo.Channel,
      }

      result.Videos = append(result.Videos, tmpResult)
    }

    // Set the token to retrieve the next page of results
    // or exit the loop if all results have been retrieved.
    nextPageToken = playlistResponse.NextPageToken
    if nextPageToken == "" {
      break
    }
  }

  return result, nil
}
