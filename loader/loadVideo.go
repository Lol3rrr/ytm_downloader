package loader

import (
  "github.com/Lol3rrr/ytdl"
)

func LoadVideo(id string) ([]byte, error) {
  if !vidExists(id) {
    dlURL := "https://www.youtube.com/watch?v=" + id
    rawContent, err := ytdl.DownloadURL(dlURL)
    if err != nil {
      return nil, err
    }

    err = convertWebmToMP3(id, rawContent);
    if err != nil {
      return nil, err
    }
  }

  content, err := readData(id);
  if err != nil {
    return nil, err
  }

  return content, nil;
}
