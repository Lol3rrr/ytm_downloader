package download

import (
  "github.com/Lol3rrr/ytdl"
)

func LoadVideo(id string) ([]byte, error) {
  dlURL := "https://www.youtube.com/watch?v=" + id
  content, err := ytdl.DownloadURL(dlURL)
  if err != nil {
    return nil, err
  }

  return content, nil
}
